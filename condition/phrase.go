package condition

// PhraseCondition 被双引号包含的字符串.
type PhraseCondition struct {
	field string
	value string
}

// NewPhraseCondition .
func NewPhraseCondition(value string) Condition {
	return &PhraseCondition{
		value: value,
	}
}

func (q *PhraseCondition) GetField() string {
	return q.field
}

func (q *PhraseCondition) GetValue() interface{} {
	return q.value
}

func (q *PhraseCondition) SetField(field string) Condition {
	q.field = field

	return q
}
