#==================================#
#         FRONTEND BUILD           #
#==================================#
FROM node:12-alpine as frontend

WORKDIR /usr/src

COPY ./frontend/package.json .

RUN yarn install

COPY ./frontend .

RUN yarn build

#==================================#
#          BACKEND BUILD           #
#==================================#
FROM golang:1.13 as builder

WORKDIR /go-n-reactjs

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN rm -rf ./frontend

RUN mkdir -p ./frontend

COPY --from=frontend /usr/src/build ./frontend/build

RUN go get github.com/GeertJohan/go.rice && \
    go get github.com/GeertJohan/go.rice/rice && \
    rice embed-go && \
    CGO_ENABLED=0 GOOS=linux go build -o app .

#==================================#
#          COPY BIN FILE           #
#==================================#
FROM alpine:3.6

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Ho_Chi_Minh

COPY --from=builder /go-n-reactjs/app /go/bin/app

EXPOSE 8080

CMD ["/go/bin/app"]

