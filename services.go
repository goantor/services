package services

import "context"

type IService interface {
	Boot() error
	Shutdown(ctx context.Context) error
}
