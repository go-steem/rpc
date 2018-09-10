package types

import (
	"encoding/json"
)

//OperationObject type from parameter JSON
type OperationObject struct {
	BlockNumber            uint32    `json:"block"`
	TransactionID          string    `json:"trx_id"`
	TransactionInBlock     uint32    `json:"trx_in_block"`
	Operation              Operation `json:"op"`
	OperationType          OpType    `json:"-"`
	OperationInTransaction uint16    `json:"op_in_trx"`
	VirtualOperation       uint64    `json:"virtual_op"`
	Timestamp              *Time     `json:"timestamp"`
}

type rawOperationObject struct {
	BlockNumber            uint32          `json:"block"`
	TransactionID          string          `json:"trx_id"`
	TransactionInBlock     uint32          `json:"trx_in_block"`
	Operation              *operationTuple `json:"op"`
	OperationInTransaction uint16          `json:"op_in_trx"`
	VirtualOperation       uint64          `json:"virtual_op"`
	Timestamp              *Time           `json:"timestamp"`
}

//UnmarshalJSON unpacking the JSON parameter in the OperationObject type.
func (op *OperationObject) UnmarshalJSON(p []byte) error {
	var raw rawOperationObject
	if err := json.Unmarshal(p, &raw); err != nil {
		return err
	}

	op.BlockNumber = raw.BlockNumber
	op.TransactionID = raw.TransactionID
	op.TransactionInBlock = raw.TransactionInBlock
	op.Operation = raw.Operation.Data
	op.OperationType = raw.Operation.Type
	op.OperationInTransaction = raw.OperationInTransaction
	op.VirtualOperation = raw.VirtualOperation
	op.Timestamp = raw.Timestamp
	return nil
}

//MarshalJSON function for packing the OperationObject type in JSON.
func (op *OperationObject) MarshalJSON() ([]byte, error) {
	return JSONMarshal(&rawOperationObject{
		BlockNumber:            op.BlockNumber,
		TransactionID:          op.TransactionID,
		TransactionInBlock:     op.TransactionInBlock,
		Operation:              &operationTuple{op.Operation.Type(), op.Operation},
		OperationInTransaction: op.OperationInTransaction,
		VirtualOperation:       op.VirtualOperation,
		Timestamp:              op.Timestamp,
	})
}
