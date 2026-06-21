# 升级Go版本到1.25，满足项目最低编译要求
FROM golang:1.25-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go mod tidy && go build -o freebuff

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/freebuff .
# 复制后台web前端文件夹，解决/admin 404
COPY --from=builder /app/web ./web
RUN apk add --no-cache ca-certificates tzdata

# 启动脚本（你是硬盘挂载方案，环境变量无需填真实值）
COPY <<-'EOF' /app/start.sh
#!/bin/sh
cat > /app/config.yaml << YAML
auth:
  api_keys:
    - ${CB_API_TOKEN}
server:
  listen: :8080
  api_keys:
    - ${CLIENT_API_KEY}
limits:
  global_rpm: 60
YAML
echo "${ADMIN_PASSWORD}" > /app/token.key
mkdir -p /app/auths
exec ./freebuff
EOF
RUN chmod +x /app/start.sh

EXPOSE 8080
CMD ["/app/start.sh"]
