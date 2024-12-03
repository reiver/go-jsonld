package jsonld

func forContextNames(fn func(string,string), contexts ...Context) {
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

		var prefix string = context.Prefix

		for _, name := range context.Names {
			fn(name, prefix)
		}
	}
}
