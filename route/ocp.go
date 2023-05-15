package route

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
)

type ObOcpModel struct {
	servers   []*ObServerAddr
	clusterId int
}

func newOcpModel(servers []*ObServerAddr, clusterId int) *ObOcpModel {
	return &ObOcpModel{servers, clusterId}
}

func (o *ObOcpModel) GetServerAddressRandomly() *ObServerAddr {
	idx := rand.Intn(len(o.servers))
	return o.servers[idx]
}

func LoadOcpModel(
	configUrl string,
	fileName string,
	timeout time.Duration,
	retryTimes int,
	retryInternal time.Duration) (*ObOcpModel, error) {
	servers := make([]*ObServerAddr, 0, 3)
	var resp obHttpRslistResponse
	err := getRemoteOcpResponseOrNull(configUrl, fileName, timeout, retryTimes, retryInternal, &resp)
	if err != nil {
		return nil, errors.WithMessagef(err, "get remote ocp response, url:%s", configUrl)
	}

	if err != nil && len(strings.TrimSpace(fileName)) != 0 {
		return nil, errors.New("not support get config from local file now")
	}

	for _, server := range resp.Data.RsList {
		// split ip and port, server.Address(xx.xx.xx.xx:xx)
		res := strings.Split(server.Address, ":")
		if len(res) != 2 {
			return nil, errors.Errorf("fail to split ip and port, server:%s", server.Address)
		}
		ip := res[0]
		svrPort, err := strconv.Atoi(res[1])
		if err != nil {
			return nil, errors.Errorf("fail to convert server port to int, port:%s", res[1])
		}
		addr := &ObServerAddr{ip: ip, sqlPort: server.SqlPort, svrPort: svrPort}
		servers = append(servers, addr)
	}

	if len(servers) == 0 {
		return nil, errors.Errorf("failed to load Rslist, url:%s", configUrl)
	}

	return newOcpModel(servers, resp.Data.ObClusterId), nil
}

type obHttpRslistResponse struct {
	Code int `json:"Code"`
	Cost int `json:"Cost"`
	Data struct {
		ObCluster   string `json:"ObCluster"`
		Type        string `json:"Type"`
		ObRegionId  int    `json:"ObRegionId"`
		ObClusterId int    `json:"ObClusterId"`
		RsList      []struct {
			SqlPort int    `json:"sql_port"`
			Address string `json:"address"`
			Role    string `json:"role"`
		} `json:"RsList"`
		ReadonlyRsList []string `json:"ReadonlyRsList"`
		ObRegion       string   `json:"ObRegion"`
		Timestamp      int64    `json:"timestamp"`
	} `json:"Data"`
	Message string `json:"Message"`
	Server  string `json:"Server"`
	Success bool   `json:"Success"`
	Trace   string `json:"Trace"`
}

func getRemoteOcpResponseOrNull(
	url string,
	fileName string,
	timeout time.Duration,
	retryTimes int,
	retryInternal time.Duration,
	resp *obHttpRslistResponse) error {
	var httpResp *http.Response
	var err error
	var times int
	cli := http.Client{Timeout: timeout}
	for times = 0; times < retryTimes; times++ {
		httpResp, err = cli.Get(url)
		if err != nil {
			log.Warn("failed to http get", log.String("url", url), log.Int("times", times))
			time.Sleep(retryInternal)
		} else {
			break
		}
	}
	if times == retryTimes {
		return errors.Errorf("failed to http get after some retry, url:%s, times:%d", url, times)
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	decoder := json.NewDecoder(httpResp.Body)
	for decoder.More() {
		err := decoder.Decode(resp)
		if err != nil {
			return errors.WithMessagef(err, "decode http response, url:%s", url)
		}
	}

	return nil
}
