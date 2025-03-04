package files

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func DoFiles21() {
	zipFilePath := "./files/3.5.2/task.zip"
	destDir := "./files/tmp/extracted_files"

	err := unzip(zipFilePath, destDir)
	if err != nil {
		fmt.Println("Failed to unzip:", err)
		return
	}
	defer removeDir(destDir)

	err1 := filepath.Walk(destDir, myFuncWalk)
	if err1 != nil {
		fmt.Println("обход произошёл с ошибкой: " + err.Error())
	}
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path")
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(fPath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
				return err
			}
			outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			rc, err := f.Open()
			if err != nil {
				return err
			}
			_, err = io.Copy(outFile, rc)

			outFile.Close()
			rc.Close()

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func removeDir(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println("Error deleting directory:", err)
		return
	}
}

func myFuncWalk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if strings.Contains(info.Name(), ".txt") {
		rc, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer rc.Close()

		reader := csv.NewReader(rc)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("reading file error = %s", err)
			return err
		}

		if len(records) == 10 {
			fmt.Println(records[4][2])
			return nil
		}
	}

	return nil
}
