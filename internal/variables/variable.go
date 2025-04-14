package variables

import "fmt"

type VariableType string

const (
	VariableType_NULL VariableType = "NULL"
	VariableType_INT  VariableType = "INT"
)

var NullVariable = simpleVariable{variableType: VariableType_NULL, value: nil}

type simpleVariable struct {
	variableType VariableType
	value        any
}

func (v simpleVariable) String() string {
	return fmt.Sprintf("[%s] {%v}", v.variableType, v.value)
}

func (v simpleVariable) Type() VariableType {
	return v.variableType
}

func (v simpleVariable) Value() any {
	return v.value
}

func (v simpleVariable) Int() (int, bool) {
	val, ok := v.value.(int)
	return val, ok
}

type Variable interface {
	Type() VariableType
	Value() any
}

func Int(value int) Variable {
	return &simpleVariable{
		variableType: VariableType_INT,
		value:        value,
	}
}
