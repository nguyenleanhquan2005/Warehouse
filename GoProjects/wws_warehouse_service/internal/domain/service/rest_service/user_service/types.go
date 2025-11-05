package user_service

type UserInformationResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Birthdate []int  `json:"birthdate"`
	Email     string `json:"email"`
	VendorID  int64  `json:"vendor_id"`
}
