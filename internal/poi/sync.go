package poi

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/paulmach/orb"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/geom"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
)

var (
	PHYSICAL_TAG_KEYS = []string{"amenity", "shop", "tourism", "office", "place", "barrier", "highway", "historic", "leisure", "man_made", "natural", "religion"}
	ALL_TAG_KEYS      = append(PHYSICAL_TAG_KEYS, "sport", "station", "cuisine", "building")
)

func (r *PoiResolver) SyncronizePois() {
	ctx := context.Background()

	r.syncronizeBtcMapPois(ctx)
}

func (r *PoiResolver) syncronizeBtcMapPois(ctx context.Context) {
	btcMapUrl := "https://api.btcmap.org/v2/elements"
	limit := 500
	requestUrl, err := url.Parse(btcMapUrl)

	if err != nil {
		metrics.RecordError("API058", "Error parsing url", err)
		log.Printf("API058: Url=%v", btcMapUrl)
		return
	}

	query := requestUrl.Query()
	query.Set("limit", strconv.Itoa(limit))

	if poi, err := r.Repository.GetPoiByLastUpdated(ctx); err == nil {
		query.Set("updated_since", poi.LastUpdated.Format(time.RFC3339))
	}

	for {
		requestUrl.RawQuery = query.Encode()
		request, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)

		if err != nil {
			metrics.RecordError("API059", "Error making request", err)
			util.LogHttpRequest("API059", requestUrl.String(), request, false)
			return
		}

		response, err := r.HTTPRequester.Do(request)

		if err != nil {
			metrics.RecordError("API060", "Error getting response", err)
			util.LogHttpResponse("API060", requestUrl.String(), response, true)
			return
		}

		elementsDto, err := UnmarshalDto(response.Body)
		defer response.Body.Close()

		if err != nil {
			metrics.RecordError("API061", "Error unmarshaling", err)
			util.LogHttpResponse("API061", requestUrl.String(), response, true)
			return
		}

		for _, elementDto := range elementsDto {
			r.processElement(ctx, elementDto)
			query.Set("updated_since", elementDto.UpdatedAt)
		}

		log.Printf("Page limit=%v since=%s count=%v", limit, query.Get("updated_since"), len(elementsDto))

		if len(elementsDto) < limit {
			break
		}
	}
}

func (r *PoiResolver) processElement(ctx context.Context, elementDto *ElementDto) {
	osmJsonDto := elementDto.OsmJson

	if osmJsonDto != nil && osmJsonDto.Type == "node" {
		tagsDto := osmJsonDto.Tags

		if name, ok := tagsDto["name"]; ok {
			tagsDto := osmJsonDto.Tags

			if len(elementDto.DeletedAt) > 0 {
				r.Repository.DeletePoiByUid(ctx, elementDto.ID)
			} else {
				poi, err := r.Repository.GetPoiByUid(ctx, elementDto.ID)
				tagKey, tagValue := r.getTag(tagsDto)

				if err == nil {
					updatePoiByUidParams := param.NewUpdatePoiByUidParams(poi)
					updatePoiByUidParams.Name = name
					updatePoiByUidParams.Description = util.SqlNullString(tagsDto["description"])
					updatePoiByUidParams.Geom = r.getGeom(osmJsonDto)
					updatePoiByUidParams.Address = util.SqlNullString(r.getAddress(tagsDto))
					updatePoiByUidParams.City = util.SqlNullString(tagsDto["addr:city"])
					updatePoiByUidParams.PostalCode = util.SqlNullString(tagsDto["addr:postcode"])
					updatePoiByUidParams.TagKey = tagKey
					updatePoiByUidParams.TagValue = tagValue
					updatePoiByUidParams.PaymentOnChain = r.getBool(tagsDto["payment:onchain"]) || r.getBool(tagsDto["payment:bitcoin"])
					updatePoiByUidParams.PaymentLn = r.getBool(tagsDto["payment:lightning"])
					updatePoiByUidParams.PaymentLnTap = r.getBool(tagsDto["payment:lightning_contactless"])
					updatePoiByUidParams.OpeningTimes = util.SqlNullString(tagsDto["opening_hours"])
					updatePoiByUidParams.Phone = util.SqlNullString(tagsDto["phone"])
					updatePoiByUidParams.Website = util.SqlNullString(tagsDto["website"])
					updatePoiByUidParams.LastUpdated = parseTime(elementDto.UpdatedAt, time.Now())

					updatedPoi, err := r.Repository.UpdatePoiByUid(ctx, updatePoiByUidParams)

					if err != nil {
						metrics.RecordError("API062", "Error updating poi", err)
						log.Printf("API062: Params=%#v", updatePoiByUidParams)
					}

					poi = updatedPoi
				} else {
					createPoiParams := NewCreatePoiParams(elementDto)
					createPoiParams.Name = name
					createPoiParams.Description = util.SqlNullString(tagsDto["description"])
					createPoiParams.Geom = r.getGeom(osmJsonDto)
					createPoiParams.Address = util.SqlNullString(r.getAddress(tagsDto))
					createPoiParams.City = util.SqlNullString(tagsDto["addr:city"])
					createPoiParams.PostalCode = util.SqlNullString(tagsDto["addr:postcode"])
					createPoiParams.TagKey = tagKey
					createPoiParams.TagValue = tagValue
					createPoiParams.PaymentOnChain = r.getBool(tagsDto["payment:onchain"]) || r.getBool(tagsDto["payment:bitcoin"])
					createPoiParams.PaymentLn = r.getBool(tagsDto["payment:lightning"])
					createPoiParams.PaymentLnTap = r.getBool(tagsDto["payment:lightning_contactless"])
					createPoiParams.OpeningTimes = util.SqlNullString(tagsDto["opening_hours"])
					createPoiParams.Phone = util.SqlNullString(tagsDto["phone"])
					createPoiParams.Website = util.SqlNullString(tagsDto["website"])

					poi, err = r.Repository.CreatePoi(ctx, createPoiParams)

					if err != nil {
						metrics.RecordError("API063", "Error creating poi", err)
						log.Printf("API063: Params=%#v", createPoiParams)
					}
				}

				r.processTags(ctx, poi.ID, tagsDto)
			}
		}
	}
}

func (r *PoiResolver) getAddress(tagsDto TagsDto) string {
	addressParts := []string{}

	if housenumber, ok := tagsDto["addr:housenumber"]; ok {
		addressParts = append(addressParts, housenumber)
	}

	if street, ok := tagsDto["addr:street"]; ok {
		addressParts = append(addressParts, street)
	}

	return strings.Join(addressParts, " ")
}

func (r *PoiResolver) getBool(str string) bool {
	return str == "yes"
}

func (r *PoiResolver) getGeom(osmJson *OsmJsonDto) geom.Geometry4326 {
	point := orb.Point{osmJson.Lon, osmJson.Lat}

	return geom.Geometry4326{
		Coordinates: point,
		Type:        point.GeoJSONType(),
	}
}

func (r *PoiResolver) getTag(tagsDto TagsDto) (string, string) {
	for _, key := range PHYSICAL_TAG_KEYS {
		if value, ok := tagsDto[key]; ok {
			return key, value
		}
	}

	return "other", "other"
}

func (r *PoiResolver) processTagValue(value string) string {
	processedValue := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(value, "-", "_"), " ", "_"), ",", ";"))
	processedSplitValues := []string{}

	for _, splitValue := range strings.Split(processedValue, ";") {
		processedSplitValue := strings.Trim(splitValue, "_")

		if len(processedSplitValue) > 0 {
			processedSplitValues = append(processedSplitValues, processedSplitValue)
		}
	}

	return strings.Join(processedSplitValues, ";")
}

func (r *PoiResolver) processTags(ctx context.Context, poiID int64, tagsDto TagsDto) {
	r.Repository.UnsetPoiTags(ctx, poiID)

	for _, key := range ALL_TAG_KEYS {
		if value, ok := tagsDto[key]; ok {
			processedValue := r.processTagValue(value)
			getTagByKeyValueParams := db.GetTagByKeyValueParams{
				Key:   key,
				Value: processedValue,
			}

			tag, err := r.Repository.GetTagByKeyValue(ctx, getTagByKeyValueParams)

			if err != nil {
				createTagParams := db.CreateTagParams{
					Key:   key,
					Value: processedValue,
				}

				createdTag, err := r.Repository.CreateTag(ctx, createTagParams)

				if err != nil {
					metrics.RecordError("API064", "Error creating tag", err)
					log.Printf("API064: Params=%#v", createTagParams)
				}

				tag = createdTag
			}

			setPoiTagParams := db.SetPoiTagParams{
				PoiID: poiID,
				TagID: tag.ID,
			}

			err = r.Repository.SetPoiTag(ctx, setPoiTagParams)

			if err != nil {
				metrics.RecordError("API065", "Error setting poi tag", err)
				log.Printf("API065: Params=%#v", setPoiTagParams)
			}
		}
	}
}
