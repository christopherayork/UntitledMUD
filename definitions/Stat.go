package definitions

type Stat struct {
	current float64
	max float64
	multiplier float64
}

func (s Stat) Add(x float64) (float64, bool) {
	if x < 0 { return s.current, false }
	added := x
	if s.current + x > s.max {
		added = s.max - s.current
	}
	s.current += added
	return s.current, true
	// we save added for informational purposes to the player
	// later we might output this
}

func (s Stat) Sub(x float64) (float64, bool) {
	if x < 0 { return s.current, false}
	subbed := x
	if s.current - x < 0 {
		subbed = s.current
	}
	s.current -= subbed
	return s.current, true
}

func (s Stat) Fill() bool {
	if s.current >= s.max { return false }
	s.current = s.max
	return true
}

func (s Stat) Set(x float64) (float64, bool) {
	if s.current == x { return s.current, false }
	if x < 0 {
		s.current = 0.0
		return s.current, true
	} else if x > s.max {
		s.current = s.max
		return s.current, true
	} else {
		s.current = x
		return s.current, true
	}
}

func (s Stat) Get() float64 {
	return s.current
}
