package render

const errStart = "ERROR performing render: "

var err = struct {
	FieldNotFound   string
	UnsupportedType string
	UnsupportedMod  string
	Internal        string
}{
	"Field '%s' not found. Instead found fields: %v\n",
	"Field %s is of unsupported type '%s'. Supported types include: %v\n",
	"Field %s of type %s doesn't support modifier '%s'. Supported modifiers include: %v\n",
	"Field %s encountered an internal error: %v. Please open a github issue.\n",
}
