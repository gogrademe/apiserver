package envconfig

import "testing"

func TestString(t *testing.T) {
	s := String("path", "d", "some test env var")
	t.Logf("%v", s)

	t.Log(environment)
}

func TestStringOption(t *testing.T) {
	s := StringOption("test1", "d", []string{"cat", "dog", "d"}, "some test env var")
	t.Logf("%v", s)

	t.Log(environment)
}

func TestBool(t *testing.T) {
	s := Bool("something", true, "some test env var")
	t.Logf("%v", s)

	t.Log(environment)
}

func TestPrintDefaults(t *testing.T) {
	PrintDefaults()
}
