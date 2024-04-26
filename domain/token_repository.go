package domain

import (
	"context"
)

type TokenRepository interface {
	CreateTokens(string, any) error
	TokenExists(context.Context, string) bool
}
