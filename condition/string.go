package condition

// StringCondition 普通字符串.
type StringCondition struct {
	field string
	value string
}

// NewStringCondition .
func NewStringCondition(value string) Condition {
	return &StringCondition{
		value: value,
	}
}

func (q *StringCondition) GetField() string {
	return q.field
}

func (q *StringCondition) GetValue() interface{} {
	return q.value
}

func (q *StringCondition) SetField(field string) Condition {
	q.field = field

	return q
}
