package model

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenDto struct {
	UID          string `json:"id"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"isAdmin"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (t Token) ToDto(u *User) *TokenDto {
	return &TokenDto{
		UID:          u.ModelUID.ID,
		Email:        u.Email,
		IsAdmin:      u.IsAdmin,
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
	}
}
