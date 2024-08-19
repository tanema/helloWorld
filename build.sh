set -euxo pipefail

mkdir -p functions
go build -o functions/hello .
chmod +x ./functions/hello
go env
