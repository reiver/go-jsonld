package jsonld

func forImplicitNameSpaces(fn func(string), contexts ...Context) {
	if nil == fn {
		return
	}

	for _, context := range contexts{

		if "" != context.Prefix {
			continue
		}
		if "" == context.NameSpace {
			continue
		}

		var implicitNamespace string = context.NameSpace

		fn(implicitNamespace)
	}
}
