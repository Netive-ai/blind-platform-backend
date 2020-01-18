sudo mkdir /app
sudo mkdir /app/cfg
sudo cp cfg/* /app/cfg
go test github.com/blind-platform/cmd/platform/
go test github.com/blind-platform/pkg/api
go test github.com/blind-platform/pkg/conf
