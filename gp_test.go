package gp

import "testing"

const (
	codeExpected = "package dummy\n"
	testExpected = "package dummy\n"
)

func createInfo() Information {
	name := "dummy"
	info := New(name, name, false)
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
