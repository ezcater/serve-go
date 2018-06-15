# multi-stage docker build
# build stage
FROM golang:latest AS build-env
ADD . /src/ch
ENV GOPATH=/

# check for dep binary so we dont run this
# check file || curl ....
#RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN cd /src/ch \
#    && dep ensure \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o certificate-host

# final stage
FROM alpine
WORKDIR /app
ADD ./files ./files
COPY --from=build-env /src/ch/certificate-host /app/
RUN pwd && ls -la .
ENTRYPOINT ./certificate-host
