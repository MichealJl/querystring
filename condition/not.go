package condition

// NotCondition .
type NotCondition struct {
	Left  Condition
	Right Condition
}

// NewNotCondition .
func NewNotCondition(left, right Condition) Condition {
	return &NotCondition{
		left, right,
	}
}

func (q *NotCondition) GetField() string {
	return ""
}

func (q *NotCondition) GetValue() interface{} {
	return nil
}

func (q *NotCondition) SetField(field string) Condition {
	q.Left.SetField(field)
	q.Right.SetField(field)

	return q
}
