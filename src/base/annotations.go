package base

import (
	"fmt"
	"reflect"
	"strings"
)

type Annotation interface {
	SetTag(tag reflect.StructTag)
	String() string
}

var AnnotationList []Annotation

func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnotationList {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}

func init() {
	AnnotationList = make([]Annotation, 0)
	AnnotationList = append(AnnotationList, new(Value))
}

type Value struct {
	tag         reflect.StructTag
	Beanfactory *BeanFactory
}

func (s *Value) SetTag(tag reflect.StructTag) {
	s.tag = tag
}

func (s *Value) String() string {
	getPrefix := s.tag.Get("prefix")
	if getPrefix == "" {
		return ""
	}
	prefix := strings.Split(getPrefix, ".")
	if config := s.Beanfactory.GetBean(new(SysConfig)); config != nil {
		getValue := GetConfigValue(config.(*SysConfig).Config, prefix, 0)
		if getValue != nil {
			return fmt.Sprintf("%v", getValue)
		} else {
			return ""
		}
	}
	return ""
}
