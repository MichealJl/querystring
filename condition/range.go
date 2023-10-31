package condition

// RangeCondition .
type RangeCondition struct {
	field string
	start interface{}
	end   interface{}
}

// NewRangeCondition .
func NewRangeCondition(start, end interface{}) Condition {
	return &RangeCondition{
		start: start,
		end:   end,
	}
}

func (q *RangeCondition) GetField() string {
	return q.field
}

func (q *RangeCondition) GetValue() interface{} {
	return []interface{}{q.start, q.end}
}

// SetField .
func (q *RangeCondition) SetField(field string) Condition {
	q.field = field

	return q
}
