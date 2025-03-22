module github.com/karagatandev/cligo-example

go 1.23.0

toolchain go1.24.1

//replace go.arpabet.com/cligo => ../cligo

require (
	go.arpabet.com/cligo v0.1.2
	go.arpabet.com/glue v1.1.2
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
