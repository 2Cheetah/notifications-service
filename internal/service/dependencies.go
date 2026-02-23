package service

import (
	"context"
)

type NovuClient interface {
	TriggerNotification(ctx context.Context, subscriberId string) (string, error)
}
