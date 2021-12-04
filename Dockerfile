FROM alpine

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
COPY . /usr/bin

EXPOSE 900

CMD ["./izharishaksa-com"]
