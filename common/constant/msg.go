package constant

var statusMessages = map[any]string{
	SUCCESS: "success",
	ERROR:   "error",
}

func GetMessage(statusCode any) string {
	if msg, exists := statusMessages[statusCode]; exists {
		return msg
	}
	return statusCode.(string)
}
