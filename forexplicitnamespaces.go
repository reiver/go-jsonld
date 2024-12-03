package jsonld

func forExplicitNameSpaces(fn func(string,string), contexts ...Context) {
	if nil == fn {
		return
	}

	for _, context := range contexts{

		if "" == context.Prefix {
			continue
		}
		if "" == context.NameSpace {
			continue
		}

		var explicitNamespace string = context.NameSpace
		var prefix string = context.Prefix

		fn(prefix, explicitNamespace)
	}
}
