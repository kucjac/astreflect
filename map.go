package astreflect

import "fmt"

var _ Type = (*MapType)(nil)

// MapType is the type wrapper for the standar key value map type.
type MapType struct {
	Key   Type
	Value Type
}

// Name implements Type interface.
func (m *MapType) Name(identified bool, packageContext string) string {
	return fmt.Sprintf("map[%s]%s", m.Key.Name(identified, packageContext), m.Value.Name(identified, packageContext))
}

// FullName implements Type interface.
func (m *MapType) FullName() string {
	return fmt.Sprintf("map[%s]%s", m.Key.FullName(), m.Value.FullName())
}

// PkgPath implements Type interface.
func (m *MapType) PkgPath() PkgPath {
	return builtInPkgPath
}

// Kind implements Type interface.
func (m *MapType) Kind() Kind {
	return Map
}

// Elem as the map has both the key and value it needs to be dereferenced manually.
func (m *MapType) Elem() Type {
	return nil
}

// String implements Type interface.
func (m *MapType) String() string {
	return m.Name(true, "")
}
