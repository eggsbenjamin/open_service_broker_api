package service

import "github.com/pkg/errors"

var (
	ErrServiceNotFound = errors.New("service not found")
	ErrPlanNotFound    = errors.New("plan not found")
)
