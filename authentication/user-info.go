package authentication

import (
	"auth-repo/types"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
)

func (svc *service) Register(ctx context.Context, userInfo types.UserInfo) (*types.UserInfo, error) {
	user, err := svc.userInfoRepo.RegisterUser(ctx, &userInfo)
	if err != nil {
		return nil, err
	}

	opt := random6Digit()
	err = svc.emailRepo.Send("Confirm register", fmt.Sprintf("Your code: %d", opt), user.Email)
	if err != nil {
		return nil, err
	}

	err = svc.cache.AddOtp(ctx, user.Email, fmt.Sprintf("%d", opt))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func random6Digit() int {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return 0
	}
	return int(n.Int64()) + 100000
}
