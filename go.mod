module github.com/infra-io/postar

go 1.21

//replace github.com/infra-io/servicex => ../servicex

replace github.com/infra-io/postar/api => ./api

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/FishGoddess/cryptox v0.4.0
	github.com/FishGoddess/errors v0.5.2
	github.com/FishGoddess/logit v1.5.10
	github.com/go-sql-driver/mysql v1.7.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	github.com/infra-io/postar/api v0.0.0-00010101000000-000000000000
	github.com/infra-io/servicex v0.3.1
	github.com/wneessen/go-mail v0.4.0
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	go.uber.org/automaxprocs v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240108191215-35c7eff3a6b1 // indirect
)
