package transactions

type Chain struct {
	ID string
}

var SteemChain = &Chain{
	ID: "0000000000000000000000000000000000000000000000000000000000000000",
}

var TestChain = &Chain{
	ID: "5876894a41e6361bde2e73278f07340f2eb8b41c2facd29099de9deef6cdb679",
}
