package jsonld

// Context represents a single namespace from a JSON-LD @context.
//
// Example:
//
//	context, err := jsonld.ContextOf(something)
//
// For multiple namespaces, see [Contexts].
type Context struct {
	NameSpace string
	Prefix string
	Names []string
}

// Ex:
//	["http://example.com/ns#"]
func (context Context) implicitNameSpace() (ns string, found bool) {

	if "" != context.Prefix {
		var nada string
		return nada, false
	}

	if "" == context.NameSpace {
		var nada string
		return nada, false
	}

	return context.NameSpace, true
}

// Ex:
//	{"something":"http://example.com/ns#"}
func (context Context) explicitNameSpace() (prefix string, ns string, found bool) {

	if "" == context.Prefix {
		var nada string
		return nada, nada, false
	}

	if "" == context.NameSpace {
		var nada string
		return nada, nada, false
	}

	return context.Prefix, context.NameSpace, true
}

// Ex:
//	"apple":"something:apple","banana":"something:banana","cherry":"something:cherry"
func (context Context) forNames(fn func(string,string)) {
	if nil == fn {
		return
	}

	if "" == context.Prefix {
		return
	}

	{
		var prefix string = context.Prefix

		for _, name := range context.Names {
			fn(prefix, name)
		}
	}
}
