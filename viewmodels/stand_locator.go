package viewmodels

import ()

type StandLocator struct {
	Title  string
	Active string
}

func GetStandLocator() StandLocator {
	return StandLocator{
		Title:  "Stand Locator Page",
		Active: "stand_locator",
	}
}

func GetStandLocations(searchTerm string) []string {
	switch searchTerm {
	case "hong kong":
		return []string{"Sai Kung", "Central"}
	case "japan":
		return []string{"Osaka"}
	}
	return []string{}
}
