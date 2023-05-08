package protocol

type ObVString struct {
	stringVal   string
	bytesVal    []byte
	encodeBytes []byte
}

func (v *ObVString) String() string {
	return "ObVString{" +
		"stringVal:" + v.stringVal + ", " +
		"bytesVal:" + string(v.bytesVal) + ", " +
		"encodeBytes:" + string(v.encodeBytes) +
		"}"
}
