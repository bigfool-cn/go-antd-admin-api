FROM golang:alpine as builder

# 开启Go Module, 设置GO Proxy代理
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 新建项目目录
RUN mkdir -p /react-antd-admin

# 指定工作目录
WORKDIR /react-antd-admin

# 复制源代码到工作目录
COPY . .

# 设置操作系统, 操作系统构架
RUN GOOS=linux GOARCH=amd64

# 删除exe文件
RUN rm -rf *.exe*

RUN go build -o server .

# 删除多余文件
RUN rm -rf `ls |egrep -v '(static|templates|server|configs)'` \
    && rm -rf `ls configs|egrep -v configs.yaml`

# 添加可执行权限
RUN chmod +x /react-antd-admin/server


FROM alpine

# MAINTAINER
LABEL name="react-antd-admin"
LABEL version="1.0.0"
LABEL author="bigfool <1063944784@qq.com>"
LABEL maintainer="bigfool <1063944784@qq.com>"
LABEL description="react-antd-admin application"

# 复制builder相关文件到基础镜像alpine
COPY --from=builder /react-antd-admin /react-antd-admin

# 设置时区
RUN apk add -U tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/lcoaltime \
    && echo 'Asia/Shanghai' > /etc/timezone

ENV TZ=Asia/Shanghai

# 新建一个用户www 并设置项目目录用户组
RUN adduser -D -H www \
    && chown -R www /react-antd-admin

# 执行用户
USER www

# 指定工作目录
WORKDIR /react-antd-admin

EXPOSE 8004

ENTRYPOINT ["./server"]
