package transaction

//TransactionMarshaller interface for converting data into byte
type TransactionMarshaller interface {
	MarshalTransaction(*Encoder) error
}
