package main

import (
	"fmt"
	"github.com/lpuig/scopeleccustomerweb/custwebrecords"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"os"
)

const (
	testinfile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\exportAll_test.csv`
)

func main() {

}

func NewRecordSetFromFile(file string) (rs *recordset.RecordSet, err error) {
	rs = recordset.NewRecordSet()
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rs.AddCSVDataFrom(f)
	if err != nil {
		return nil, err
	}
	return
}

type Converter struct {
	source *recordset.RecordSet
	target *recordset.RecordSet
	index  map[string]int
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

	return Converter{
		source: sourceRs,
		target: custwebrecords.NewCustomerWebRecords(),
		index:  ind,
	}
}

func (c Converter) Convert() error {
	for i := 0; i < c.source.Len(); i++ {
		cwr, err := c.convertRecord(i)
		if err != nil {
			return err
		}
		c.target.Add(cwr)
	}
	return nil
}

func (c Converter) convertRecord(i int) (cwr recordset.Record, err error) {
	woid := fmt.Sprintf("RECIFS_%05d", i)
	prjcode := "RECETTE IFS"
	cwr = recordset.Record{
		//"Identity.WorkOrderId",
		woid,
		//"Identity.CustomerReference",
		woid,
		//"Identity.ProjectCode",
		prjcode,
		//"Identity.Contract",
		//"Identity.WorkOrderState",
		//"Planning.PlannedDateOrdered",
		//"Planning.DateRdv",
		//"Planning.Marge",
		//"Planning.ContractualDate",
		//"Templating.Criteria1",
		//"Templating.Criteria2",
		//"Templating.Criteria3",
		//"Templating.Criteria4",
		//"Coordinates.StreetNumber",
		//"Coordinates.NumberComplement",
		//"Coordinates.StreetName",
		//"Coordinates.AddressComplement",
		//"Coordinates.PostalCode",
		//"Coordinates.City",
		//"Coordinates.Country",
		//"Coordinates.Latitude",
		//"Coordinates.Longitude",
		//"Client.ClientFullName",
		//"Client.TelMobile",
		//"Client.TelFix",
		//"Client.MarketType",
		//"Client.GTI",
		//"Client.GTR",
		//"Connection.Node",
		//"Connection.Information",
		//"ParticularConditions.Criteria1",
		//"ParticularConditions.Criteria2",
		//"ParticularConditions.Criteria3",
		//"ParticularConditions.Criteria4",
		//"ParticularConditions.Criteria5",
		//"ParticularConditions.Criteria6",
		//"ParticularConditions.Criteria7",
		//"ParticularConditions.Criteria8",
		//"ParticularConditions.Criteria9",
		//"ParticularConditions.Criteria10",
		//"Routing.Criteria1",
		//"Routing.Criteria2",
		//"Routing.Criteria3",
		//"Routing.Criteria4",
		//"Dispatching.Criteria1",
		//"Dispatching.Criteria2",
		//"Dispatching.Criteria3",
		//"Dispatching.Criteria4",
		//"AdditionalInformation",
		//"Comment",
		//"Attachment.File1",
		//"Attachment.Name1",
		//"Attachment.File2",
		//"Attachment.Name2",
		//"Attachment.File3",
		//"Attachment.Name3",
		//"Extension.Key1",
		//"Extension.Value1",
		//"Extension.Key2",
		//"Extension.Value2",
		//"Extension.Key3",
		//"Extension.Value3",
		//"réponse SPRINT",
	}
	return
}

func (c Converter) getSourceCol(i, j int) string {
	return c.source.Get(i)[j]
}

func (c Converter) getSourceCR(i int) map[string]string {
	res := make(map[string]string)

}
