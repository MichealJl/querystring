package condition

// OrCondition .
type OrCondition struct {
	Left  Condition
	Right Condition
}

// NewOrCondition .
func NewOrCondition(left, right Condition) Condition {
	return &OrCondition{Left: left, Right: right}
}

func (q *OrCondition) GetField() string {
	return ""
}

func (q *OrCondition) GetValue() interface{} {
	return nil
}

func (q *OrCondition) SetField(field string) Condition {
	q.Left.SetField(field)
	q.Right.SetField(field)

	return q
}
