package api

func ValidateToken(token string) bool {
	client := NewSourceCraftClient(token)
	body, err := client.DoRequest("GET", "/me/issues")

	return body != nil && err == nil
}
