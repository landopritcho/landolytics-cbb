package espn

type ParsedGame struct {
	Date             string
	LongDate         string
	Time             string
	LongTitle        string
	ShortTitle       string
	VenueName        string
	VenueCity        string
	VenueState       string
	VenueLocation    string
	EventTitle       string
	IsNeutralSite    bool
	IsConferenceGame bool
	IsSpecialEvent   bool
	HomeTeam         ParsedTeam
	AwayTeam         ParsedTeam
}

type ParsedTeam struct {
	School         string
	Mascot         string
	Abbreviation   string
	FullName       string
	ShortName      string
	PrimaryColor   string
	AlternateColor string
	Rank           string
}
