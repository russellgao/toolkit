# 运行阶段
FROM russellgao/alpine:3.9.4
LABEL PROJECT="toolkit" \
      VERSION="1.0.0"             \
      AUTHOR="gaoweizong@hd123.com"              \
      COMPANY="Shanghai HEADING Information Engineering Co., Ltd."
MAINTAINER gaoweizong "gaoweizong@hd123.com"
ENV LC_ALL en_US.UTF-8
ADD target/tkctl /bin
RUN chmod +x /bin/tkctl
CMD ["init"]
