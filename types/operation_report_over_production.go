package types

//ReportOverProductionOperation represents report_over_production operation data.
type ReportOverProductionOperation struct {
	Reporter string `json:"reporter"`
}

//Type function that defines the type of operation ReportOverProductionOperation.
func (op *ReportOverProductionOperation) Type() OpType {
	return TypeReportOverProduction
}

//Data returns the operation data ReportOverProductionOperation.
func (op *ReportOverProductionOperation) Data() interface{} {
	return op
}
