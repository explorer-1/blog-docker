FROM golang:1.22.5 as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /work/blog
COPY . /work/blog/
RUN go build -o app ./cmdd

EXPOSE 8000
ENTRYPOINT [ "./app" ]