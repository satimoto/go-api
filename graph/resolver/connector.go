package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Identifier is the resolver for the identifier field.
func (r *connectorResolver) Identifier(ctx context.Context, obj *db.Connector) (*string, error) {
	return util.NullString(obj.Identifier)
}

// Evse is the resolver for the evse field.
func (r *connectorResolver) Evse(ctx context.Context, obj *db.Connector) (*db.Evse, error) {
	if evse, err := r.EvseRepository.GetEvse(ctx, obj.EvseID); err == nil {
		return &evse, nil
	}

	return nil, gqlerror.Errorf("Evse not found")
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

// GetConnector is the resolver for the getConnector field.
func (r *queryResolver) GetConnector(reqCtx context.Context, input graph.GetConnectorInput) (*db.Connector, error) {
	ctx := context.Background()
	
	if userID := middleware.GetUserID(reqCtx); userID != nil {
		if input.ID != nil {
			if connector, err := r.ConnectorRepository.GetConnector(ctx, *input.ID); err == nil {
				return &connector, nil
			}
		} else if input.Identifier != nil {
			dashRegex := regexp.MustCompile(`-`)
			identifier := strings.ToUpper(dashRegex.ReplaceAllString(*input.Identifier, "*"))

			if connector, err := r.ConnectorRepository.GetConnectorByIdentifier(ctx, dbUtil.SqlNullString(identifier)); err == nil {
				return &connector, nil
			}
		}

		return nil, gqlerror.Errorf("Connector not found")
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// Connector returns graph.ConnectorResolver implementation.
func (r *Resolver) Connector() graph.ConnectorResolver { return &connectorResolver{r} }

type connectorResolver struct{ *Resolver }
