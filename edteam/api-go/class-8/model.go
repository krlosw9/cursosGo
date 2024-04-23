package main

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GeneralResponse struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
