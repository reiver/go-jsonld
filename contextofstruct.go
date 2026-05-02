package jsonld

import (
	"reflect"
	"sort"
)

func contextOfStruct(value any) (Context, error) {
	if nil == value {
		var nada Context
		return nada, errNilValue
	}

	var reflectedType reflect.Type
	{
		reflectedType  = reflect.TypeOf(value)
		if nil == reflectedType {
			var nada Context
			return nada, errNilReflectedType
		}

		if reflect.Struct != reflectedType.Kind() {
			var nada Context
			return nada, errNotStruct
		}
	}

	var reflectedValue reflect.Value = reflect.ValueOf(value)

	var context Context
	collectContextFields(reflectedType, reflectedValue, &context)

	sort.Strings(context.Names)

	return context, nil
}

// collectContextFields recursively collects context information from struct fields, flattening anonymous (embedded) struct fields.
func collectContextFields(reflectedType reflect.Type, reflectedValue reflect.Value, context *Context) {
	var limit int = reflectedType.NumField()

	for index:=0; index<limit; index++ {

		var structField reflect.StructField = reflectedType.Field(index)
		var reflectedFieldValue reflect.Value = reflectedValue.Field(index)

		if !structField.IsExported() {
			continue
		}

		// If this is an anonymous (embedded) struct field, recurse into it.
		if structField.Anonymous && structField.Type.Kind() == reflect.Struct {
			// Check if it's a NameSpace or Prefix first.
			switch reflectedFieldValue.Interface().(type) {
			case NameSpace:
				value, found := structField.Tag.Lookup(structTagNameJSONLD)
				if found {
					context.NameSpace = value
				}
				continue
			case Prefix:
				value, found := structField.Tag.Lookup(structTagNameJSONLD)
				if found {
					context.Prefix = value
				}
				continue
			}

			collectContextFields(structField.Type, reflectedFieldValue, context)
			continue
		}

		switch reflectedFieldValue.Interface().(type) {
		case NameSpace:
			value, found := structField.Tag.Lookup(structTagNameJSONLD)
			if found {
				context.NameSpace = value
			}

		case Prefix:
			value, found := structField.Tag.Lookup(structTagNameJSONLD)
			if found {
				context.Prefix = value
			}

		default:
			name := structFieldName(structField)

			context.Names = append(context.Names, name)
		}
	}
}
