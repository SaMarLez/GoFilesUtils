package gofilesutils

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func File_Exists(ptf string) bool {
	_, err := os.Stat(ptf)
	return !errors.Is(err, os.ErrNotExist)
}

func Dir_Exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func TxtFile_as_strslice(ptf string) (lines []string) {

	file, err := os.Open(ptf)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error durante el escaneo: %s", err)
	}

	return lines
}

func Get_file_list(pattern string, dirpath string) (lst_files []string) {
	lknf := dirpath + "\\" + pattern
	matches, err := filepath.Glob(lknf)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		lst_files = append(lst_files, match)
	}
	return lst_files
}

func Copy_file(sfp string, dfp string) {

	sf, err := os.Open(sfp)
	if err != nil {
		log.Fatal(err)
	}
	defer sf.Close()

	df, err := os.Create(dfp)
	if err != nil {
		log.Fatal(err)
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		log.Fatal(err)
	}

	err = df.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
