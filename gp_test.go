package gp

import "testing"

const (
	codeExpected = "package dummy\n"
	testExpected = "package dummy\n"
)

func createInfo() Information {
	name := "dummy"
	info := New(name, name, false, false, false)
	return info
}

func TestTemplatePackageCode(t *testing.T) {
	info := createInfo()
	code, err := info.CreatePackageCode()
	if err != nil {
		t.Fatalf("CreatePackageCode() failed: %s", err)
	}

	if code != codeExpected {
		t.Fatalf("generated code is %q, expected %q", code, codeExpected)
	}
}

func TestTemplatePackageTest(t *testing.T) {
	info := createInfo()
	test, err := info.CreatePackageTest()
	if err != nil {
		t.Fatalf("CreatePackageTest() failed: %s", err)
	}

	if test != testExpected {
		t.Fatalf("generated test is %q, expected %q", test, testExpected)
	}
}

func TestFirstUpper(t *testing.T) {
	for _, tv := range []struct {
		i string
		e string
	}{
		{i: "", e: ""},
		{i: "h", e: "H"},
		{i: "H", e: "H"},
		{i: "hello", e: "Hello"},
		{i: "Hello", e: "Hello"},
		{i: "hELLO", e: "HELLO"},
		{i: "HELLO", e: "HELLO"},
	} {
		if r := firstUpper(tv.i); r != tv.e {
			t.Errorf("firstUpper(%q) is %q, expected %q", tv.i, r, tv.e)
		}
	}
}
