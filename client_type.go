package rpc

const fdt = `"20060102t150405"`

var Key_List = make(map[string]Keys)

type Keys struct {
	PKey string
	AKey string
	OKey string
	MKey string
}

type BResp struct {
	ID       string
	BlockNum uint32
	TrxNum   uint32
	Expired  bool
}

type Beneficiarie struct {
	Account string
	Weight  uint16
}

type PC_Options struct {
	Percent   uint16
	BenefList []Beneficiarie
}

type PC_Vote struct {
	Weight int
}

type ArrTransfer struct {
	To      string
	Memo    string
	Ammount string
}
