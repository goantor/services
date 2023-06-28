package services

import "context"

type IService interface {
	Boot() error
	ShutDown(ctx context.Context) error
}
