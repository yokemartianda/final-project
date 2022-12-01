package http_request

type RequestCustomer struct {
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Alamat      string `json:"alamat"`
	PhoneNumber string `json:"phone_number"`
	CreatedTime string `json:"created_time"`
}
