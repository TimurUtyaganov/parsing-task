package main

import "strings"

type Cond struct {
	Type      string
	Field     string
	Condition string
	Value     any
}

func Parser(condStr string) (*Cond, error) {
	delimIdx := strings.Index(condStr, ":")
	if delimIdx == -1 {
		return &Cond{Type: condStr[:1], Field: condStr[1:], Condition: "=", Value: nil}, nil
	}
	c := &Cond{
		Type:  condStr[:1],
		Field: condStr[1:delimIdx],
	}
	if condStr[delimIdx+1:delimIdx+2] < "0" || condStr[delimIdx+1:delimIdx+2] > "9" {
		condition := make([]byte, 0, 2)
		i := delimIdx + 1
		for i < len(condStr)-1 {
			if condStr[i:i+1] > "0" && condStr[i:i+1] < "9" {
				break
			}
			condition = append(condition, condStr[i])
			i++
		}
		c.Condition = string(condition)
		c.Value = condStr[i:]
	} else {
		c.Condition = "="
		c.Value = condStr[delimIdx+1:]
	}
	return c, nil
}
