package ports

import "context"

type AuthPort interface {
	Authenticate(ctx context.Context, accessToken string) (string, error)
	RoleAuthorize(ctx context.Context, accessToken string) error
}
