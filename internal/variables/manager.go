package variables

type VariableUpdateFn func(Variable) (Variable, error)

type VariableManager struct {
	variables map[string]Variable
}

func (m *VariableManager) Get(key string) Variable {
	v, ok := m.variables[key]
	if !ok {
		return simpleVariable{variableType: VariableType_NULL, value: nil}
	}
	return v
}

func (m *VariableManager) Set(key string, value Variable) {
	m.variables[key] = value
}

func (m *VariableManager) Update(key string, updateFn VariableUpdateFn) Variable {
	currentVal := m.Get(key)
	newVal, err := updateFn(currentVal)
	if err != nil {
		return currentVal
	}
	m.Set(key, newVal)
	return newVal
}

func NewManager() *VariableManager {
	return &VariableManager{
		variables: make(map[string]Variable),
	}
}
