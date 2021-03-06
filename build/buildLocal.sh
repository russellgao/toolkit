## 在toolkit目录下运行

flags="-X 'github.com/prometheus/common/version.Version=`cat VERSION`' -X 'github.com/prometheus/common/version.BuildUser=russellgaorr@gmail.com' -X 'github.com/prometheus/common/version.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X 'github.com/prometheus/common/version.Branch=HEAD' -X 'github.com/prometheus/common/version.Revision=`git log -n1 --format=%H`'"
go build -ldflags="$flags" -o target/tkctl
cp target/tkctl $GOPATH/bin