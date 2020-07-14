version=$1
set -ex
docker build -t russellgao/toolkit:$version -f build/Dockerfile .
docker push russellgao/toolkit:$version
