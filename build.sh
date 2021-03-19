#git fetch
commit=$(git rev-parse --short HEAD)
#git checkout -f $commit
export CGO_ENABLED=0
go build -o test.bin *.go

if test $? -ne 0 ;then
	echo "go build error,exit"
	exit 1
fi

echo "[Current workdir]:"$PWD
imageName="harbor.gome.com.cn/test/gotest:$commit"
docker build -t $imageName .

if test $? -ne 0 ;then
        echo "docker build error,exit"
        exit 1
fi

docker push $imageName

if test $? -ne 0 ;then
        echo "docker push error,exit"
        exit 1
fi

sed  -i "s#- image: .*#- image: $imageName#" deploy.yaml 
sed -i "s#kubernetes.io/change-cause: .*#kubernetes.io/change-cause: $imageName#" deploy.yaml

kubectl apply -f deploy.yaml

