package client

import (
	// Stdlib
	"encoding/json"
	"io/ioutil"
	"log"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/transports/websocket"
	"github.com/asuleymanov/golos-go/types"
)

type User struct {
	Name string `json:"username"`
	PKey string `json:"posting_key"`
	AKey string `json:"active_key"`
	OKey string `json:"owner_key"`
	MKey string `json:"memo_key"`
}

type Golos struct {
	Rpc  *rpc.Client
	User *User
}

func readconfig() *User {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Println(errors.Wrapf(e, "Error read config.json: "))
	}

	var jsontype *User
	if erru := json.Unmarshal(file, &jsontype); erru != nil {
		log.Println(errors.Wrapf(erru, "Error UnMarshal config.json: "))
		return nil
	}
	return jsontype
}

func initclient(url string) *rpc.Client {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url)
	if err != nil {
		panic(errors.Wrapf(err, "Error Websocket: "))
	}

	// Инициализация RPC клиента
	client, err := rpc.NewClient(t)
	if err != nil {
		panic(errors.Wrapf(err, "Error RPC: "))
	}
	//defer client.Close()
	return client
}

func NewApi(url string) *Golos {
	return &Golos{
		Rpc:  initclient(url),
		User: readconfig(),
	}
}

func (api *Golos) Send_Trx(strx types.Operation, chain string) error {
	var ChainId *transactions.Chain
	// Получение необходимых параметров
	props, err := api.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(strx)

	// Определяем ChainId
	switch chain {
	case "steem":
		ChainId = transactions.SteemChain
	case "golos":
		ChainId = transactions.GolosChain
	case "test":
		ChainId = transactions.TestChain
	}
	// Подписываем транзакцию
	if err := tx.Sign(privKeys, ChainId); err != nil {
		return errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.Rpc.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	}
	log.Printf("%+v\n", *resp)
	return nil
}

func (api *Golos) Send_Arr_Trx(strx []types.Operation, chain string) error {
	var ChainId *transactions.Chain
	// Получение необходимых параметров
	props, err := api.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	for _, val := range strx {
		tx.PushOperation(val)
	}

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(strx[0])

	// Определяем ChainId
	switch chain {
	case "steem":
		ChainId = transactions.SteemChain
	case "golos":
		ChainId = transactions.GolosChain
	case "test":
		ChainId = transactions.TestChain
	}
	// Подписываем транзакцию
	if err := tx.Sign(privKeys, ChainId); err != nil {
		return errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.Rpc.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	}
	log.Printf("%+v\n", *resp)
	return nil
}
