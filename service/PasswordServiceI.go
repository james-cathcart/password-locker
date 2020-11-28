package service

import . "pwlocker/model"

type PasswordRepositoryI interface {

	GetAll() []PasswordRecord
	LookupPassword(serviceName string) PasswordRecord
	GetByServiceName(serviceName string) (bool, PasswordRecord)
	Store(record PasswordRecord) (bool, PasswordRecord)
	Update()
}
