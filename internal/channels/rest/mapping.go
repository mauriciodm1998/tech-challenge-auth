package rest

import "tech-challenge-auth/internal/canonical"

func (c *LoginRequest) toCanonical() canonical.Login {
	return canonical.Login{
		Document: c.Document,
		Email:    c.Email,
		Password: c.Password,
	}
}
