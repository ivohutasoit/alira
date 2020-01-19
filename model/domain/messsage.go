package domain

type Response struct {
	Code    int               `json:"code" bson:"code"`
	Status  string            `json:"status" bson:"status"`
	Message string            `json:"message" bson:"message"`
	Error   string            `json:"error" bson:"error"`
	Data    map[string]string `json:"data" bson:"data"`
}
