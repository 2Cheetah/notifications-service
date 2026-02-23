package server

import "context"

type Service interface {
	TriggerNotification(ctx context.Context, subscriberId string) error
}
