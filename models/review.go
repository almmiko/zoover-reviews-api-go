package models

import (
	"time"
)

type Review struct {
	Parents []struct {
		ID string `json:"id"`
	} `json:"parents"`
	ID           string `json:"id"`
	TraveledWith string `json:"traveledWith"`
	EntryDate    int64  `json:"entryDate"`
	TravelDate   int64  `json:"travelDate"`
	Ratings      struct {
		General struct {
			General float32 `json:"general"`
		} `json:"general"`
		Aspects struct {
			Location        int `json:"location"`
			Service         int `json:"service"`
			PriceQuality    int `json:"priceQuality"`
			Food            int `json:"food"`
			Room            int `json:"room"`
			ChildFriendly   int `json:"childFriendly"`
			Interior        int `json:"interior"`
			Size            int `json:"size"`
			Activities      int `json:"activities"`
			Restaurants     int `json:"restaurants"`
			SanitaryState   int `json:"sanitaryState"`
			Accessibility   int `json:"accessibility"`
			Nightlife       int `json:"nightlife"`
			Culture         int `json:"culture"`
			Surrounding     int `json:"surrounding"`
			Atmosphere      int `json:"atmosphere"`
			NoviceSkiArea   int `json:"noviceSkiArea"`
			AdvancedSkiArea int `json:"advancedSkiArea"`
			ApresSki        int `json:"apresSki"`
			Beach           int `json:"beach"`
			Entertainment   int `json:"entertainment"`
			Environmental   int `json:"environmental"`
			Pool            int `json:"pool"`
			Terrace         int `json:"terrace"`
		} `json:"aspects"`
	} `json:"ratings"`
	Titles struct {
		Nl string `json:"nl"`
	} `json:"titles"`
	Texts struct {
		Nl string `json:"nl"`
	} `json:"texts"`
	User   string `json:"user"`
	Locale string `json:"locale"`
}

func (r *Review) ReviewWeight() float32 {
	const OlderThanYears = 5

	reviewYear := time.Unix(0, r.EntryDate * 1000000).Year()
	currentYear := time.Now().Year()

	if currentYear - OlderThanYears > reviewYear {
		return 0.5
	}

	return 1 - float32(currentYear - reviewYear) * 0.1
}