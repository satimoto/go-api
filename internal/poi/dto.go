package poi

type ElementDto struct {
	ID        string      `json:"id"`
	OsmJson   *OsmJsonDto `json:"osm_json"`
	Tags      TagsDto     `json:"tags"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt string      `json:"deleted_at"`
}

type OsmJsonDto struct {
	Lat    *float64   `json:"lat"`
	Lon    *float64   `json:"lon"`
	Bounds *BoundsDto `json:"bounds"`
	Tags   TagsDto    `json:"tags"`
	Type   string     `json:"type"`
}

type BoundsDto struct {
	MaxLat  float64 `json:"maxlat"`
	MaxLon  float64 `json:"maxlon"`
	MinLat  float64 `json:"minlat"`
	MinLon  float64 `json:"minlon"`
}

type TagsDto map[string]string
