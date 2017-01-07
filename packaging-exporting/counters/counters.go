package counters

type AlertCounter int
type length float64
type Time int

func ReturnLength(value float64) length {
	return length(value)
}
