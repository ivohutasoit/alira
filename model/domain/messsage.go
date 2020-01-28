package domain

import "encoding/json"

type Response struct {
	Code    int             `json:"code" bson:"code"`
	Status  string          `json:"status" bson:"status"`
	Message string          `json:"message" bson:"message"`
	Error   string          `json:"error" bson:"error"`
	Data    json.RawMessage `json:"data" bson:"data"`
}

type AuthenticatedUser struct {
	UserID   string `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`
}
