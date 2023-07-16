package helpers

import "time"

func IsThreeHours(update time.Time) float64 {
	lastUpdate := update
	current := time.Now()
	gap := current.Sub(lastUpdate)
	hour := gap.Hours()

	return hour

}
