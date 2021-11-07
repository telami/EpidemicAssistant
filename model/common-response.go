package model

type Response struct {
	Results []struct {
		ProvinceShortName     string
		CurrentConfirmedCount int
		ConfirmedCount        int
		SuspectedCount        int
		CuredCount            int
		DeadCount             int
		Comment               string
		cities                []struct {
			CityName                 string
			CurrentConfirmedCount    int
			ConfirmedCount           int
			SuspectedCount           int
			CuredCount               int
			DeadCount                int
			HighDangerCount          int
			MidDangerCount           int
			LocationId               int
			CurrentConfirmedCountStr int
		}
	}
	Success bool
}
