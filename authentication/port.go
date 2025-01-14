package authentication

import (
	"auth-repo/types"
	"context"
)

type Service interface {
	Register(ctx context.Context, userInfo types.UserInfo) (*types.UserInfo, error)
}

type UserInfoRepo interface {
	RegisterUser(ctx context.Context, userInfo *types.UserInfo) (*types.UserInfo, error)
}

type EmailRepo interface {
	Send(subject, body, recipient string) error
}

type Cache interface {
	AddOtp(ctx context.Context, email string, otp string) error
}
