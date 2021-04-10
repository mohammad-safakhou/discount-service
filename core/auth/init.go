package auth

import "discount-service/transport"

type AuthenticateContext struct {
	*transport.ApplicationContext
}