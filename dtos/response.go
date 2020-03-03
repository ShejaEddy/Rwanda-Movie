package dtos

type Response struct {
	Message string      `json:"message,omitemty"`
	Data    interface{} `json:"data,omitemty"`
	Status  bool        `json:"status,omitemty"`
	Code    int         `json:"code,omitemty"`
}
