package variables

import "errors"

func Increment(amount int) VariableUpdateFn {
	return func(v Variable) (Variable, error) {
		if v.Type() != VariableType_INT {
			return Int(0), nil
		}
		val, ok := v.Value().(int)
		if !ok {
			return NullVariable, errors.New("failed to parse variable")
		}
		newVal := simpleVariable{variableType: VariableType_INT, value: val + amount}
		return newVal, nil
	}
}
