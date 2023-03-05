package priority

type Priority uint8

const (
	LO Priority = iota
	MED
	HI
)

type TagsPriorityMap map[string]Priority

type PriorityConfig struct {
	Field   string
	Map     TagsPriorityMap
	Default Priority
}
