package ginvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name string `binding:"required"`
}

func TestEngine_engineEqualToValidator(t *testing.T) {
	engine := ValidatorV9{}
	assert.Equal(t, engine.Engine(), engine.validate)
}

func TestValidateStruct_valid(t *testing.T) {
	engine := ValidatorV9{}
	data := User{Name: "fajar"}
	assert.NoError(t, engine.ValidateStruct(&data))
}

func TestValidateStruct_invalid(t *testing.T) {
	engine := ValidatorV9{}
	data := User{Name: ""}
	assert.Error(t, engine.ValidateStruct(&data))
}

func Test_lazyInit(t *testing.T) {
	counter := 0
	engine := ValidatorV9{ConfigFn: func(v *validator.Validate) { counter++ }}
	engine.lazyinit()
	assert.True(t, engine.initialized)

	engine.lazyinit()
	assert.Equal(t, 1, counter)
}
