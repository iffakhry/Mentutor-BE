package delivery

import "be12/mentutor/features/user"

type UserResponse struct {
	ID     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Images string `json:"images" form:"images"`
	Class  string `json:"class" form:"class"`
	Role   string `json:"role"`
}

func toResponList(data []user.Core) []UserResponse {
	var dataAll []UserResponse
	for key := range data {
		dataAll = append(dataAll, FromCore(data[key]))
	}

	return dataAll
}

func FromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:     data.ID,
		Name:   data.Name,
		Images: data.Images,
		Role:   data.Role,
		Class:  data.Class,
	}
}
