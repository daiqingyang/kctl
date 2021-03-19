git fetch
commit=$(git rev-parse --short origin/HEAD)
git checkout -f $commit
export CGO_ENABLED=0
go build -o test.bin *.go
echo $PWD
imageName="harbor.gome.com.cn/test/gotest:$commit"
docker build -t $imageName .
docker push $imageName
sed  -i "s#- image: .*#- image: $imageName#" deploy.yaml 
sed -i "s#kubernetes.io/change-cause: .*#kubernetes.io/change-cause: $imageName#" deploy.yaml
kubectl apply -f deploy.yaml

