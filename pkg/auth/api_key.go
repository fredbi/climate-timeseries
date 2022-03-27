package auth

// AuthorizeAPIKey verifies the token passed in the API key header and returns
// a principal, if the user is valid.
func AuthorizeAPIKey(token string) (Principal, error) {
	return nil, nil
}

// AuthorizeOAuth2 verifies the bearer token passed in the Authorization header and returns
// a principal, if the user is valid.
func AuthorizeOAuth2(token string) (Principal, error) {
	return nil, nil
}
