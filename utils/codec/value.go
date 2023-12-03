package codec

type ValuePtr struct {
}

func NewValuePtr(entry *Entry) *ValuePtr {
	return &ValuePtr{}
}

func IsValuePtr(entry *Entry) bool {
	return false
}

func ValuePtrDecode(data []byte) *ValuePtr {
	return nil
}
