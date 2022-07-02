package model

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Type    string  `json:"type"`
	Message Message `json:"message"`
}

func (o Order) String() string {
	marshal, _ := json.Marshal(&o)
	return fmt.Sprintf(string(marshal))
}
