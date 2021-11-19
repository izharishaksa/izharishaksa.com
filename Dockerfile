FROM golang:1.17.3-alpine3.14 AS builder

WORKDIR ./app
COPY . ./

RUN apk update && apk add --no-cache git
RUN go mod download
RUN go build -o /usr/bin/izharishaksa-com ./cmd

############################### SECOND STAGE ###############################

FROM alpine

WORKDIR ./app
COPY internal ./internal

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

COPY --from=builder /usr/bin/izharishaksa-com /usr/bin/izharishaksa-com

EXPOSE 900

CMD ["izharishaksa-com"]
