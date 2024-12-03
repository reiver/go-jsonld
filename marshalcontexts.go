package jsonld

import (
	"github.com/reiver/go-json"
)

func MarshalContexts(contexts ...Context) ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var squared bool = false
	{

		forImplicitNameSpaces(func(namespace string){
			if !squared {
				bytes = append(bytes, '[')
				squared = true
			} else {
				bytes = append(bytes, ',')
			}

			bytes = append(bytes, json.MarshalString(namespace)...)
		}, contexts...)
	}


	var curlyed bool = false
	{
		forExplicitNameSpaces(func(prefix string, namespace string){
			if !curlyed {
				if squared {
					bytes = append(bytes, ',')
				}
				bytes = append(bytes, '{')
				curlyed = true
			} else {
				bytes = append(bytes, ',')
			}

			bytes = append(bytes, json.MarshalString(prefix)...)
			bytes = append(bytes, ':')
			bytes = append(bytes, json.MarshalString(namespace)...)
		}, contexts...)
	}

	forContextNames(func(name string, prefix string){
		if !curlyed {
			if squared {
				bytes = append(bytes, ',')
			}
			bytes = append(bytes, '{')
			curlyed = true
		} else {
			bytes = append(bytes, ',')
		}

		bytes = append(bytes, json.MarshalString(name)...)
		bytes = append(bytes, ':')
		bytes = append(bytes, json.MarshalString(prefix+":"+name)...)
	}, contexts...)

	if curlyed {
		bytes = append(bytes, '}')
	}

	if squared {
		bytes = append(bytes, ']')
	}

	if len(bytes) <= 0 {
		return []byte(`{}`), nil
	}
	return bytes, nil
}
