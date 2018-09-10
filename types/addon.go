package types

// Add-on struct

//POW is an additional structure used by other structures.
type POW struct {
	Worker    string `json:"worker"`
	Input     string `json:"input"`
	Signature string `json:"signature"`
	Work      string `json:"work"`
}

//POW2Input is an additional structure used by other structures.
type POW2Input struct {
	WorkerAccount string `json:"worker_account"`
	PrevBlock     []byte `json:"prev_block"`
	Nonce         uint64 `json:"nonce"`
}

//Beneficiary is an additional structure used by other structures.
type Beneficiary struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

//CommentPayoutBeneficiaries is an additional structure used by other structures.
type CommentPayoutBeneficiaries struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
}
