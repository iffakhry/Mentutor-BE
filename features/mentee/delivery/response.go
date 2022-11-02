package delivery


type UpdateResponse struct {
	IdUser uint `json:"id_user"`
	Name string `json:"name"`
	Email string `json:"email"`
	Images string `json:"images"`
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data": data,
	}
}

func FailedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}