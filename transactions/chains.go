package transactions

type Chain struct {
	ID string
}

var SteemChain = &Chain{
	ID: "0000000000000000000000000000000000000000000000000000000000000000",
}

var GolosChain = &Chain{
	ID: "782a3039b478c839e4cb0c941ff4eaeb7df40bdd68bd441afd444b9da763de12",
}
