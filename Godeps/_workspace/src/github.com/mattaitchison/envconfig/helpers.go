// generated code -- DO NOT EDIT

package envconfig

import "net"

// String retrieves a environment variable by name and parses it to a string
// defaultVal will be returned if the variable is not found.
func String(name string, defaultVal string, description string) string {
	v := NewVar(newStringValue(defaultVal), name, description)
	return v.Value.Get().(string)
}

// StringOption like String except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func StringOption(name string, defaultVal string, options []string, description string) string {
	v := NewVar(newStringValue(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(string) {
			return v.Value.Get().(string)
		}
	}
	return defaultVal
}

// Bool retrieves a environment variable by name and parses it to a bool
// defaultVal will be returned if the variable is not found.
func Bool(name string, defaultVal bool, description string) bool {
	v := NewVar(newBoolValue(defaultVal), name, description)
	return v.Value.Get().(bool)
}

// Float64 retrieves a environment variable by name and parses it to a float64
// defaultVal will be returned if the variable is not found.
func Float64(name string, defaultVal float64, description string) float64 {
	v := NewVar(newFloat64Value(defaultVal), name, description)
	return v.Value.Get().(float64)
}

// Float64Option like Float64 except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func Float64Option(name string, defaultVal float64, options []float64, description string) float64 {
	v := NewVar(newFloat64Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(float64) {
			return v.Value.Get().(float64)
		}
	}
	return defaultVal
}

// Int retrieves a environment variable by name and parses it to a int
// defaultVal will be returned if the variable is not found.
func Int(name string, defaultVal int, description string) int {
	v := NewVar(newIntValue(defaultVal), name, description)
	return v.Value.Get().(int)
}

// IntOption like Int except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func IntOption(name string, defaultVal int, options []int, description string) int {
	v := NewVar(newIntValue(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(int) {
			return v.Value.Get().(int)
		}
	}
	return defaultVal
}

// Int64 retrieves a environment variable by name and parses it to a int64
// defaultVal will be returned if the variable is not found.
func Int64(name string, defaultVal int64, description string) int64 {
	v := NewVar(newInt64Value(defaultVal), name, description)
	return v.Value.Get().(int64)
}

// Int64Option like Int64 except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func Int64Option(name string, defaultVal int64, options []int64, description string) int64 {
	v := NewVar(newInt64Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(int64) {
			return v.Value.Get().(int64)
		}
	}
	return defaultVal
}

// Uint retrieves a environment variable by name and parses it to a uint
// defaultVal will be returned if the variable is not found.
func Uint(name string, defaultVal uint, description string) uint {
	v := NewVar(newUintValue(defaultVal), name, description)
	return v.Value.Get().(uint)
}

// UintOption like Uint except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func UintOption(name string, defaultVal uint, options []uint, description string) uint {
	v := NewVar(newUintValue(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(uint) {
			return v.Value.Get().(uint)
		}
	}
	return defaultVal
}

// Uint64 retrieves a environment variable by name and parses it to a uint64
// defaultVal will be returned if the variable is not found.
func Uint64(name string, defaultVal uint64, description string) uint64 {
	v := NewVar(newUint64Value(defaultVal), name, description)
	return v.Value.Get().(uint64)
}

// Uint64Option like Uint64 except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func Uint64Option(name string, defaultVal uint64, options []uint64, description string) uint64 {
	v := NewVar(newUint64Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(uint64) {
			return v.Value.Get().(uint64)
		}
	}
	return defaultVal
}

// IP retrieves a environment variable by name and parses it to a net.IP
// defaultVal will be returned if the variable is not found.
func IP(name string, defaultVal net.IP, description string) net.IP {
	v := NewVar(newIPValue(defaultVal), name, description)
	return v.Value.Get().(net.IP)
}
