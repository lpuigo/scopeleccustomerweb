package custwebrecords

import (
	"github.com/lpuig/scopeleccustomerweb/convert"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

type CWConverter struct {
	convert.FileConverter
}

func NewCWConverter(infile string) CWConverter {
	cf := func(r io.Reader) io.Reader {
		return transform.NewReader(r, charmap.Windows1252.NewDecoder())
	}
	return CWConverter{FileConverter: convert.NewFileConverter(infile, cf)}
}

func NewCustomerWebRecords() *recordset.RecordSet {
	rs := recordset.NewRecordSet()
	rs.AddHeader(recordset.Record{
		"Identity.WorkOrderId",
		"Identity.CustomerReference",
		"Identity.ProjectCode",
		"Identity.Contract",
		"Identity.WorkOrderState",
		"Planning.PlannedDateOrdered",
		"Planning.DateRdv",
		"Planning.Marge",
		"Planning.ContractualDate",
		"Templating.Criteria1",
		"Templating.Criteria2",
		"Templating.Criteria3",
		"Templating.Criteria4",
		"Coordinates.StreetNumber",
		"Coordinates.NumberComplement",
		"Coordinates.StreetName",
		"Coordinates.AddressComplement",
		"Coordinates.PostalCode",
		"Coordinates.City",
		"Coordinates.Country",
		"Coordinates.Latitude",
		"Coordinates.Longitude",
		"Client.ClientFullName",
		"Client.TelMobile",
		"Client.TelFix",
		"Client.MarketType",
		"Client.GTI",
		"Client.GTR",
		"Connection.Node",
		"Connection.Information",
		"ParticularConditions.Criteria1",
		"ParticularConditions.Criteria2",
		"ParticularConditions.Criteria3",
		"ParticularConditions.Criteria4",
		"ParticularConditions.Criteria5",
		"ParticularConditions.Criteria6",
		"ParticularConditions.Criteria7",
		"ParticularConditions.Criteria8",
		"ParticularConditions.Criteria9",
		"ParticularConditions.Criteria10",
		"Routing.Criteria1",
		"Routing.Criteria2",
		"Routing.Criteria3",
		"Routing.Criteria4",
		"Dispatching.Criteria1",
		"Dispatching.Criteria2",
		"Dispatching.Criteria3",
		"Dispatching.Criteria4",
		"AdditionalInformation",
		"Comment",
		"Attachment.File1",
		"Attachment.Name1",
		"Attachment.File2",
		"Attachment.Name2",
		"Attachment.File3",
		"Attachment.Name3",
		"Extension.Key1",
		"Extension.Value1",
		"Extension.Key2",
		"Extension.Value2",
		"Extension.Key3",
		"Extension.Value3",
		"réponse SPRINT",
	})
	return rs
}

func (c CWConverter) getSourceCol(i, j int) string {
	return c.Source.Get(i)[j]
}

func (c CWConverter) SqltoCustomerweb(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	prjcode := "RecetteIFS_R2"
	addrs := rec[c.Index["Adresse site"]]
	addrss := strings.Split(addrs, ",")
	num_voie := ""
	voie := addrss[0]
	if len(addrss) > 1 {
		num_voie = strings.Trim(addrss[0], " ")
		voie = strings.Trim(addrss[1], " ")
	}
	cp := rec[c.Index["Code postal site"]]
	dep := ""
	if len(cp) > 2 {
		dep = cp[0:2]
	}

	cwr = recordset.Record{
		//"Identity.WorkOrderId",
		rec[c.Index["NUMINT"]],
		//"Identity.CustomerReference",
		rec[c.Index["NUMINT"]],
		//"Identity.ProjectCode",
		prjcode,
		//"Identity.Contract",
		rec[c.Index["Code Centre"]] + rec[c.Index["Code ETL"]] + rec[c.Index["UI"]],
		//"Identity.WorkOrderState",
		rec[c.Index["Statut"]],
		//"Planning.PlannedDateOrdered",
		rec[c.Index["Planification FT"]],
		//"Planning.DateRdv",
		rec[c.Index["Date de RDV"]],
		//"Planning.Marge",
		"1",
		//"Planning.ContractualDate",
		rec[c.Index["Date contractuelle"]],
		//"Templating.Criteria1",
		rec[c.Index["Activite"]],
		//"Templating.Criteria2",
		rec[c.Index["Produit"]],
		//"Templating.Criteria3",
		rec[c.Index["UI"]],
		//"Templating.Criteria4",
		"",
		//"Coordinates.StreetNumber",
		num_voie,
		//"Coordinates.NumberComplement",
		"",
		//"Coordinates.StreetName",
		voie,
		//"Coordinates.AddressComplement",
		"",
		//"Coordinates.PostalCode",
		rec[c.Index["Code postal site"]],
		//"Coordinates.City",
		rec[c.Index["Ville site"]],
		//"Coordinates.Country",
		"FRANCE",
		//"Coordinates.Latitude",
		"",
		//"Coordinates.Longitude",
		"",
		//"Client.ClientFullName",
		rec[c.Index["Nom client"]],
		//"Client.TelMobile",
		rec[c.Index["Telephone site"]],
		//"Client.TelFix",
		rec[c.Index["ND principal"]],
		//"Client.MarketType",
		"",
		//"Client.GTI",
		"",
		//"Client.GTR",
		"",
		//"Connection.Node",
		rec[c.Index["Code Centre"]] + rec[c.Index["codeEquipement"]],
		//"Connection.Information",
		"",
		//"ParticularConditions.Criteria1",
		"",
		//"ParticularConditions.Criteria2",
		"",
		//"ParticularConditions.Criteria3",
		"",
		//"ParticularConditions.Criteria4",
		"",
		//"ParticularConditions.Criteria5",
		"",
		//"ParticularConditions.Criteria6",
		"",
		//"ParticularConditions.Criteria7",
		"",
		//"ParticularConditions.Criteria8",
		"",
		//"ParticularConditions.Criteria9",
		"",
		//"ParticularConditions.Criteria10",
		"",
		//"Routing.Criteria1",
		"",
		//"Routing.Criteria2",
		"",
		//"Routing.Criteria3",
		"",
		//"Routing.Criteria4",
		"",
		//"Dispatching.Criteria1",
		dep,
		//"Dispatching.Criteria2",
		rec[c.Index["Activite"]] + rec[c.Index["Produit"]],
		//"Dispatching.Criteria3",
		"",
		//"Dispatching.Criteria4",
		"",
		//"AdditionalInformation",
		"",
		//"Comment",
		"",
		//"Attachment.File1",
		"",
		//"Attachment.Name1",
		"",
		//"Attachment.File2",
		"",
		//"Attachment.Name2",
		"",
		//"Attachment.File3",
		"",
		//"Attachment.Name3",
		"",
		//"Extension.Key1",
		"",
		//"Extension.Value1",
		"",
		//"Extension.Key2",
		"",
		//"Extension.Value2",
		"",
		//"Extension.Key3",
		"",
		//"Extension.Value3",
		"",
		//"réponse SPRINT",
		"",
	}
	return
}
