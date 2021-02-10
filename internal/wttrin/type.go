package wttrin

type response struct {
	CurrentConditions []conditions `json:"current_condition"`
}

type conditions struct {
	FeelsLikeC          string           `json:"FeelsLikeC"`
	FeelsLikeF          string           `json:"FeelsLikeF"`
	CloudCover          string           `json:"cloudcover"`
	Humidity            string           `json:"humidity"`
	LocalObsDateTime    string           `json:"localObsDateTime"`
	ObservationTime     string           `json:"observation_time"`
	PrecipMM            string           `json:"precipMM"`
	Pressure            string           `json:"pressure"`
	TempC               string           `json:"temp_C"`
	TempF               string           `json:"temp_F"`
	UVIndex             string           `json:"uvIndex"`
	Visibility          string           `json:"visibility"`
	WeatherCode         string           `json:"weatherCode"`
	WeatherDescriptions []weatherDesc    `json:"weatherDesc"`
	WeatherIconURLS     []weatherIconURL `json:"weatherIconUrl"`
	WindDir16Point      string           `json:"winddir16Point"`
	WindDirDegree       string           `json:"winddirDegree"`
	WindSpeedKPH        string           `json:"windspeedKmph"`
	WindSpeedMPH        string           `json:"windspeedMiles"`
}

type weatherDesc struct {
	Value string `json:"value"`
}

type weatherIconURL struct {
	Value string `json:"value"`
}
