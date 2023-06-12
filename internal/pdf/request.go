package pdf

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/template"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
)

func (r *PdfResolver) GetInvoicePdf(rw http.ResponseWriter, request *http.Request) {
	reqCtx := request.Context()

	if user := middleware.GetUser(reqCtx, r.UserRepository); user != nil {
		uid := chi.URLParam(request, "uid")
		ctx := context.Background()

		if session, err := r.SessionRepository.GetSessionByUid(ctx, uid); err == nil && (user.ID == session.UserID || user.IsAdmin) {
			if bytes, err := r.invoicePdfBytes(ctx, session); err == nil {
				contentDisposition := fmt.Sprintf("attachment; filename=\"%s.pdf\"", uid)

				rw.WriteHeader(http.StatusOK)
				rw.Header().Set("Content-Disposition", contentDisposition)
				rw.Header().Set("Content-Type", "application/octet-stream")
				rw.Write(bytes)
			}
		}
	}

	rw.WriteHeader(http.StatusNotFound)
}

func (r *PdfResolver) invoicePdfBytes(ctx context.Context, session db.Session) ([]byte, error) {
	currency := session.Currency
	user, err := r.UserRepository.GetUser(ctx, session.UserID)

	if err != nil {
		metrics.RecordError("API069", "Error getting user", err)
		log.Printf("API069: Uid=%v, UserID=%v", session.Uid, session.UserID)
		return nil, errors.New("error getting user")
	}

	location, err := r.LocationRepository.GetLocation(ctx, session.LocationID)

	if err != nil {
		metrics.RecordError("API070", "Error getting location", err)
		log.Printf("API070: Uid=%v, LocationID=%#v", session.Uid, session.LocationID)
		return nil, errors.New("error getting location")
	}

	evse, err := r.LocationRepository.GetEvse(ctx, session.EvseID)

	if err != nil {
		metrics.RecordError("API071", "Error getting evse", err)
		log.Printf("API071: Uid=%v, EvseID=%#v", session.Uid, session.EvseID)
		return nil, errors.New("error getting evse")
	}

	if connector, err := r.LocationRepository.GetConnector(ctx, session.ConnectorID); err != nil && connector.TariffID.Valid {
		if tariff, err := r.TariffRepository.GetTariffByUid(ctx, connector.TariffID.String); err != nil {
			currency = tariff.Currency
		}
	}

	var currencyText, currencyDecimalFormat = "Fiat", "%.2f"

	if currency, err := r.AccountResolver.Repository.GetCurrencyByCode(ctx, currency); err == nil {
		currencyText = currency.Name
		currencyDecimalFormat = fmt.Sprintf("%%.%df", currency.Decimals)
	}

	taxPercent := r.AccountResolver.GetTaxPercentByCountry(ctx, location.Country, dbUtil.GetEnvFloat64("DEFAULT_TAX_PERCENT", 19))
	sessionInvoices, err := r.SessionRepository.ListSessionInvoicesBySessionID(ctx, session.ID)

	if err != nil {
		metrics.RecordError("API072", "Error listing session invoices", err)
		log.Printf("API072: Uid=%v", session.Uid)
		return nil, errors.New("error listing session invoices")
	}

	logoBytes, err := template.ReadFile("image/logo.png")

	if err != nil {
		metrics.RecordError("API073", "Error reading image file", err)
		log.Printf("API073: Uid=%v", session.Uid)
		return nil, errors.New("error reading image file")
	}

	logoBase64 := base64.StdEncoding.EncodeToString(logoBytes)

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetDefaultFontFamily(consts.Helvetica)
	m.SetPageMargins(20, 10, 20)

	m.RegisterHeader(func() {
		m.Row(30, func() {
			m.Col(9, func() {
				m.Text("Invoice", props.Text{
					Top:   5,
					Size:  40,
					Style: consts.Bold,
					Align: consts.Left,
					Left:  1.5,
				})
			})
			m.Col(3, func() {
				_ = m.Base64Image(logoBase64, consts.Png, props.Rect{})
			})
		})
	})

	m.RegisterFooter(func() {
		m.Row(12, func() {
			m.Col(4, func() {
				m.Text("Satimoto", props.Text{
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Left,
					Left:  1.5,
				})
				m.Text("Postfach 21 11 05", props.Text{
					Top:   4,
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Left,
					Left:  1.5,
				})
				m.Text("04122 Leipzig", props.Text{
					Top:   8,
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Left,
					Left:  1.5,
				})
				m.Text("Germany", props.Text{
					Top:   12,
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Left,
					Left:  1.5,
				})
			})
			m.Col(4, func() {
				m.Text("https://satimoto.com", props.Text{
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Center,
					Left:  1.5,
				})
				m.Text("hello@satimoto.com", props.Text{
					Top:   4,
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Center,
					Left:  1.5,
				})
			})

			m.Col(4, func() {
				m.Text("Radicle UG", props.Text{
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Right,
					Left:  1.5,
				})
				m.Text("HRB No.: 231270", props.Text{
					Top:   4,
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Right,
					Left:  1.5,
				})
				m.Text("USt-IdNr.: DE310932213", props.Text{
					Top:   8,
					Size:  r.fontFooter,
					Color: r.grayColor,
					Align: consts.Right,
					Left:  1.5,
				})
			})
		})
	})

	address := util.DeleteEmpty([]string{user.PostalCode.String, user.City.String})
	r.renderInfoLine(m, 5, user.Name.String, "", "Satimoto")
	r.renderInfoLine(m, 5, user.Address.String, "", "Postfach 21 11 05")
	r.renderInfoLine(m, 5, strings.Join(address, " "), "", "04122 Leipzig")
	r.renderInfoLine(m, 5, "", "", "Germany")
	r.renderInfoLine(m, 5, "", "Invoice No.", fmt.Sprintf("SAT/%d/%s", session.ID, session.StartDatetime.Format("2006")))
	r.renderInfoLine(m, 5, "", "Invoice Date", session.StartDatetime.Format("02.01.2006"))
	r.renderInfoLine(m, 5, "", "Payment Date", session.StartDatetime.Format("02.01.2006"))

	address = util.DeleteEmpty([]string{location.Address, fmt.Sprintf("%s %s", location.PostalCode, location.City)})
	r.renderLine(m, 6, fmt.Sprintf("%s / %s", session.StartDatetime.Format("02.01.2006"), dbUtil.DefaultString(evse.EvseID, dbUtil.DefaultString(evse.Identifier, evse.Uid))), consts.Bold)
	r.renderLine(m, 6, strings.Join(address, ", "), consts.Normal)
	r.renderLine(m, 10, location.Name.String, consts.Normal)

	r.renderTableHeader(m, "Payments", currencyText, "Satoshis")

	var commissionFiat, taxFiat, totalFiat = 0.0, 0.0, 0.0
	var commissionMsat, taxMsat, totalMsat = int64(0), int64(0), int64(0)

	for _, sessionInvoice := range sessionInvoices {
		title := fmt.Sprintf("%s...", sessionInvoice.PaymentRequest[:32])
		amountFiat := fmt.Sprintf(currencyDecimalFormat, sessionInvoice.PriceFiat)
		amountMsat := formatSatoshis(sessionInvoice.PriceMsat)

		r.renderTableRow(m, title, amountFiat, amountMsat)

		commissionFiat += sessionInvoice.CommissionFiat
		commissionMsat += sessionInvoice.CommissionMsat
		taxFiat += sessionInvoice.TaxFiat
		taxMsat += sessionInvoice.TaxMsat
		totalFiat += sessionInvoice.TotalFiat
		totalMsat += sessionInvoice.TotalMsat
	}

	invoiceRequests, err := r.InvoiceRequestRepository.ListInvoiceRequestsBySessionID(ctx, dbUtil.SqlNullInt64(session.ID))

	if err != nil {
		metrics.RecordError("API067", "Error listing invoice requests", err)
		log.Printf("API067: Uid=%v", session.Uid)
		return nil, errors.New("error listing invoice requests")
	}


	if len(invoiceRequests) > 0 {
		r.renderTableHeader(m, "Rebates", currencyText, "Satoshis")

		for _, invoiceRequest := range invoiceRequests {
			title := fmt.Sprintf("%s...", invoiceRequest.PaymentRequest.String[:32])
			amountFiat := fmt.Sprintf("- "+currencyDecimalFormat, invoiceRequest.PriceFiat.Float64)
			amountMsat := fmt.Sprintf("- %s", formatSatoshis(invoiceRequest.PriceMsat.Int64))
	
			r.renderTableRow(m, title, amountFiat, amountMsat)
	
			commissionFiat -= invoiceRequest.CommissionFiat.Float64
			commissionMsat -= invoiceRequest.CommissionMsat.Int64
			taxFiat -= invoiceRequest.TaxFiat.Float64
			taxMsat -= invoiceRequest.TaxMsat.Int64
			totalFiat -= invoiceRequest.TotalFiat
			totalMsat -= invoiceRequest.TotalMsat
		}
	}

	subtotalFiat := totalFiat - taxFiat
	subtotalMsat := totalMsat - taxMsat

	r.renderSubtotal(m, fmt.Sprintf("Commission (%.1f%%)", user.CommissionPercent), fmt.Sprintf(currencyDecimalFormat, commissionFiat), formatSatoshis(commissionMsat), consts.Normal)
	r.renderSubtotal(m, "Total excl. VAT", fmt.Sprintf(currencyDecimalFormat, subtotalFiat), formatSatoshis(subtotalMsat), consts.Bold)
	r.renderSubtotal(m, fmt.Sprintf("VAT (%.1f%%)", taxPercent), fmt.Sprintf(currencyDecimalFormat, taxFiat), formatSatoshis(taxMsat), consts.Normal)
	r.renderSubtotal(m, "Total", fmt.Sprintf(currencyDecimalFormat, totalFiat), formatSatoshis(totalMsat), consts.Bold)

	outputBuffer, err := m.Output()

	if err != nil {
		metrics.RecordError("API040", "Error encoding pdf", err)
		log.Printf("API040: Uid=%v", session.Uid)
		return nil, errors.New("error encoding pdf")
	}

	return outputBuffer.Bytes(), nil
}

func (r *PdfResolver) renderLine(m pdf.Maroto, rowHeight float64, text string, style consts.Style) {
	m.Row(rowHeight, func() {
		m.Col(12, func() {
			m.Text(text, props.Text{
				Size:  r.fontText,
				Style: style,
				Align: consts.Left,
				Left:  1.5,
			})
		})
	})
}

func (r *PdfResolver) renderInfoLine(m pdf.Maroto, rowHeight float64, leftText, rightLabel, rightText string) {
	m.Row(rowHeight, func() {
		m.Col(3, func() {
			m.Text(leftText, props.Text{
				Size:  r.fontText,
				Align: consts.Left,
			})
		})
		m.ColSpace(4)
		m.Col(2, func() {
			m.Text(rightLabel, props.Text{
				Top:   0.8,
				Size:  r.fontSmall,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(rightText, props.Text{
				Size:  r.fontText,
				Align: consts.Left,
			})
		})
	})
}

func (r *PdfResolver) renderTableHeader(m pdf.Maroto, title, titleFiat, titleMsat string) {
	rowHeight := 6.0

	if len(titleFiat) > 16 {
		rowHeight = 9.0
	}

	m.SetBackgroundColor(r.blueColor)

	m.Row(rowHeight, func() {
		m.Col(6, func() {
			m.Text(title, props.Text{
				Top:   1,
				Size:  r.fontTableHeader,
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1.5,
			})
		})
		m.Col(2, func() {
			m.Text(titleFiat, props.Text{
				Top:   1,
				Size:  r.fontTableHeader,
				Style: consts.Bold,
				Align: consts.Right,
				Right: 1.5,
			})
		})
		m.ColSpace(1)
		m.Col(3, func() {
			m.Text(titleMsat, props.Text{
				Top:   1,
				Size:  r.fontTableHeader,
				Style: consts.Bold,
				Align: consts.Right,
				Right: 9,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())
}

func (r *PdfResolver) renderTableRow(m pdf.Maroto, title, amountFiat, amountMsat string) {
	m.Row(7, func() {
		m.Col(6, func() {
			m.Text(title, props.Text{
				Top:    1,
				Size:   r.fontText,
				Family: consts.Courier,
				Align:  consts.Left,
				Left:   1.5,
			})
		})
		m.Col(2, func() {
			m.Text(amountFiat, props.Text{
				Top:    1,
				Size:   r.fontText,
				Family: consts.Courier,
				Align:  consts.Right,
				Right:  1.5,
			})
		})
		m.ColSpace(1)
		m.Col(3, func() {
			splitMsat := strings.Split(amountMsat, " ")
			satsLen := len(splitMsat) - 1

			if satsLen > 0 {
				m.Text(strings.Join(splitMsat[0:satsLen], " "), props.Text{
					Top:    1,
					Size:   r.fontText,
					Family: consts.Courier,
					Align:  consts.Right,
					Right:  9,
				})
			}
			m.Text(" "+splitMsat[satsLen], props.Text{
				Top:    1,
				Size:   r.fontText,
				Family: consts.Courier,
				Color:  r.grayColor,
				Align:  consts.Right,
				Right:  1.5,
			})
		})
	})
}

func (r *PdfResolver) renderSubtotal(m pdf.Maroto, title, amountFiat, amountMsat string, style consts.Style) {
	m.Row(6, func() {
		m.ColSpace(4)
		m.Col(3, func() {
			m.Text(title, props.Text{
				Top:   3,
				Size:  r.fontText,
				Style: style,
				Align: consts.Right,
			})
		})
		m.Col(1, func() {
			m.Text(amountFiat, props.Text{
				Top:    3,
				Size:   r.fontText,
				Family: consts.Courier,
				Style:  style,
				Align:  consts.Right,
				Right:  1.5,
			})
		})
		m.ColSpace(1)
		m.Col(3, func() {
			splitMsat := strings.Split(amountMsat, " ")
			satsLen := len(splitMsat) - 1

			if satsLen > 0 {
				m.Text(strings.Join(splitMsat[0:satsLen], " "), props.Text{
					Top:    3,
					Size:   r.fontText,
					Family: consts.Courier,
					Style:  style,
					Align:  consts.Right,
					Right:  9,
				})
			}
			m.Text(" "+splitMsat[satsLen], props.Text{
				Top:    3,
				Size:   r.fontText,
				Family: consts.Courier,
				Color:  r.grayColor,
				Style:  style,
				Align:  consts.Right,
				Right:  1.5,
			})
		})
	})
}
