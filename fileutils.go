package gofilesutils

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

func File_Exists(ptf string) (err error) {
	if _, err := os.Stat(ptf); err != nil {
		return err
	}
	return nil
}

func Dir_Exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// TxtFile_as_strslice Splits the content of a plain text file to a slice of strings
func TxtFile_as_strslice(ptf string) (lines []string, err error) {

	if err := File_Exists(ptf); err != nil {
		return nil, err
	}

	file, err := os.Open(ptf)
	if err != nil {
		log.Printf("Error al abrir el archivo: %s", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error durante el escaneo: %s", err)
		return nil, err
	}

	return lines, nil
}

// File_as_str Returns the content of an plain text file into a single string
func File_as_str(path string) (string, error) {
	var content string = ""
	if err := File_Exists(path); err != nil {
		return content, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error al abrir el archivo: %s", err)
		return content, err
	}

	content = string(file)
	return content, nil
}

// Get_file_list Returns a list of filenames that match with the given pattern in the directory
func Get_file_list(pattern string, dirpath string) (lst_files []string, err error) {
	lknf := dirpath + "\\" + pattern
	matches, err := filepath.Glob(lknf)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	for _, match := range matches {
		lst_files = append(lst_files, match)
	}
	return lst_files, nil
}

// Copy_file Creates a new file, then syncs the content of the source file
func Copy_file(sfp string, dfp string) error {

	sf, err := os.Open(sfp)
	if err != nil {
		log.Println(err)
		return err
	}
	defer sf.Close()

	df, err := os.Create(dfp)
	if err != nil {
		log.Println(err)
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		log.Println(err)
		return err
	}

	err = df.Sync()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Create_dir Creates a directory and any missig parent folders recursively if not exists
func Create_dir(dir_path string) (de bool, err error) {
	if Dir_Exists(dir_path) {
		return true, nil
	}
	err = os.MkdirAll(dir_path, 0755)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

// Write_Txt_File Writes a plain text file given a set of lines, it uses .txt as default extension
// and add a newline at the end of each line of lines, also it creates the path to the file if it don´t exists
func Write_Txt_File(lines []string, path string, file_name string, ext string) error {
	if ext == "" {
		ext = ".txt"
	}

	if _, err := Create_dir(path); err != nil {
		return err
	}

	fn := path + file_name + ext
	nwfile, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer nwfile.Close()

	for _, line := range lines {
		_, err := nwfile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
