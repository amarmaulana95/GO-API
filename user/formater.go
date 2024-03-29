package user

type UserFormater struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Expired    string `json:"expired"`
}

func FormatUser(user User, token string) UserFormater {
	formatter := UserFormater{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Name,
		Token:      token,
	}
	return formatter
}
