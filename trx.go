package rpc

import (
	"github.com/asuleymanov/rpc/transactions"
	"github.com/asuleymanov/rpc/types"
	"github.com/pkg/errors"
	_ "time"
)

/*//SendTrx generates and sends an array of transactions to GOLOS.
func (client *Client) SendTrx(username string, strx []types.Operation) (*BResp, error) {
	// Получение необходимых параметров
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
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
	privKeys, err := client.SigningKeys(strx[0])
	if err != nil {
		return nil, err
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, client.Chain); err != nil {
		return nil, err
	}

	// Отправка транзакции
	resp, err := client.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return nil, err
	}
	var bresp BResp

	bresp.ID = resp.ID
	bresp.BlockNum = resp.BlockNum
	bresp.TrxNum = resp.TrxNum
	bresp.Expired = resp.Expired

	return &bresp, nil
}*/

//Старые методы надо разобраться с ними
func (api *Client) Send_Trx(username string, strx types.Operation) (*BResp, error) {
	// Получение необходимых параметров
	props, err := api.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(username, strx)

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return nil, errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return nil, errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	} else {
		var bresp BResp

		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}
}

func (api *Client) Send_Arr_Trx(username string, strx []types.Operation) (*BResp, error) {
	// Получение необходимых параметров
	props, err := api.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
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
	privKeys := api.Signing_Keys(username, strx[0])

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return nil, errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return nil, errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	} else {
		var bresp BResp

		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}
}

func (api *Client) Verify_Trx(username string, strx types.Operation) (bool, error) {
	// Получение необходимых параметров
	props, err := api.Database.GetDynamicGlobalProperties()
	if err != nil {
		return false, errors.Wrapf(err, "Error get DynamicGlobalProperties: ")
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return false, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	privKeys := api.Signing_Keys(username, strx)

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return false, errors.Wrapf(err, "Error Sign: ")
	}

	// Отправка транзакции
	resp, err := api.Database.GetVerifyAuthoruty(tx.Transaction)

	if err != nil {
		return false, errors.Wrapf(err, "Error BroadcastTransactionSynchronous: ")
	} else {
		return resp, nil
	}
}
