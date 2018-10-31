package main

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
)

func TestWalkPath(t *testing.T) {
	path, err := walkPath(".", []string{".git", "main.go"})
	if err != nil {
		t.Error("FAILED")
	}
	_, ok := path["files_test.go"]
	if ok {
		t.Log("PASS")
	} else {
		t.Error("FAILED")
	}

	_, ok = path["main.go"]
	if !ok {
		t.Log("PASS")
	} else {
		t.Error("FAILED")
	}
}

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}
func TestCopy(t *testing.T) {
	tempPath := "../tempTest"
	_, err := prepareTestDirTree(tempPath)
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tempPath)
	error := Copy(".", tempPath, []string{".idea", ".git", "files_test.go"})
	if error != nil {
		t.Error("FAILED")
	}
	ok, err := exists("../tempTest/main.go")
	if err != nil {
		t.Error("FAILED")
	}
	if ok {
		t.Log("PASS")
	} else {
		t.Error("FAILED")
	}
	ok, err = exists("../tempTest/files_test.go")
	if !ok {
		t.Log("PASS")
	} else {
		t.Error("FAILED")
	}

}

func TestCopy2(t *testing.T) {
	//src := "/testWalkDir"
	//dest := "/testWalkDirDest"
	//error := Copy(src, dest, []string{})
	//if error != nil{
	//	t.Error("FAILED")
	//}
	//pathSrc,err := walkPath(src,[]string{".DS_Store"})
	//if err != nil{
	//	t.Error("FAILED")
	//}
	//pathDes,e := walkPath(dest,[]string{".DS_Store"})
	//if e != nil{
	//	t.Error("FAILED")
	//}
	//for _,v := range pathSrc {
	//	if _, ok := pathDes[path.Join(dest,v)]; ok {
	//		continue
	//	} else {
	//		t.Error("FAILED")
	//		return
	//	}
	//}
	//t.Log("PASS")
}
