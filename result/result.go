package result

import "encoding/json"

type Result struct {
	Start            string        `json:"realtime_start"`
	End              string        `json:"realtime_end"`
	ObservationStart string        `json:"observation_start"`
	ObservationEnd   string        `json:"observation_end"`
	Units            string        `json:"units"`
	OutputType       int           `json:"output_type"`
	FileType         string        `json:"file_type"`
	OrderBy          string        `json:"order_by"`
	SortOrder        string        `json:"sort_order"`
	Count            int           `json:"count"`
	Offset           int           `json:"offset"`
	Limit            int           `json:"limit"`
	Series           []Series      `json:"seriess,omitempty"`
	Observations     []Observation `json:"observations,omitempty"`
	Releases         []Release     `json:"releases,omitempty"`
	Categories       []Category    `json:"categories,omitempty"`
}

func (r *Result) JSON() (string, error) {
	jsonData, err := json.Marshal(r)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (r *Result) PrettyJSON() (string, error) {
	jsonData, err := json.MarshalIndent(r, "", " ")

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

type Series struct {
	ID                      string `json:"id" csv:"id"`
	RealtimeStart           string `json:"realtime_start" csv:"realtime_start"`
	RealtimeEnd             string `json:"realtime_end" csv:"realtime_end"`
	Title                   string `json:"title" csv:"title"`
	ObservationStart        string `json:"observation_start" csv:"observation_start"`
	ObservationEnd          string `json:"observation_end" csv:"observation_end"`
	Frequency               string `json:"frequency" csv:"frequency"`
	FrequencyShort          string `json:"frequency_short" csv:"frequency_short"`
	Units                   string `json:"units" csv:"units"`
	UnitsShort              string `json:"units_short" csv:"units_short"`
	SeasonalAdjustment      string `json:"seasonal_adjustment" csv:"seasonal_adjustment"`
	SeasonalAdjustmentShort string `json:"seasonal_adjustment_short" csv:"seasonal_adjustment_short"`
	LastUpdated             string `json:"last_updated" csv:"last_updated"`
	Popularity              int    `json:"popularity" csv:"popularity"`
	Notes                   string `json:"notes"`
}

type Observation struct {
	Date          string `json:"date"`
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime_end"`
	Value         string `json:"value"`
}

type Release struct {
	ID            int
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime_end"`
	Name          string `json:"name"`
	PressRelease  string `json:"press_release"`
	Link          string `json:"link"`
}

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
}
