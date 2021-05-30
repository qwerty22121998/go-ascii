package dto

type AsciiFromUrlRequest struct {
	Url  string `json:"url" query:"url" validate:"url"`
	Size uint   `json:"size" query:"size" validate:"required"`
}
