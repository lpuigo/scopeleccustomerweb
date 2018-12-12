package convert

import (
	"fmt"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"io"
)

type Convertfunc func(i int) (cwr recordset.Record, err error)

type target struct {
	name          string
	target        *recordset.RecordSet
	convertRecord Convertfunc
	writer        io.Writer
}

type Converter struct {
	source  *recordset.RecordSet
	index   map[string]int
	targets []target
}

func NewConverterToCustomerWeb(sourceRs *recordset.RecordSet) Converter {
	ind := make(map[string]int)
	for _, col := range sourceRs.GetHeader().GetKeys() {
		colinds, err := sourceRs.GetRecordColNumByName(col)
		if err != nil || len(colinds) < 1 {
			panic("could not create converter index")
		}
		ind[col] = colinds[0]
	}

	ctr := Converter{
		source: sourceRs,
		//target: custwebrecords.NewCustomerWebRecords(),
		index: ind,
	}
	//ctr.convertRecord = ctr.convertExportALLRecord
	return ctr
}

func (c *Converter) AddTarget(name string, targetrecordset *recordset.RecordSet, convertfunc Convertfunc, w io.Writer) {
	c.targets = append(c.targets, target{
		name:          name,
		target:        targetrecordset,
		convertRecord: convertfunc,
		writer:        w,
	})
}

func (c Converter) Convert() error {
	for i := 0; i < c.source.Len(); i++ {
		for _, tgt := range c.targets {
			cwr, err := tgt.convertRecord(i)
			if err != nil {
				return err
			}
			tgt.target.Add(cwr)
		}
	}
	return nil
}

func (c Converter) WriteCSV() error {
	for _, tgt := range c.targets {
		err := tgt.target.WriteCSVTo(tgt.writer)
		if err != nil {
			return fmt.Errorf("%s:%v", tgt.name, err)
		}
	}
	return nil
}
