package model

// User is a structure which defines the values returned by google.
type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	GiveName      string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Error         GoogleError
}
