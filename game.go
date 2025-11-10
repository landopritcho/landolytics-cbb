package cbb

type Game struct {
	Date             string
	Time             string
	LongDate         string
	LongTitle        string
	ShortTitle       string
	EventTitle       string
	Venue            Venue
	HomeTeam         Team
	AwayTeam         Team
	IsNeutralSite    bool
	IsConferenceGame bool
	IsSpecialEvent   bool
}

type Team struct {
	School         string
	Mascot         string
	Abbreviation   string
	FullName       string
	ShortName      string
	PrimaryColor   string
	AlternateColor string
	Rank           string
}

type Venue struct {
	Name     string
	City     string
	State    string
	Location string
}

type InternalTeam struct {
	Id               string
	Slug             string
	Abbreviation     string
	DisplayName      string
	ShortDisplayName string
	Mascot           string
}
