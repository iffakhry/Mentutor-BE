package helper


// SUCCESS RESPONSE WITH DATA
func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data": data,
	}
}

//  SUCCESS RESPONSE NO DATA
func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

// FAILED RESEPONSE
func FailedResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message error": msg,
	}
}

