package priority

const errStart = "Priority Assignment error: "

var errType = struct {
	FieldNotFound   string
	UnsupportedType string
}{
	FieldNotFound:   errStart + "field '%s' not found. Instead found fields: %v\n",
	UnsupportedType: errStart + "field '%s' is not a multi-select field\n",
}
