package jsonld

func MarshalContext(context Context) ([]byte, error) {
	return MarshalContexts(context)
}
