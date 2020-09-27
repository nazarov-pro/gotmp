FROM golang:alpine AS building
RUN apk update && apk add --no-cache git tzdata
ENV USER=appuser
ENV UID=10001 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/bettech.com/betting/scripts/

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/eventbook bettech.com/betting/scripts/cmd/eventbook

FROM scratch
COPY --from=building /etc/passwd /etc/passwd
COPY --from=building /etc/group /etc/group
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=building /go/bin/eventbook /app/


USER appuser:appuser

ENTRYPOINT ["/app/eventbook"]
