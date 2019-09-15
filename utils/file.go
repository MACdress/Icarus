package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// 如果文件存在，返回true
func Exists(filename string) bool {
	info, err := os.Stat(filename)

	return err == nil && !info.IsDir()
}

//用数据覆盖文件
func Overwrite(fileName string, data []byte) bool {
	f, err := os.Create(fileName)
	if err != nil {
		return false
	}

	_, err = f.Write(data)
	return err == nil
}

//返回完整的项目路径
func ExpandedFilename(filename string) string {
	if filename == "" {
		return ""
	}

	if len(filename) > 2 && filename[:2] == "~/" {
		if usr, err := user.Current(); err == nil {
			filename = filepath.Join(usr.HomeDir, filename[2:])
		}
	}

	result, err := filepath.Abs(filename)

	if err != nil {
		panic(err)
	}

	return result
}

//在目录中提取zip文件
func Unzip(src, dest string) (fileNames []string, err error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fileNames, err
	}

	defer r.Close()

	for _, f := range r.File {
		// Skip directories like __OSX
		if strings.HasPrefix(f.Name, "__") {
			continue
		}

		fn, err := copyToFile(f, dest)
		if err != nil {
			return fileNames, err
		}

		fileNames = append(fileNames, fn)
	}

	return fileNames, nil
}

//拷贝ZIP文件复制到指定路径
func copyToFile(f *zip.File, dest string) (fileName string, err error) {
	rc, err := f.Open()
	if err != nil {
		return fileName, err
	}

	defer rc.Close()

	// Store filename/path for returning and using later on
	fileName = filepath.Join(dest, f.Name)

	if f.FileInfo().IsDir() {
		// Make Folder
		return fileName, os.MkdirAll(fileName, os.ModePerm)
	}

	// Make File
	var fdir string
	if lastIndex := strings.LastIndex(fileName, string(os.PathSeparator)); lastIndex > -1 {
		fdir = fileName[:lastIndex]
	}

	err = os.MkdirAll(fdir, os.ModePerm)
	if err != nil {
		return fileName, err
	}

	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return fileName, err
	}

	defer fd.Close()
	_, err = io.Copy(fd, rc)
	if err != nil {
		return fileName, err
	}

	return fileName, nil
}

// 从URL出下载文件
func Download(filepath string, url string) error {
	os.MkdirAll("/tmp/Icarus", os.ModePerm)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func DirectoryIsEmpty(path string) bool {
	f, err := os.Open(path)

	if err != nil {
		return false
	}

	defer f.Close()

	_, err = f.Readdirnames(1)

	if err == io.EOF {
		return true
	}

	return false
}
