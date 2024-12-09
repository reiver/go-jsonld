package jsonld


import (
	"slices"
	"sort"
	"reflect"

	"github.com/reiver/go-erorr"
)

func DeepContextsOf(value any) ([]Context, error) {
	if nil == value {
		return nil, errNilValue
	}

	var contexts []Context
	var contextsMap map[string]*Context = map[string]*Context{}

	{
		ctx, err := ContextOf(value)
		if nil != err {
			return nil, err
		}

		{
			var nada Context
			if !reflect.DeepEqual(nada, ctx) {
				contexts = append(contexts, ctx)
				contextsMap[ctx.NameSpace] = &(contexts[len(contexts)-1])
			}
		}
	}

	{
		var reflectedType reflect.Type = reflect.TypeOf(value)

		var kind reflect.Kind = reflectedType.Kind()
		if reflect.Struct != kind {
			return  nil, erorr.Errorf("jsonld: cannot get context-of something of kind %q", kind)
		}

		var reflectedStructValue reflect.Value = reflect.ValueOf(value)

		var limit int = reflectedStructValue.NumField()

		for index:=0; index<limit; index++ {

			var reflectedStructField reflect.StructField = reflectedType.Field(index)
			if reflect.Struct != reflectedStructField.Type.Kind() {
				continue
			}
			if !reflectedStructField.IsExported() {
				continue
			}

			var fieldReflectedValue reflect.Value = reflectedStructValue.Field(index)

			ctxs, err := DeepContextsOf(fieldReflectedValue.Interface())
			if nil != err {
				return nil, err
			}

			for _, ctx := range ctxs {
				var ns string = ctx.NameSpace

				contextPointer, found := contextsMap[ns]
				if found {
					contextPointer.Names = append(contextPointer.Names, ctx.Names...)
					sort.Strings(contextPointer.Names)
					slices.Compact(contextPointer.Names)
				} else {
					var nada Context
					if !reflect.DeepEqual(nada, ctx) {
						contexts = append(contexts, ctx)
						contextsMap[ctx.NameSpace] = &(contexts[len(contexts)-1])
					}
				}
			}
		}
	}

	return contexts, nil
}
