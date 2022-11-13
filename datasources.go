package core

type DataSource string

const (
	DATASOURCE_BILLINGREPORT DataSource = "billingreport"
	DATASOURCE_WEATHER       DataSource = "weather"
	DATASOURCE_INDOORCLIMATE DataSource = "indoorclimate"
	DATASOURCE_EXCHANGERATE  DataSource = "exchangerate"
	DATASOURCE_STRAVE        DataSource = "strava"
)
