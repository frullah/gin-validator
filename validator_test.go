package ginvalidator

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name string `json:"name" binding:"required"`
}

func Test_Engine_engineEqualToValidator(t *testing.T) {
	engine := Validator{}
	require.Equal(t, engine.Engine(), engine.validate)
}

func Test_ValidateStruct_valid(t *testing.T) {
	engine := Validator{}
	data := User{Name: "fajar"}
	require.NoError(t, engine.ValidateStruct(&data))
}

func Test_ValidateStruct_invalid(t *testing.T) {
	engine := Validator{}
	data := User{Name: ""}
	require.Error(t, engine.ValidateStruct(&data))
}

func Test_ValidateStruct_jsonTag(t *testing.T) {
	engine := Validator{}
	data := User{Name: ""}
	err := engine.ValidateStruct(&data)
	require.Error(t, err)
	errors := err.(validator.ValidationErrors)
	require.Equal(t, errors[0].Field(), "name")
}
func Test_ValidateStruct_emptyJsonTag(t *testing.T) {
	engine := Validator{}
	data := struct {
		Name string `binding:"required"`
	}{Name: ""}
	err := engine.ValidateStruct(&data)
	require.Error(t, err)
	errors := err.(validator.ValidationErrors)
	require.Equal(t, errors[0].Field(), "Name")
}

func Test_ValidateStruct_dashJsonName(t *testing.T) {
	engine := Validator{}
	data := struct {
		Name string `json:"-" binding:"required"`
	}{Name: ""}
	err := engine.ValidateStruct(&data)
	require.Error(t, err)
	errors := err.(validator.ValidationErrors)
	require.Equal(t, errors[0].Field(), "Name")
}

func Test_lazyInit(t *testing.T) {
	counter := 0
	engine := Validator{ConfigFn: func(v *validator.Validate) { counter++ }}
	engine.lazyinit()
	require.True(t, engine.initialized)

	engine.lazyinit()
	require.Equal(t, 1, counter)
}
