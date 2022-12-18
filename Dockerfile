FROM golang:1.18-alpine
WORKDIR /api
COPY . /api/project

RUN apk add --update make
RUN make -C project/ build
RUN cp project/api . && rm -rf project/

EXPOSE 1323
CMD ["./api"]