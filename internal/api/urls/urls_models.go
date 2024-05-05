package urls

type URLRequest struct {
	Original string `json:"original"`
}

type URLResponse struct {
	Original  string `json:"original" gorm:"unique"`
	Short     string `json:"short"`
	CreatedAt int64  `json:"created_at"`
}
