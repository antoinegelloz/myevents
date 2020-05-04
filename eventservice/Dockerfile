FROM golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY eventservice /
ADD certificate /certificate
EXPOSE 8181
CMD ["./eventservice"]
