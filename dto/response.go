package dto

type AsciiFromUrlErrorResponse struct {
	AsciiFromUrlRequest
	Message string `json:"message"`
}
