package condition

// Condition .
type Condition interface {
	SetField(field string) Condition
	GetField() string
	GetValue() interface{}
}
