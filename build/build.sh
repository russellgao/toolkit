version=$1
\cp toolkit/build/Dockerfile .
docker build -t russellgao/toolkit:$version ../
docker push russellgao/toolkit:$version
