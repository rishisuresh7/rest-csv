package factory

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"rest-csv/config"
	"rest-csv/utility"
)

var fileWriter sync.Once

type Factory interface {
	ReadWriter(file string) (*os.File, error)
}

type factory struct {
	config *config.Config
	header []string
	files  map[string]*os.File
}

func NewFactory(c *config.Config) Factory {
	f := &factory{
		config: c,
		header: []string{},
	}

	return f
}

func (f *factory) initializeFiles(fileName string) (*os.File, error) {
	fileNameType := fmt.Sprintf("%s.csv", fileName)
	file, err := os.OpenFile(filepath.Join(f.config.DataLocation, fileNameType), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() == 0 {
		csvWriter := csv.NewWriter(file)
		err = csvWriter.Write(f.header)
		csvWriter.Flush()
	}

	return file, nil
}

func (f *factory) initialize() (map[string]*os.File, error) {
	var scErr error
	fileWriter.Do(func() {
		f.files = make(map[string]*os.File)
		for _, name := range f.config.Categories {
			fileName := utility.SanitizeFileName(name)
			file, err := f.initializeFiles(fileName)
			if err != nil {
				scErr = err
				f.files = nil
				return
			}

			f.files[fileName] = file
		}
	})

	return f.files, scErr
}

func (f *factory) ReadWriter(file string) (*os.File, error) {
	files, err := f.initialize()
	if err != nil {
		return nil, fmt.Errorf("ReadWriter: unable to get file(category) pointer: %s", err)
	}

	fileName := utility.SanitizeFileName(file)
	fl, ok := files[fileName]
	if !ok {
		return nil, fmt.Errorf("ReadWriter: invalid category: %s", err)
	}

	return fl, nil
}
