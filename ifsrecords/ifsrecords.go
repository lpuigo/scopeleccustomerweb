package ifsrecords

import (
	"github.com/lpuig/scopeleccustomerweb/convert"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type IFSConverter struct {
	sourceFileName string
	convert.Converter
}

func NewIFSConverter(infile string) IFSConverter {
	sourceRs, err := NewRecordSetFromFile(infile)
	if err != nil {
		log.Fatal(err)
	}
	ic := IFSConverter{Converter: convert.NewConverter(sourceRs)}
	ic.sourceFileName = strings.TrimSuffix(infile, filepath.Ext(infile))
	return ic
}

func (c *IFSConverter) AddTarget(name string, rs *recordset.RecordSet, convertfunc convert.Convertfunc) func() {
	of, err := os.Create(c.sourceFileName + "_" + name + ".csv")
	if err != nil {
		log.Fatal("could not create target file:", err)
	}
	c.Converter.AddTarget(name, rs, convertfunc, of)
	return func() { of.Close() }
}

func NewRecordSetFromFile(file string) (rs *recordset.RecordSet, err error) {
	rs = recordset.NewRecordSet()
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = rs.AddCSVDataFrom(transform.NewReader(f, charmap.Windows1252.NewDecoder()))
	if err != nil {
		return nil, err
	}
	return
}
