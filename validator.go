package ginvalidator

import (
	"sync"

	"gopkg.in/go-playground/validator.v9"
)

// ValidatorV9 for gin binding
type ValidatorV9 struct {
	validate    *validator.Validate
	once        sync.Once
	ConfigFn    func(*validator.Validate)
	initialized bool
}

// ValidateStruct any obj
func (v *ValidatorV9) ValidateStruct(obj interface{}) (err error) {
	v.lazyinit()
	return v.validate.Struct(obj)
}

// Engine gives validator engine
func (v *ValidatorV9) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *ValidatorV9) lazyinit() {
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
