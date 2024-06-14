/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package route

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
)

// ObConfigServerInfo contains information about rslist and IDC.
type ObConfigServerInfo struct {
	configUrl     string
	file          string // read from local file
	timeout       time.Duration
	retryTimes    int
	retryInterval time.Duration
	rslist        *ObRslist
}

func NewConfigServerInfo() *ObConfigServerInfo {
	return &ObConfigServerInfo{
		rslist: NewRslist(),
	}
}

// GetServerAddressRandomly get one randomly server from all the servers
func (i *ObConfigServerInfo) GetServerAddressRandomly() (*ObServerAddr, error) {
	return i.rslist.GetServerRandomly()
}

// FetchRslist fetch the rslist information from the configUrl using the http get service.
func (i *ObConfigServerInfo) FetchRslist() (*ObRslist, error) {
	var resp obHttpRslistResponse
	err := getConfigServerResponseOrNull(i.configUrl, i.timeout, i.retryTimes, i.retryInterval, &resp)
	if err != nil {
		return nil, errors.WithMessagef(err, "get remote ocp response, url:%s", i.configUrl)
	}

	if err != nil && len(strings.TrimSpace(i.file)) != 0 {
		return nil, errors.New("not support get config from local file now")
	}

	rslist := NewRslist()
	for _, server := range resp.Data.RsList {
		// split ip and port, server.Address(xx.xx.xx.xx:xx)
		res := strings.Split(server.Address, ":")
		if len(res) != 2 {
			return nil, errors.Errorf("fail to split ip and port, server:%s", server.Address)
		}
		ip := res[0]
		if ip == "172.16.46.180" {
			ip = "115.29.212.38"
			println(ip)
		}
		svrPort, err := strconv.Atoi(res[1])
		if err != nil {
			return nil, errors.Errorf("fail to convert server port to int, port:%s", res[1])
		}
		serverAddr := NewObServerAddr(ip, server.SqlPort, svrPort)
		println(serverAddr.String())
		rslist.Append(serverAddr)
	}

	if rslist.Size() == 0 {
		return nil, errors.Errorf("failed to load Rslist, url:%s", i.configUrl)
	}

	return rslist, nil
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

// getConfigServerResponseOrNull parse the response returned by http get and parse the json.
func getConfigServerResponseOrNull(
	url string,
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
			log.Warn("Monitor", nil, "failed to http get", log.String("url", url), log.Int("times", times))
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
