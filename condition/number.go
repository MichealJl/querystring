package condition

// NumberCondition 可以包含整数（例如123）、
// 浮点数（例如3.14）、科学计数法表示的数字（例如1.23e+4或1.23E-4）等各种数字类型
type NumberCondition struct {
	field string
	value string
}

// NewNumberCondition .
func NewNumberCondition(value string) Condition {
	return &NumberCondition{
		value: value,
	}
}

func (q *NumberCondition) GetField() string {
	return q.field
}

func (q *NumberCondition) GetValue() interface{} {
	return q.value
}

func (q *NumberCondition) SetField(field string) Condition {
	q.field = field

	return q
}
