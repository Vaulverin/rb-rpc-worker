package main

import "encoding/json"

type PCRequestBody struct {
	request string
	params string
}

type PCResponse struct {
	HasInternalErrors bool
	InternalErrors    []string
	Result            *PCResult
}
type PCResult struct {
	HttpCode     string
	ResponseBody string
}

type IProcedureCall interface {
	execute(params string) (*PCResult, error)
}

type ProcedureCallManager struct {}

func (pcManager ProcedureCallManager) GetResponse(requestBody []byte) (response []byte, err error) {
	var reqBdy PCRequestBody
	json.Unmarshal(requestBody, &reqBdy)
	pcResponse := PCResponse{
		HasInternalErrors:false}
	var pCall IProcedureCall
	switch reqBdy.request {
	case "GetStatus":
		pCall = GetStatus{}
		break
	case "RegisterClient":
		break
	case "GetHistory":
		break
	case "GetBalance":
		break
	default:
		pcResponse.HasInternalErrors = true
		pcResponse.InternalErrors = []string{"Request not implemented!"}
	}
	pcResponse.Result, err = pCall.execute(reqBdy.params)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(pcResponse)
	if err != nil {
		res = nil
	}
	return res, err
}
