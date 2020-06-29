# 编译阶段
FROM russellgao/golang:1.14.2 as builder
ADD toolkit /go/src/toolkit
WORKDIR /go/src/toolkit
RUN flags="-X 'github.com/prometheus/common/version.Version=`cat VERSION`' -X 'github.com/prometheus/common/version.BuildUser=gaoweizong@hd123.com' -X 'github.com/prometheus/common/version.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X 'github.com/prometheus/common/version.Branch=HEAD' -X 'github.com/prometheus/common/version.Revision=`git log -n1 --format=%H`'" \
    &&GOOS=linux GOARCH=amd64 go build -ldflags="$flags" -o /tmp/tkctl

# 运行阶段
FROM russellgao/alpine:3.9.4
LABEL PROJECT="toolkit" \
      VERSION="1.0.0"             \
      AUTHOR="gaoweizong@hd123.com"              \
      COMPANY="Shanghai HEADING Information Engineering Co., Ltd."
MAINTAINER gaoweizong "gaoweizong@hd123.com"
ENV LC_ALL en_US.UTF-8
COPY --from=builder /tmp/tkctl /bin
RUN chmod +x /bin/tkctl
CMD ["/bin/tkctl"]
