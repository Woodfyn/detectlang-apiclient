package detectlanguage

import (
	"fmt"
	"time"
)

type languageResponce struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type accountStatusResponce struct {
	Date               string    `json:"date"`
	Requests           int       `json:"requests"`
	Bytes              int       `json:"bytes"`
	Plan               string    `json:"plan"`
	PlanExpires        time.Time `json:"plan_expires"`
	DailyRequestsLimit int       `json:"daily_requests_limit"`
	DailyBytesLimit    int       `json:"daily_bytes_limit"`
	Status             string    `json:"status"`
}

type detectResponce struct {
	Data DetectionData `json:"data"`
}

type DetectionData struct {
	Detections []Detection `json:"detections"`
}

type Detection struct {
	Language   string  `json:"language"`
	IsReliable bool    `json:"isReliable"`
	Confidence float64 `json:"confidence"`
}

func Info(l interface{}) string {
	switch v := l.(type) {
	case languageResponce:
		return fmt.Sprintf("[CODE] %s | [NAME] %s", v.Code, v.Name)
	case accountStatusResponce:
		return fmt.Sprintf("[DATE] %s | [REQUESTS] %d | [BYTES] %d | [PLAN] %s | [PLANEXPIRES] %s | [DAILYREQUESTSLIMIT] %d | [DAILYBYTESLIMIT] %d, [STATUS] %s",
			v.Date, v.Requests, v.Bytes, v.Plan, v.PlanExpires.Format(time.ANSIC), v.DailyRequestsLimit, v.DailyBytesLimit, v.Status)
	case detectResponce:
		var info string
		for _, d := range v.Data.Detections {
			info += fmt.Sprintf("[LANGUAGE] %s | [ISRELIABLE] %t | [CONFIDENCE] %.2f\n", d.Language, d.IsReliable, d.Confidence)
		}
		return info
	default:
		return "Not implemented"
	}
}
