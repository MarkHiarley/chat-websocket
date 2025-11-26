package model

import (
	"gopkg.in/guregu/null.v4"
)

type LoginUser struct {
	Email    null.String `json:"email,omitempty" binding:"required"`
	Password null.String `json:"password,omitempty" binding:"required"`
}
