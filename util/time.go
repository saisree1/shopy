package util

import "time"

var istLocation *time.Location

func init() {
	istLocation, _ = time.LoadLocation("Asia/Kolkata")
}

func GetIST() time.Time {
	t := time.Now().In(istLocation)
	updatedDate, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return updatedDate
}
