FROM golang:1.22.1-alpine AS builder
WORKDIR /build/
RUN rm -rf ./*
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -ldflags '-s' -installsuffix cgo -o main .

# FINAL STAGE
FROM alpine:3.15 AS runner
WORKDIR /app/
RUN mkdir -p /logs && chmod 755 /logs
RUN apk add --no-cache curl tar \
    && curl -L https://github.com/grafana/loki/releases/download/v2.9.2/promtail-linux-arm64-v2.9.2.tar.gz -o /promtail.tar.gz \
    && tar -xzf /promtail.tar.gz -C /usr/local/bin \
    && chmod +x /usr/local/bin/promtail
COPY --from=builder /build/main .
RUN echo 'export $(strings /proc/1/environ | grep AWS_CONTAINER_CREDENTIALS_RELATIVE_URI)' >> /root/.profile
ARG VERSION
ENV VER $VERSION
EXPOSE 80
EXPOSE 3100
COPY promtail-config.yaml /etc/promtail/promtail-config.yaml
CMD ["sh", "-c", "promtail -config.file=/etc/promtail/promtail-config.yaml & ./main"]