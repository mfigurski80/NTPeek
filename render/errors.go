package render

const errStart = "ERROR performing render: "

var errType = struct {
	FieldNotFound   string
	UnsupportedType string
	UnsupportedMod  string
	Internal        string
}{
	errStart + "Field '%s' not found. Instead found fields: %v\n",
	errStart + "Field %s is of unsupported type '%s'. Supported types include: %v\n",
	errStart + "Field %s of type %s doesn't support modifier '%s'. Supported modifiers include: %v\n",
	errStart + "Field %s encountered an internal error: %v. Please open a github issue.\n",
}
