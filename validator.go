package ginvalidator

import (
	"sync"

	"gopkg.in/go-playground/validator.v9"
)

// Validator for gin binding
type Validator struct {
	validate    *validator.Validate
	once        sync.Once
	ConfigFn    func(*validator.Validate)
	initialized bool
}

// ValidateStruct any obj
func (v *Validator) ValidateStruct(obj interface{}) (err error) {
	v.lazyinit()
	return v.validate.Struct(obj)
}

// Engine gives validator engine
func (v *Validator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *Validator) lazyinit() {
	if !v.initialized {
		v.initialized = true
		v.once.Do(func() {
			v.validate = validator.New()
			v.validate.SetTagName("binding")
			if v.ConfigFn != nil {
				v.ConfigFn(v.validate)
			}
		})
	}
}
