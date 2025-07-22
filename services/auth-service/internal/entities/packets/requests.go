package packets

type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAdmin   bool   `json:"isAdmin"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	UserID    string `json:"userID" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty"`
}

type ForgotPasswordRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type RevokeTokenRequest struct {
	UserID string `json:"userID" validate:"required"`
}

type ResetPasswordRequest struct {
	TempToken string `json:"tempToken"`
	Password  string `json:"password" validate:"required"`
}

type ListUsersRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Search string `json:"search"`
}
