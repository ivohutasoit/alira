package alira

import "encoding/json"

type Response struct {
	Code    int             `json:"code" bson:"code"`
	Status  string          `json:"status" bson:"stats"`
	Message string          `json:"message" bson:"message`
	Error   string          `json:"error" bson:"error"`
	Data    json.RawMessage `json:"data" bson:"data"`
}
