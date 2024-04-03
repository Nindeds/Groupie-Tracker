package data_structure

type Relations struct {
	Relations []DatesLocation `json:"index"`
}
type DatesLocation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type SpotifyToken []struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}
type Concert struct {
	Location string
	Dates    string
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	RelationsUrl string   `json:"relations"`
	PastConcert  []Concert
	FuturConcert []Concert
}
