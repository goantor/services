package services

type IService interface {
	Boot() error
	ShutDown() error
}
