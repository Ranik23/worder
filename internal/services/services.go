package services

import (
	analytics_service "Worder/internal/services/analytics"
	corrector_service "Worder/internal/services/corrector"
)


type AppServices struct {
	Corrector corrector_service.Corrector
	Analytic  analytics_service.Analytic
}