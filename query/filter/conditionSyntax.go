package filter

import "fmt"

type condition struct {
	Simple *filter        `@@`
	Start  *condition     `| "(" @@`
	List   *conditionList `@@? ")"`
}

type conditionList struct {
	AndConditions []*condition `( "AND" @@ )+`
	OrConditions  []*condition `| ( "OR" @@ )+`
}

/// STRINGIFY

func (c *condition) String() string {
	if c.Simple != nil {
		return c.Simple.String()
	}
	if c.List == nil {
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
	for _, cond := range iterable {
		ret += glueSymbol
		ret += cond.String()
	}
	return ret
}

/// RENDER FUNCTION

func (c *condition) Render() string {
	if c.Simple != nil {
		return c.Simple.Render()
	} else if c.List == nil {
		return c.Start.Render()
	}
	keyword := "and"
	conds := c.List.AndConditions
	if len(conds) == 0 {
		keyword = "or"
		conds = c.List.OrConditions
	}
	ret := `{"` + keyword + `": [` + c.Start.Render() + ", "
	for _, cond := range conds {
		ret += cond.Render() + ", "
	}
	ret = ret[:len(ret)-2] + "]}" // remove last comma
	return ret
}
