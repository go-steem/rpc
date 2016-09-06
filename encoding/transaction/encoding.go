package transaction

type TransactionMarshaller interface {
	MarshalTransaction(*Encoder) error
}
