package pdf

import (
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/satimoto/go-api/internal/account"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/invoicerequest"
	"github.com/satimoto/go-datastore/pkg/location"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-datastore/pkg/user"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type PdfResolver struct {
	AccountResolver          account.AccountResolver
	InvoiceRequestRepository invoicerequest.InvoiceRequestRepository
	LocationRepository       location.LocationRepository
	NodeRepository           node.NodeRepository
	SessionRepository        session.SessionRepository
	TariffRepository         tariff.TariffRepository
	UserRepository           user.UserRepository
	blueColor                color.Color
	grayColor                color.Color
	fontSmall                float64
	fontFooter               float64
	fontTableHeader          float64
	fontText                 float64
}

func NewResolver(repositoryService *db.RepositoryService) *PdfResolver {
	return &PdfResolver{
		AccountResolver:          *account.NewResolver(repositoryService),
		InvoiceRequestRepository: invoicerequest.NewRepository(repositoryService),
		LocationRepository:       location.NewRepository(repositoryService),
		NodeRepository:           node.NewRepository(repositoryService),
		SessionRepository:        session.NewRepository(repositoryService),
		TariffRepository:         tariff.NewRepository(repositoryService),
		UserRepository:           user.NewRepository(repositoryService),
		blueColor:                getBlueColor(),
		grayColor:                getGrayColor(),
		fontSmall:                7,
		fontFooter:               8,
		fontTableHeader:          8,
		fontText:                 9,
	}
}

func getColor(r, g, b int) color.Color {
	return color.Color{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

func getBlueColor() color.Color {
	return getColor(0, 187, 255)
}

func getGrayColor() color.Color {
	return getColor(206, 212, 218)
}

func formatSatoshis(satoshis int64) string {
	p := message.NewPrinter(language.English)
	return strings.ReplaceAll(p.Sprintf("%d", satoshis), ",", " ")
}
