package graph

import (
	"github.com/GSalise/lms/patron-service/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          *pgxpool.Pool
	PatronStore map[string]model.Patron
}
