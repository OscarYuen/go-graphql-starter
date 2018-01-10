package model

type Notification struct {
	From    int64  `json:"from"`
	To      int64  `json:"to"`
	Message string `json:"message"`
}
