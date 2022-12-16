module github.com/mhuggins7278/AdventOfCode/main

go 1.19

require github.com/mhuggins7278/AdventOfCode/utils v0.0.0

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/k0kubun/pp v3.0.1+incompatible
	golang.org/x/exp v0.0.0-20221126150942-6ab00d035af9
)

require (
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace github.com/mhuggins7278/AdventOfCode/utils v0.0.0 => ../utils
