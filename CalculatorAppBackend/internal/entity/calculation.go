package entity

import (
	"errors"
	"strings"
)

var (
	ErrEmptyID         = errors.New("ID cannot be empty")
	ErrEmptyExpression = errors.New("expression cannot be empty")
	ErrEmptyResult     = errors.New("result cannot be empty")
)

type Calculation struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

func (c *Calculation) Validate() error {
	if strings.TrimSpace(c.ID) == "" {
		return ErrEmptyID
	}
	if strings.TrimSpace(c.Expression) == "" {
		return ErrEmptyExpression
	}
	if strings.TrimSpace(c.Result) == "" {
		return ErrEmptyResult
	}
	return nil
}

func (cr *CalculationRequest) Validate() error {
	if strings.TrimSpace(cr.Expression) == "" {
		return ErrEmptyExpression
	}
	return nil
}
