package customer

import "tech-challenge-auth/internal/canonical"

func translateToGRPC(customer canonical.Login) *Customer {
	return &Customer{
		Email:    customer.Email,
		Document: customer.Document,
		Password: customer.Password,
	}
}

func translateToCanonical(customer *Customer) canonical.Login {
	return canonical.Login{
		Document: customer.Document,
		Email:    customer.Document,
		Password: customer.Password,
	}
}
