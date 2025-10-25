package api

func ValidateToken(token string) bool {
	client := NewSourceCraftClient(token)
	body, err := client.DoRequest("GET", "/me/issues", nil)

	return body != nil && err == nil
}
