package hesta

import "encoding/xml"

type Respondents struct {
	XMLName    xml.Name   `xml:"Respondents"`
	Respondent Respondent `xml:"Respondent"`
}

type Respondent struct {
	EstBurNr              string         `xml:"EstBurNr"`
	Period                string         `xml:"Period"`
	AutomaticTransmission string         `xml:"AutomaticTransmission"`
	RespondentData        RespondentData `xml:"RespondentData"`
}

type RespondentData struct {
	Globals Globals `xml:"Globals"`
	Cursor  Cursor  `xml:"Cursor"`
}

type Globals struct {
	EstNbrRoom         int     `xml:"EstNbrRoom"`
	EstNbrBed          int     `xml:"EstNbrBed"`
	EstRoomForDisabled int     `xml:"EstRoomForDisabled"`
	EstContactName     string  `xml:"EstContactName"`
	EstContactPhone    string  `xml:"EstContactPhone"`
	EstContactEMail    string  `xml:"EstContactEMail"`
	EstNextClosingFrom string  `xml:"EstNextClosingFrom"`
	EstNextClosingTo   string  `xml:"EstNextClosingTo"`
	TotalArrivals      int     `xml:"TotalArrivals"`
	TotalNights        int     `xml:"TotalNights"`
	TotalOccupiedRooms int     `xml:"TotalOccupiedRooms"`
	EstAvgNightIncome  float64 `xml:"EstAvgNightIncome"`
	TotalOpenDays      int     `xml:"TotalOpenDays"`
}

type Cursor struct {
	ID  string `xml:"id,attr"`
	Row Rows   `xml:"Row"`
}

type Rows []Row

type Row struct {
	Country  string `xml:"country"`
	Arrivals int    `xml:"arrivals"`
	Nights   int    `xml:"nights"`
}
