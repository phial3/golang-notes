package result

import (
	"encoding/json"
)

const (
	PAGE_NO   = 1
	PAGE_SIZE = 20
)

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

type Order string

type Result struct {
	ErrCode int32       `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}

func ToResult(code int32, msg string, payload interface{}) []byte {
	result := Result{
		ErrCode: code,
		ErrMsg:  msg,
		Data:    payload,
	}
	res, err := json.Marshal(result)
	if err != nil {
		return []byte("{\"errCode\":" + string(code) + ",\"errMsg\": \"" + msg + "\",\"data\":\"convert json error!\"}")
	}
	return res
}
