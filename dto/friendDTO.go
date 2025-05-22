package dto

type friendListResponse struct {
	Age    float64 `json:"age"`
	Name   string  `json:"name"`
	UserID float64 `json:"user_id"`
}

func GetFriendListResponseDTO(age float64, name string, user_id float64) *friendListResponse {
	return &friendListResponse{
		Age:    age,
		Name:   name,
		UserID: user_id,
	}
}
