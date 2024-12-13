/*
Copyright © 2024 JAROD GUICHARD
*/
package models

type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Flags    struct {
		NSFW      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int    `json:"id"`
	Safe bool   `json:"safe"`
	Lang string `json:"lang"`
}

type JokeResponseError struct {
	Error          bool     `json:"error"`
	InternalError  bool     `json:"internalError"`
	Code           uint     `json:"code"`
	Message        string   `json:"message"`
	CausedBy       []string `json:"causedBy"`
	AdditionalInfo string   `json:"additionalInfo"`
	Timestamp      uint     `json:"timestamp"`
}

type JokeResult struct {
	Joke    string
	Error   string
	IsError bool
}
