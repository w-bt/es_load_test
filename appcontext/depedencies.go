package appcontext

import (
	userSvc "es_load_test/internal/services/item"
)

type ServerDependencies struct {
	Services Services
}

type Services struct {
	Item userSvc.Service
}
