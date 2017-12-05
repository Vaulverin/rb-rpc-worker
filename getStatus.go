package main

import (
	"encoding/json"
	"net/url"
	"io/ioutil"
)

type GetStatus struct {}

func (gs GetStatus) execute(sParams string) (*PCResult, error) {
	var params DefaultParams
	json.Unmarshal([]byte(sParams), &params)

	u, err := url.Parse(Config.EndPoint.PAPI + "/SmBP/loyaltyManagement/loyaltyProgramMember/1.0.0")
	if err == nil {
		DefaultQueryParams(u, params.msisdn)

		resp, err := DefaultRequest("GET", u.String(), nil)
		if err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				var result *PCResult
				result.HttpCode = resp.Status
				result.ResponseBody = string(body)
				return result, nil
			}
		}
	}
	return nil, err
}
