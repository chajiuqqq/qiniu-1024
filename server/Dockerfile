# 使用官方Golang基础镜像作为构建环境
# 您可以选择一个特定版本的golang镜像，如果需要的话
FROM golang:1.21 as builder

# 设置工作目录
WORKDIR /app

# 将源代码复制到容器中
COPY ./server/* .

# 获取依赖并缓存它们
# 假设您的项目使用go modules来管理依赖
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o api

# 使用scratch作为最小的基础镜像
FROM scratch

# 从builder阶段复制可执行文件
COPY --from=builder /app/api /

# 提供挂载点
VOLUME ["/conf"]

# 暴露需要的端口
EXPOSE 9133

# 运行你的go程序
# 我们假设您的应用程序可以使用-f标志接受配置文件路径
CMD ["/api", "-f", "/conf/config.yaml"]
