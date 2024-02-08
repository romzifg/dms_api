package auth

type AuthRegisterDto struct {
	Username        string `gorm:"varchar" json:"username"`
	Email           string `gorm:"varchar" json:"email"`
	Password        string `gorm:"varchar" json:"password"`
	ConfirmPassword string `gorm:"varchar" json:"confirm_password"`
}

type AuthLoginDto struct {
	Email    string `gorm:"varchar" json:"email"`
	Password string `gorm:"varchar" json:"password"`
}