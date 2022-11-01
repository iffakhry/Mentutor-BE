package delivery

import "be12/mentutor/features/user"

type UserResponse struct {
	ID     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Images string `json:"images" form:"images"`
	Class  string `json:"class" form:"class"`
	Role   string `json:"role"`
	// Bio    string `json:"bio" form:"bio"`
	// Email  string `json:"email" form:"email"`
	// Product []user.ProductCore `json:"product" form:"product"`
}

type ProductResponse struct {
	ID     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Images string `json:"images" form:"images"`
	Stock  int    `json:"stock" form:"stock"`
	Price  int    `json:"price" form:"price"`
}

// func FromUser(data user.ProductCore) ProductResponse {
// 	return ProductResponse{
// 		Name:   data.Name,
// 		Images: data.Images,
// 		Stock:  data.Stock,
// 		Price:  data.Price,
// 	}
// }

func FromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:     data.ID,
		Name:   data.Name,
		Images: data.Images,
		Role:   data.Role,
		Class:  data.Class,
		// Saldo:   int(data.Saldo),
		// Email:   data.Email,
		// Product: data.Product,
	}
}
