package client

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
