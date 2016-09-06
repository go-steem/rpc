package transactions

type Chain struct {
	ID string
}

var SteemChain = &Chain{
	ID: "0000000000000000000000000000000000000000000000000000000000000000",
}

var TestChain = &Chain{
	ID: "18dcf0a285365fc58b71f18b3d3fec954aa0c141c44e4e5cb4cf777b9eab274e",
}
