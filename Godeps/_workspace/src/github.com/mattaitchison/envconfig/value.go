package envconfig

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Value interface {
	String() string
	Set(string) error
	Get() interface{}
}

// -- string Value
type stringValue string

func newStringValue(val string) *stringValue {
	return (*stringValue)(&val)
}
func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) Get() interface{} { return string(*s) }

func (s *stringValue) String() string { return fmt.Sprintf("%s", *s) }

// -- bool Value
type boolValue bool

func newBoolValue(val bool) *boolValue {
	return (*boolValue)(&val)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) Get() interface{} { return bool(*b) }

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

func (b *boolValue) IsBoolFlag() bool { return true }

// -- int Value
type intValue int

func newIntValue(val int) *intValue {
	return (*intValue)(&val)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

func (i *intValue) Get() interface{} { return int(*i) }

func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

// -- int64 Value
type int64Value int64

func newInt64Value(val int64) *int64Value {
	return (*int64Value)(&val)
}

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

func (i *int64Value) Get() interface{} { return int64(*i) }

func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- uint Value
type uintValue uint

func newUintValue(val uint) *uintValue {
	return (*uintValue)(&val)
}

func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

func (i *uintValue) Get() interface{} { return uint(*i) }

func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64) *uint64Value {
	return (*uint64Value)(&val)
}

func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

func (i *uint64Value) Get() interface{} { return uint64(*i) }

func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- float64 Value
type float64Value float64

func newFloat64Value(val float64) *float64Value {
	return (*float64Value)(&val)
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

func (f *float64Value) Get() interface{} { return float64(*f) }

func (f *float64Value) String() string { return fmt.Sprintf("%v", *f) }

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration) *durationValue {
	return (*durationValue)(&val)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

func (d *durationValue) Get() interface{} { return time.Duration(*d) }

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

// -- float64 Value
type ipValue net.IP

func newIPValue(val net.IP) *ipValue {
	return (*ipValue)(&val)
}

func (f *ipValue) Set(s string) error {
	v := net.ParseIP(s)
	*f = ipValue(v)
	if v != nil {
		return errors.New("invalid IP")
	}
	return nil
}

func (f *ipValue) Get() interface{} { return net.IP(*f) }

func (f *ipValue) String() string { return fmt.Sprintf("%v", *f) }
