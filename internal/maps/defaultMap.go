package maps

type Builder[TKey comparable, TValue any] interface {
	Build(key TKey) TValue
}

// Default map implements a map w/ default value assertions
// If value doesn't exist generate one using the NewFn
type DefaultMap[TKey comparable, TValue any] struct {
	ChannelMap map[TKey]TValue
	Builder    Builder[TKey, TValue]
}

func (m *DefaultMap[TKey, TValue]) Get(key TKey) TValue {
	_, ok := m.ChannelMap[key]
	if !ok {
		m.ChannelMap[key] = m.Builder.Build(key)
	}
	return m.ChannelMap[key]
}

func NewDefaultMap[TKey comparable, TValue any](builder Builder[TKey, TValue]) *DefaultMap[TKey, TValue] {
	return &DefaultMap[TKey, TValue]{
		ChannelMap: make(map[TKey]TValue),
		Builder:    builder,
	}
}
