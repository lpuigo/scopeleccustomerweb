package convert

import (
	"fmt"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"strings"
)

func (c Converter) ExportallToCustomerweb(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	cr := c.getSourceCR(i)
	woid := fmt.Sprintf("RECIFS_%05d", i)
	prjcode := "RecetteIFS_R2"
	addrs := rec[c.Index["ADRESSE SITE"]]
	addrss := strings.Split(addrs, ",")
	num_voie := ""
	voie := addrss[0]
	if len(addrss) > 1 {
		num_voie = strings.Trim(addrss[0], " ")
		voie = strings.Trim(addrss[1], " ")
	}
	cp := rec[c.Index["CODE POSTAL SITE"]]
	dep := ""
	if len(cp) > 2 {
		dep = cp[0:2]
	}

	cwr = recordset.Record{
		//"Identity.WorkOrderId",
		woid,
		//"Identity.CustomerReference",
		woid,
		//"Identity.ProjectCode",
		prjcode,
		//"Identity.Contract",
		cr["Instgpc"] + cr["Agentetl"] + cr["Cui"] + cr["Cequip"],
		//"Identity.WorkOrderState",
		rec[c.Index["STATUT"]],
		//"Planning.PlannedDateOrdered",
		rec[c.Index["DATE AU PLUS TOT"]],
		//"Planning.DateRdv",
		rec[c.Index["DATE RDV"]],
		//"Planning.Marge",
		cr["marge"],
		//"Planning.ContractualDate",
		cr["DateContractuelle"],
		//"Templating.Criteria1",
		rec[c.Index["NOM EQUIPEMENT"]][0:3],
		//"Templating.Criteria2",
		rec[c.Index["NOM EQUIPEMENT"]][3:],
		//"Templating.Criteria3",
		cr["Cui"],
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
		cp,
		//"Coordinates.City",
		rec[c.Index["VILLE SITE"]],
		//"Coordinates.Country",
		"FRANCE",
		//"Coordinates.Latitude",
		"",
		//"Coordinates.Longitude",
		"",
		//"Client.ClientFullName",
		rec[c.Index["NOM CLIENT"]],
		//"Client.TelMobile",
		rec[c.Index["TELEPHONE SITE"]],
		//"Client.TelFix",
		rec[c.Index["NOM SITE"]],
		//"Client.MarketType",
		cr["Cegma"],
		//"Client.GTI",
		"",
		//"Client.GTR",
		cr["Engagt"],
		//"Connection.Node",
		cr["Centre"] + cr["Zone"] + cr["Catpc"],
		//"Connection.Information",
		rec[c.Index["DESCRIPTION SITE"]],
		//"ParticularConditions.Criteria1",
		cr["securite"],
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
		rec[c.Index["NOM EQUIPEMENT"]],
		//"Dispatching.Criteria3",
		"",
		//"Dispatching.Criteria4",
		"",
		//"AdditionalInformation",
		rec[c.Index["DESCRIPTION SITE"]],
		//"Comment",
		rec[c.Index["DESCRIPTION"]],
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

func (c Converter) getSourceCol(i, j int) string {
	return c.Source.Get(i)[j]
}

func (c Converter) getSourceCR(i int) map[string]string {
	res := make(map[string]string)
	cr := c.getSourceCol(i, c.Index["COMPTE-RENDU"])
	pairs := strings.Split(cr, "#")
	for _, p := range pairs {
		vals := strings.Split(p, "=")
		val := ""
		if len(vals) > 1 {
			val = vals[1]
		}
		res[vals[0]] = val
	}
	return res
}
