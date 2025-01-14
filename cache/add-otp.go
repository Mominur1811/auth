package cache

import (
	"context"
	"time"
)

func (svc *cache) AddOtp(ctx context.Context, email string, otp string) error {
	err := svc.client.Set(ctx, email, otp, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
