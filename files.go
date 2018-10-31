package main

import (
	"strings"
	"path/filepath"
	"os"
	"fmt"
	"io"
)
//get all subdirectory of root  exclude special directory by fuzzy string match
func walkPath(root string, exclude []string) (map[string]string, error) {
	pathSet := map[string]string{}

	if strings.TrimSpace(root) == "" {
		return pathSet, nil
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		pathSet[path] = path
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
		return pathSet, nil
	}
	//delete excluded path from map
	for _, path := range exclude {
		for k := range pathSet {
			if strings.Contains(k, path) {
				delete(pathSet, k)
			}
		}
	}
	//delete root path
	for k := range pathSet {
		if root != "." && root != "./" {
			pathSet[k] = pathSet[k][len(root):]
			if pathSet[k] == "" {
				delete(pathSet, k)
			}
		}
	}
	return pathSet, nil
}

// Copy file src to dest,
func Copy(src, dest string, exclude []string) error {
	pathSet,error := walkPath(src,exclude)
	if error != nil{
		return error
	}
	//walk through pathSet to create all directory
	for path,name := range pathSet{
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			destAll := filepath.Join(dest,name)
			exist,err := exists(destAll)
			if err != nil{
				return err
			}
			//directory is exists
			if exist{
				continue
			}
			if err := os.MkdirAll(destAll, info.Mode()); err != nil {
				return err
			}
		}
	}
	//copy file and symlink
	for path,name := range pathSet{
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		//copy symlink
		if info.Mode()&os.ModeSymlink != 0 {
			src, err := os.Readlink(src)
			if err != nil {
				return err
			}
			destName := filepath.Join(dest,name)
			error := os.Symlink(src, destName)
			if error != nil{
				return error
			}
		}
		//copy file
		if info.Mode()&os.ModeType == 0 {
			destName := filepath.Join(dest,name)
			f, err := os.Create(destName)
			if err != nil {
				return err
			}
			defer f.Close()

			if err = os.Chmod(f.Name(), info.Mode()); err != nil {
				return err
			}
			s, err := os.Open(path)
			if err != nil {
				return err
			}
			defer s.Close()

			_, err = io.Copy(f, s)
			if err != nil{
				return err
			}
		}
	}
	return nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
