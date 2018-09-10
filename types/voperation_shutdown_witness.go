package types

//ShutdownWitnessOperation represents shutdown_witness operation data.
type ShutdownWitnessOperation struct {
	Owner string `json:"owner"`
}

//Type function that defines the type of operation ShutdownWitnessOperation.
func (op *ShutdownWitnessOperation) Type() OpType {
	return TypeShutdownWitness
}

//Data returns the operation data ShutdownWitnessOperation.
func (op *ShutdownWitnessOperation) Data() interface{} {
	return op
}
