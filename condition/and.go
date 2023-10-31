package condition

// AndCondition .
type AndCondition struct {
	Left  Condition
	Right Condition
}

// NewAndCondition .
func NewAndCondition(left, right Condition) Condition {
	return &AndCondition{Left: left, Right: right}
}

func (q *AndCondition) GetField() string {
	return ""
}

func (q *AndCondition) GetValue() interface{} {
	return nil
}

func (q *AndCondition) SetField(field string) Condition {
	q.Left.SetField(field)
	q.Right.SetField(field)

	return q
}
