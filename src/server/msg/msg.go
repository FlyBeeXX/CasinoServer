package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
    Processor.Register(&Request{})
}

type Request struct {
    Method string `json:"method"`
    Params interface{} `json:"params"`
}
