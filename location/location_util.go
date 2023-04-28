package location

import (
	"encoding/json"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/location/model"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

// Use for development
const (
	OCP_ROOT_SERVICE_ACTION = "ObRootServiceInfo"
	OCP_IDC_REGION_ACTION   = "ObIDCRegionInfo"
)

// Use for development
var (
	home = os.Getenv("user.home")
)

type LocationUtil struct {
	OB_VERSION_SQL                string
	PROXY_PLAIN_SCHEMA_SQL_FORMAT string
	TEMPLATE_PART_ID              int
}

func (t *LocationUtil) LoadOcpModel(paramURL string, dataSourceName string, timeout int,
	retryTimes int, retryInternal int64) (model.OcpModel, error) {

	var ocpModel model.OcpModel
	var obServerAddrs []model.ObServerAddr
	ocpModel.SetObServerAddrs(obServerAddrs)

	ocpResponse, err := t.getRemoteOcpResponseOrNull(paramURL, dataSourceName, timeout, retryTimes, retryInternal)
	if err != nil {
		// TODO
		// log.Fatal(err)
	}

	if ocpResponse.IsEmpty() && len(strings.TrimSpace(dataSourceName)) != 0 {
		ocpResponse, err = t.getLocalOcpResponseOrNull(dataSourceName)
		if err != nil {
			// TODO
			// log.Fatal(err)
		}
	}
	if !ocpResponse.IsEmpty() {
		ocpResponseData := ocpResponse.GetData()
		ocpModel.SetClusterId(ocpResponseData.GetObRegionId())
		for _, responseRs := range ocpResponseData.GetRsList() {
			var obServerAddr model.ObServerAddr
			obServerAddr.SetAddress(responseRs.GetAddress())
			obServerAddr.SetSqlPort(responseRs.GetSQLPort())
			obServerAddrs = append(obServerAddrs, obServerAddr)
		}
	}

	if obServerAddrs == nil {
		// TODO
		// log and panic
		// log.Error("load rs list failed dataSource: " + dataSourceName + " paramURL:"
		//                          + paramURL + " response:" + ocpResponse);
	}

	// Get IDC -> Region map if any.
	//String obIdcRegionURL = paramURL.replace(Constants.OCP_ROOT_SERVICE_ACTION,
	//	Constants.OCP_IDC_REGION_ACTION);
	obIdcRegionURL := strings.Replace(paramURL, OCP_ROOT_SERVICE_ACTION, OCP_IDC_REGION_ACTION, -1)

	ocpResponse, err = t.getRemoteOcpIdcRegionOrNull(obIdcRegionURL, dataSourceName, timeout, retryTimes, retryInternal)
	if err != nil {
		// TODO
		// log.Fatal(err)
	}

	if !ocpResponse.IsEmpty() {
		ocpResponseData := ocpResponse.GetData()
		if !ocpResponseData.IsEmpty() && ocpResponseData.GetRsList() != nil {
			for _, idcRegion := range ocpResponseData.GetIDCList() {
				ocpModel.AddIdc2Region(idcRegion.GetIdc(), idcRegion.GetRegion())
			}
		}
	}

	return ocpModel, nil
}

func (t *LocationUtil) getRemoteOcpResponseOrNull(paramURL string, dataSourceName string, timeout int, tryTimes int,
	retryInternal int64) (model.OcpResponse, error) {

	var ocpResponse model.OcpResponse
	var tries int = 0

	// TODO retry
	for ; tries < tryTimes; tries++ {
		content, err := t.loadStringFromUrl(paramURL, timeout)
		if err != nil {
			//return ocpResponse, err
		}
		if err := json.Unmarshal([]byte(content), &ocpResponse); err == nil && ocpResponse.Validate() {
			if len(strings.TrimSpace(dataSourceName)) != 0 {
				t.saveLocalContent(dataSourceName, content)
			}
			return ocpResponse, err
		}
	}

	// TODO return RUNTIME error
	if tries >= tryTimes {

	}
	return ocpResponse, nil
}

func (t *LocationUtil) parseOcpResponse(data []byte) (model.OcpResponse, error) {
	var ocpResponse model.OcpResponse
	err := json.Unmarshal(data, &ocpResponse)
	if err != nil {
		// log.Fatal(err)
		return ocpResponse, err
	}
	return ocpResponse, err
}

func (t *LocationUtil) getLocalOcpResponseOrNull(fileName string) (model.OcpResponse, error) {
	filePath := fmt.Sprintf("%s/conf/obtable//%s", home, fileName)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// TODO
			// log.Fatal("File does not exist.")
			return model.OcpResponse{}, err
		}
		// log.Fatal(err)
		return model.OcpResponse{}, err
	}

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		// TODO
		// log.Fatal(err)
		return model.OcpResponse{}, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return model.OcpResponse{}, err
	}
	return t.parseOcpResponse(data)

}

func (t *LocationUtil) saveLocalContent(fileName string, content string) {

}

func (t *LocationUtil) loadStringFromUrl(url string, timeout int) (string, error) {
	var dumpStr string
	client := &http.Client{}
	client.Timeout = time.Second * time.Duration(timeout)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return dumpStr, err
	}
	response, err := client.Do(request)
	if err != nil {
		return dumpStr, err
	}
	defer response.Body.Close()
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return dumpStr, err
	}
	return string(dump), err
}

func (t *LocationUtil) getPlainString(str string) string {
	var start, end int
	if len(str) > 0 && str[0] == '\'' {
		start = 1
	} else {
		start = 0
	}
	if len(str) > 0 && str[len(str)-1] == '\'' {
		end = len(str) - 1
	} else {
		end = len(str)
	}
	return str[start:end]
}

func (t *LocationUtil) getRemoteOcpIdcRegionOrNull(paramURL string, dataSourceName string, timeout int, tryTimes int,
	retryInternal int64) (model.OcpResponse, error) {

	var ocpResponse model.OcpResponse
	var tries int = 0

	// TODO retry
	for ; tries < tryTimes; tries++ {
		content, err := t.loadStringFromUrl(paramURL, timeout)
		if err != nil {
			//return ocpResponse, err
		}
		if err := json.Unmarshal([]byte(content), &ocpResponse); err == nil && ocpResponse.Validate() {
			if !ocpResponse.IsEmpty() {
				return ocpResponse, err
			}
		}
	}

	return ocpResponse, nil
}
