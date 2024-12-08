# 交叉编译构建镜像
docker buildx build --platform linux/amd64,linux/arm64 -t pgl888999/ip-server --push .
