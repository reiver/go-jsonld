package jsonld

import (
	"slices"
	"sort"
	"reflect"
)

func deepContextsOfMany(values ...any) (contexts []Context, err error) {
	{
		var contextsMap map[string]*Context = map[string]*Context{}

		for _, value := range values {
			if reflect.Struct != reflect.TypeOf(value).Kind() {
				continue
			}

			ctxs, err := DeepContextsOf(value)
			if nil != err {
				return nil, err
			}

			for _, ctx := range ctxs {
				var ns string = ctx.NameSpace

				contextPointer, found := contextsMap[ns]
				if found {
					contextPointer.Names = append(contextPointer.Names)
				} else {
					contexts = append(contexts, ctx)
					contextsMap[ctx.NameSpace] = &(contexts[len(contexts)-1])
				}
			}
		}
	}
	{
		for _, context := range contexts {
			sort.Strings(context.Names)
			slices.Compact(context.Names)
		}
	}

	return contexts, nil
}
