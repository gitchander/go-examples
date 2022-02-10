package devilc

type Interval struct {
	Min, Max float64
	Count    int
}

func (i *Interval) Step() float64 {
	return (i.Max - i.Min) / float64(i.Count-1)
}
