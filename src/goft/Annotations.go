package goft

//
//import (
//	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/Injector"
//	"fmt"
//	"reflect"
//	"strings"
//)
//
//// Annotation 注解处理
//type Annotation interface {
//	SetTag(tag reflect.StructTag)
//	String() string
//}
//
//var AnnotationList []Annotation
//
//// IsAnnotation 判断当前的 注入对象 是否是注解
//func IsAnnotation(t reflect.Type) bool {
//	for _, item := range AnnotationList {
//		if reflect.TypeOf(item) == t {
//			return true
//		}
//	}
//	return false
//}
//func init() {
//	AnnotationList = make([]Annotation, 0)
//	AnnotationList = append(AnnotationList, new(Value))
//}
//
//type Value struct {
//	tag reflect.StructTag
//}
//
//func (this *Value) SetTag(tag reflect.StructTag) {
//	this.tag = tag
//}
//func (this *Value) String() string {
//	getPrefix := this.tag.Get("prefix")
//	if getPrefix == "" {
//		return ""
//	}
//	prefix := strings.Split(getPrefix, ".")
//	if config := Injector.BeanFactory.Get((*SysConfig)(nil)); config != nil {
//		getValue := GetConfigValue(config.(*SysConfig).Config, prefix, 0)
//		if getValue != nil {
//			return fmt.Sprintf("%v", getValue)
//		} else {
//			return ""
//		}
//	} else {
//		return ""
//	}
//}
