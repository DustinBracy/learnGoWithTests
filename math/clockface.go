package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func SecondsInRadians(tm time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(tm.Second())))
}

func SecondHandPoint(tm time.Time) Point {
	return angleToPoint(SecondsInRadians(tm))

}

func MinutesInRadians(tm time.Time) float64 {
	return (SecondsInRadians(tm) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(tm.Minute())))
}

func MinuteHandPoint(tm time.Time) Point {
	return angleToPoint(MinutesInRadians(tm))
}

func HoursInRadians(tm time.Time) float64 {
	return (MinutesInRadians(tm) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(tm.Hour()%hoursInClock)))
}

func HourHandPoint(tm time.Time) Point {
	return angleToPoint(HoursInRadians(tm))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
