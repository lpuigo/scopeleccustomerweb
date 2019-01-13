package convert

import (
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileConverter struct {
	sourceFileName string
	Converter
}

type MapFunc func(io.Reader) io.Reader

// transform.NewReader(f, charmap.Windows1252.NewDecoder())

func NewFileConverter(infile string, mf MapFunc) FileConverter {
	sourceRs, err := NewRecordSetFromFile(infile, mf)
	if err != nil {
		log.Fatal(err)
	}
	ic := FileConverter{Converter: NewConverter(sourceRs)}
	ic.sourceFileName = strings.TrimSuffix(infile, filepath.Ext(infile))
	return ic
}

// AddTarget adds a converter target process, with given name, given result recordset and given convertion function
//
// Result file is build fron source file name with "_<name>.csv" suffix
//
// It returns a function to be defered in order to close the target result file
//  fclose := fileconverter.AddTarget("example", targetRS, targetRSConvertFunc)
//  defer fclose()
func (c *FileConverter) AddTarget(name string, rs *recordset.RecordSet, convertfunc Convertfunc) func() {
	of, err := os.Create(c.sourceFileName + "_" + name + ".csv")
	if err != nil {
		log.Fatal("could not create target file:", err)
	}
	c.Converter.AddTarget(name, rs, convertfunc, of)
	return func() { of.Close() }
}

func NewRecordSetFromFile(file string, transform MapFunc) (rs *recordset.RecordSet, err error) {
	rs = recordset.NewRecordSet()
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if transform == nil {
		transform = func(r io.Reader) io.Reader {
			return r
		}
	}

	err = rs.AddCSVDataFrom(transform(f))
	if err != nil {
		return nil, err
	}
	return
}
