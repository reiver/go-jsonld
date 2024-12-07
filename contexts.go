package jsonld

import (
	"github.com/reiver/go-json"
)

type Contexts []Context

func (contexts Contexts) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var squared bool = false
	{
		for _, context := range contexts {
			ns, found := context.implicitNameSpace()
			if !found {
				continue
			}
			if "" == ns {
				continue
			}

			if !squared {
				bytes = append(bytes, '[')
				squared = true
			} else {
				bytes = append(bytes, ',')
			}

			bytes = append(bytes, json.MarshalString(ns)...)
		}
	}

	var curlyed bool = false
	{
		for _, context := range contexts {
			prefix, ns, found := context.explicitNameSpace()
			if !found {
				continue
			}
			if "" == ns {
				continue
			}
			if "" == prefix {
				continue
			}

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
			bytes = append(bytes, json.MarshalString(ns)...)
		}
	}

	{
		for _, context := range contexts {
			context.forNames(func(prefix string, name string){
				if "" == prefix {
					return
				}
				if "" == name {
					return
				}

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
			})
		}
	}

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
