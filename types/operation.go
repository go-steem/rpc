package types

import (
	// Stdlib
	"encoding/json"
	"reflect"

	// Vendor
	"github.com/pkg/errors"
)

// dataObjects keeps mapping operation type -> operation data object.
// This is used later on to unmarshal operation data based on the operation type.
var dataObjects = map[OpType]Operation{
	TypeVote:              &VoteOperation{},
	TypeComment:           &CommentOperation{},
	TypeTransfer:          &TransferOperation{},
	TypeTransferToVesting: &TransferToVestingOperation{},
	TypeWithdrawVesting:   &WithdrawVestingOperation{},
	TypeLimitOrderCreate:  &LimitOrderCreateOperation{},
	TypeLimitOrderCancel:  &LimitOrderCancelOperation{},
	TypeFeedPublish:       &FeedPublishOperation{},
	TypeConvert:           &ConvertOperation{},
	TypeAccountCreate:     &AccountCreateOperation{},
	TypeAccountUpdate:     &AccountUpdateOperation{},
	// TypeWitnessUpdate:  &WitnessUpdateOperation{},
	TypeAccountWitnessVote:  &AccountWitnessVoteOperation{},
	TypeAccountWitnessProxy: &AccountWitnessProxyOperation{},
	TypePOW:                 &POWOperation{},
	// TypeCustom:           &CustomOperation{},
	TypeReportOverProduction: &ReportOverProductionOperation{},
	TypeDeleteComment:        &DeleteCommentOperation{},
	TypeCustomJSON:           &CustomJSONOperation{},
	TypeCommentOptions:       &CommentOptionsOperation{},
	// TypeSetWithdrawVestingRoute: &SetWithdrawVestingRouteOperation{},
	// TypeLimitOrderCreate2:       &LimitOrderCreate2Operation{},
	// TypeChallengeAuthority:      &ChallengeAuthorityOperation{},
	// TypeProveAuthority:          &ProveAuthorityOperation{},
	// TypeRequestAccountRecovery:  &RequestAccountRecoveryOperation{},
	// TypeRecoverAccount:          &RecoverAccountOperation{},
	// TypeChangeRecoveryAccount:   &ChangeRecoverAccountOperation{},
	// TypeEscrowTransfer:          &EscrowTransferOperation{},
	// TypeEscrowDispute:           &EscrowDisputeOperation{},
	// TypeEscrowRelease:           &EescrowReleaseOperation{},
	// TypePOW2:                    &POW2Operation{},
}

// Operation represents an operation stored in a transaction.
type Operation interface {
	// Type returns the operation type as present in the operation object, element [0].
	Type() OpType

	// Data returns the operation data as present in the operation object, element [1].
	//
	// When the operation type is known to this package, this field contains
	// the operation data object associated with the given operation type,
	// e.g. Type is TypeVote -> Data contains *VoteOperation.
	// Otherwise this field contains raw JSON (type *json.RawMessage).
	Data() interface{}
}

type Operations []Operation

func (ops *Operations) UnmarshalJSON(data []byte) error {
	var tuples []*operationTuple
	if err := json.Unmarshal(data, &tuples); err != nil {
		return err
	}

	items := make([]Operation, 0, len(tuples))
	for _, tuple := range tuples {
		items = append(items, tuple.Data)
	}

	*ops = items
	return nil
}

func (ops Operations) MarshalJSON() ([]byte, error) {
	tuples := make([]*operationTuple, 0, len(ops))
	for _, op := range ops {
		tuples = append(tuples, &operationTuple{
			Type: op.Type(),
			Data: op.Data().(Operation),
		})
	}
	return json.Marshal(tuples)
}

type operationTuple struct {
	Type OpType
	Data Operation
}

func (op *operationTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		op.Type,
		op.Data,
	})
}

func (op *operationTuple) UnmarshalJSON(data []byte) error {
	// The operation object is [opType, opBody].
	raw := make([]*json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Wrapf(err, "failed to unmarshal operation object: %v", string(data))
	}
	if len(raw) != 2 {
		return errors.Errorf("invalid operation object: %v", string(data))
	}

	// Unmarshal the type.
	var opType OpType
	if err := json.Unmarshal(*raw[0], &opType); err != nil {
		return errors.Wrapf(err, "failed to unmarshal Operation.Type: %v", string(*raw[0]))
	}

	// Unmarshal the data.
	var opData Operation
	template, ok := dataObjects[opType]
	if ok {
		opData = reflect.New(
			reflect.Indirect(reflect.ValueOf(template)).Type(),
		).Interface().(Operation)

		if err := json.Unmarshal(*raw[1], opData); err != nil {
			return errors.Wrapf(err, "failed to unmarshal Operation.Data: %v", string(*raw[1]))
		}
	} else {
		opData = &UnknownOperation{opType, raw[1]}
	}

	// Update fields.
	op.Type = opType
	op.Data = opData
	return nil
}
