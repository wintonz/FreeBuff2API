FROM golang:1.22-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go mod tidy && go build -o freebuff

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/freebuff .
# 关键：复制后台前端静态文件，解决404页面缺失
COPY --from=builder /app/web ./web
RUN apk add --no-cache ca-certificates tzdata

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
