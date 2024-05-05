package shortener

type ShoternerRequest struct {
	Original string `json:"original"`
}

type ShortenerResponse struct {
	Original  string `json:"original" gorm:"unique"`
	Short     string `json:"short"`
	CreatedAt int64  `json:"created_at"`
}
