module github.com/example/blog

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.6
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.1 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.0.0-20210625155357-5a2c8d79013b
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/tools v0.1.3 // indirect
	google.golang.org/genproto v0.0.0-20211129164237-f09f9a12af12
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
