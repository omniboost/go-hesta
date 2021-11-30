package hesta

import "encoding/xml"

type Respondents struct {
	XMLName    xml.Name `xml:"Respondents"`
	Respondent struct {
		EstBurNr              string `xml:"EstBurNr"`
		Period                string `xml:"Period"`
		AutomaticTransmission string `xml:"AutomaticTransmission"`
		RespondentData        struct {
			Globals struct {
				EstNbrRoom         string `xml:"EstNbrRoom"`
				EstNbrBed          string `xml:"EstNbrBed"`
				EstRoomForDisabled string `xml:"EstRoomForDisabled"`
				EstContactName     string `xml:"EstContactName"`
				EstContactPhone    string `xml:"EstContactPhone"`
				EstContactEMail    string `xml:"EstContactEMail"`
				EstNextClosingFrom string `xml:"EstNextClosingFrom"`
				EstNextClosingTo   string `xml:"EstNextClosingTo"`
				TotalArrivals      string `xml:"TotalArrivals"`
				TotalNights        string `xml:"TotalNights"`
				TotalOccupiedRooms string `xml:"TotalOccupiedRooms"`
				EstAvgNightIncome  string `xml:"EstAvgNightIncome"`
				TotalOpenDays      string `xml:"TotalOpenDays"`
			} `xml:"Globals"`
			Cursor struct {
				ID  string `xml:"id,attr"`
				Row []struct {
					Text     string `xml:",chardata"`
					Country  string `xml:"country"`
					Arrivals string `xml:"arrivals"`
					Nights   string `xml:"nights"`
				} `xml:"Row"`
			} `xml:"Cursor"`
		} `xml:"RespondentData"`
	} `xml:"Respondent"`
}
