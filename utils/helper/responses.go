package helper


// SUCCESS RESPONSE WITH DATA

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data": data,
	}
}

// FAILED RESEPONSE

func FailedResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message error": msg,
	}
}