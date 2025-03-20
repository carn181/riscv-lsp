package utils

import (
	"path/filepath"
	"os"
	"fmt"
	"os/exec"
	
)

func Uri2Path(uri string) string{
	return uri[7:]    // file:// is 7 chars
}

func Path2Uri(path string) string{
	return "file://"+path    // file:// is 7 chars
}

func FilesInPath(path string)([]string, error){
	var files []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error{
			if err != nil {
				return err
			}
			if !info.IsDir(){
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("Couldn't Find Directory in Path")
	}
	return files, nil
}

func FileExists(path string) bool{
	_, err := os.Stat(path)

	return err == nil
}

func EnsurePathExists(path string){
	dir := filepath.Dir(path)
	if !FileExists(path) {
		os.MkdirAll(dir, 0700) // Create your file
		f, _ := os.Create(path)
		defer f.Close()
	}
}

func RunCommand(dir string, command string)([]byte, error){
	os.Chdir(dir)
	
	cmd := exec.Command("bash","-c",command)
	out, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			code := exitError.ExitCode()
			if code != 1 && code != 0{
				return nil, fmt.Errorf("Build failed with invalid exit code: %d\n",exitError.ExitCode())
			}
			out = exitError.Stderr
		}
	}
	return out, nil
}
