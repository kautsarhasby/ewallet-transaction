package services

import "kautsarhasby/ewallet-transaction/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepo
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "service healthy", nil
}
