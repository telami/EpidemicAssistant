package model

type RiskGrade struct {
	Data struct {
		End_update_time string
		Hcount          int
		Mcount          int
		Middlelist      []struct {
			Type       string
			Province   string
			City       string
			County     string
			Area_name  string
			Communitys []string
		}
		Highlist []struct {
			Type       string
			Province   string
			City       string
			County     string
			Area_name  string
			Communitys []string
		}
	}
	Code int
	Msg  string
}

func IsSuccessful(grade RiskGrade) bool {
	return grade.Code == 0
}
