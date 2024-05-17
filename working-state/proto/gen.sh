set -e;

# Get the directory where the script is located
scriptDir="$( cd -- "$( dirname -- "${BASH_SOURCE[0]:-$0}"; )" &> /dev/null && pwd 2> /dev/null; )";
echo "script directory: $scriptDir";

outDir="$scriptDir/lang_go";
mkdir -p $outDir/primary;

protoc -I="$scriptDir" \
    --go_out="$outDir" --go_opt=paths=source_relative \
    --go-grpc_out="$outDir" --go-grpc_opt=paths=source_relative \
    dependencies/secondary/secondary.proto

protoc -I="$scriptDir" \
    --go_out="$outDir/primary" --go_opt=paths=source_relative \
    --go-grpc_out="$outDir/primary" --go-grpc_opt=paths=source_relative \
    primary.proto
