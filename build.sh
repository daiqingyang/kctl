git fetch
commit=$(git rev-parse --short origin/HEAD)
git checkout -f $commit
export CGO_ENABLED=0
go build -o test.bin *.go
echo $PWD
docker build -t test:$commit .

