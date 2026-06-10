module github.com/cstrahan/packagecloud-go

go 1.25.0

require (
	github.com/cstrahan/packagecloud-go/sdk v0.0.0
	github.com/spf13/cobra v1.8.1
	github.com/vbauerster/mpb/v8 v8.12.1
	golang.org/x/term v0.44.0
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.23 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
)

replace github.com/cstrahan/packagecloud-go/sdk => ./sdk
