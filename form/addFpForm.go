package form

// AddFpForm (add favorite page form)
type AddFpForm struct {
	Email     string  `json:"email,omitempty"`
	EmailType uint8   `json:"email_type,omitempty"`
	Url       string  `json:"url,omitempty"`
	Title     string  `json:"title,omitempty"`
	Score     float32 `json:"score,omitempty"`
}
