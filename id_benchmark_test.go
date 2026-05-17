package jsonld_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
)

func BenchmarkSomeID_alreadyNormalized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("https://example.com/path/to/resource")
	}
}

func BenchmarkSomeID_needsNormalization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("HTTPS://EXAMPLE.COM/path/to/resource")
	}
}

func BenchmarkSomeID_alreadyNormalized_idn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("http://درود.com/")
	}
}

func BenchmarkSomeID_needsNormalization_idn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("http://XN--UGBAF6G.COM/")
	}
}

func BenchmarkSomeID_alreadyNormalized_percentEncoded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("http://example.com/foo%2Fbar")
	}
}

func BenchmarkSomeID_needsNormalization_percentEncoded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("http://example.com/%E4%B8%96%E7%95%8C")
	}
}

func BenchmarkSomeID_blankNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("_:abc123")
	}
}

func BenchmarkSomeID_plainString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonld.SomeID("apple")
	}
}
