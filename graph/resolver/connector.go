package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
)

// Identifier is the resolver for the identifier field.
func (r *connectorResolver) Identifier(ctx context.Context, obj *db.Connector) (*string, error) {
	return util.NullString(obj.Identifier)
}

// Standard is the resolver for the standard field.
func (r *connectorResolver) Standard(ctx context.Context, obj *db.Connector) (string, error) {
	return string(obj.Standard), nil
}

// Format is the resolver for the format field.
func (r *connectorResolver) Format(ctx context.Context, obj *db.Connector) (string, error) {
	return string(obj.Format), nil
}

// PowerType is the resolver for the powerType field.
func (r *connectorResolver) PowerType(ctx context.Context, obj *db.Connector) (string, error) {
	return string(obj.PowerType), nil
}

// TariffID is the resolver for the tariffId field.
func (r *connectorResolver) TariffID(ctx context.Context, obj *db.Connector) (*string, error) {
	return util.NullString(obj.TariffID)
}

// Tariff is the resolver for the tariff field.
func (r *connectorResolver) Tariff(ctx context.Context, obj *db.Connector) (*db.Tariff, error) {
	if obj.TariffID.Valid {
		if tariff, err := r.TariffRepository.GetTariffByUid(ctx, obj.TariffID.String); err == nil {
			return &tariff, nil
		}
	}

	return nil, nil
}

// TermsAndConditions is the resolver for the termsAndConditions field.
func (r *connectorResolver) TermsAndConditions(ctx context.Context, obj *db.Connector) (*string, error) {
	return util.NullString(obj.TermsAndConditions)
}

// LastUpdated is the resolver for the lastUpdated field.
func (r *connectorResolver) LastUpdated(ctx context.Context, obj *db.Connector) (string, error) {
	return obj.LastUpdated.Format(time.RFC3339), nil
}

// Connector returns graph.ConnectorResolver implementation.
func (r *Resolver) Connector() graph.ConnectorResolver { return &connectorResolver{r} }

type connectorResolver struct{ *Resolver }
