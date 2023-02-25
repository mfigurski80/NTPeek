package filter

import "fmt"

type condition struct {
	Simple *filter          `@@`
	Start  *condition       `| "(" @@`
	List   []*conditionList `@@? ")"`
}

type conditionList struct {
	AndConditions []*condition `( "AND" @@ )+`
	OrConditions  []*condition `| ( "OR" @@ )+`
}

func (c *condition) String() string {
	if c.Simple != nil {
		return c.Simple.String()
	}
	if c.Start != nil {
		return fmt.Sprintf("[%s]", c.Start)
	}
	return fmt.Sprintf("[%s%s]", c.Start, c.List)
}

func (c *conditionList) String() string {
	ret := ""
	iterable := c.AndConditions
	glueSymbol := " && "
	if len(c.AndConditions) == 0 {
		iterable = c.OrConditions
		glueSymbol = " || "
	}
	for i, cond := range iterable {
		ret += glueSymbol
		ret += cond.String()
	}
	return ret
}
