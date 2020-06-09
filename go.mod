module web1992

go 1.13

require (
	github.com/coreos/go-semver v0.3.0
	gitlab.com/golang-commonmark/markdown v0.0.0-20191127184510-91b5b3c99c19
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

// old/thing v1.2.3
// exclude

// bad/thing v1.4.5 => good/thing v1.4.5
// replace
