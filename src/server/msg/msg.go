package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
    Processor.Register(&Request{})
    Processor.Register(&Response{})
}

type Request struct {
    Method string `json:"method"`
    Params interface{} `json:"params"`
    Token string `json:"req"`
}

type Response struct{
    Payload interface{} `json:"payload"`
    Data interface{} `json:"data"`
    Method string `json:"method"`
    Token string `json:"resp"`
    Time int64 `json:"time"`
    Ok int8 `json:"ok"`
}
