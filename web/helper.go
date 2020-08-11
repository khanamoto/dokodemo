package web

import (
	"log"
	"strings"
	"time"
)

func stringToTime(year string, month string, day string, hour string, minute string) time.Time {
	ymds := []string{year, month, day}
	hms := []string{hour, minute}
	ymd := strings.Join(ymds, "-")
	hm := strings.Join(hms, ":")
	ymdhms := []string{ymd, hm}
	ymdhm := strings.Join(ymdhms, " ")
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation("2006-01-02 15:04", ymdhm, jst)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
