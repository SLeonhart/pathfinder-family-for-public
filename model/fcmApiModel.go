package model

type FcmApiResponse struct {
	Failure int            `json:"failure"`
	Success int            `json:"success"`
	Results []FcmApiResult `json:"results"`
}

type FcmApiResult struct {
	MessageId *string `json:"message_id"`
	Error     *string `json:"error"`
}
