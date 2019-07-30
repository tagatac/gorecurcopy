package gorecurcopy

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

const (
	testDir  = "tests"
	testDest = ".47fe2fcb-cd79-509c-ab6f-4a87d5a4369b"
)

func TestRecursiveCopy(t *testing.T) {
	if err := os.RemoveAll(testDest); err != nil {
		t.Errorf("unable to remove test destination directory, %s", err)
	}

	if err := CopyDirectory(
		filepath.Join(testDir, "test1"), filepath.Join(testDest, "test1"),
	); err != nil {
		t.Errorf("copying error, %s", err)
	}

	var files []string

	root := filepath.Join(testDest, "test1")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	if err := os.RemoveAll(testDest); err != nil {
		t.Errorf("unable to remove test destination directory, %s", err)
	}

	var shouldBe []string
	shouldBe = append(shouldBe, filepath.Join(testDest, "test1"))
	shouldBe = append(shouldBe, filepath.Join(testDest, "test1", "sub1"))
	shouldBe = append(shouldBe, filepath.Join(testDest, "test1", "sub1", "recursive_copy.txt"))
	shouldBe = append(shouldBe, filepath.Join(testDest, "test1", "sub1", "sub_sub1"))
	shouldBe = append(shouldBe, filepath.Join(testDest, "test1", "sub1", "sub_sub1", "recursive_copy.txt"))

	equals(t, files, shouldBe)
}

// ben johnson's helper functions: https://github.com/benbjohnson/testing

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
