ARG EXECUTABLE="vuln-docker-scanners"

# BASE IMAGE
# FROM golang:1.22.4-bookworm as BUILDER
FROM golang:1.22.4-alpine3.20 as BUILDER

ARG EXECUTABLE
ENV PROJECT=${EXECUTABLE}

RUN mkdir -p /$PROJECT


# TODO: fork repos to pin

# https://securitytrails.com/blog/nmap-vulnerability-scan

COPY . /$PROJECT

WORKDIR /$PROJECT

# TODO: add vulner scripts from repo

RUN go mod download && go mod verify

RUN go build -v -o /$PROJECT/bin/$PROJECT cmd/cli/main.go

# MULTI-STAGE BUILD
FROM alpine:3.20

ARG EXECUTABLE="vuln-docker-scanners"
ENV PROJECT=${EXECUTABLE}

# add nmap
# https://stackoverflow.com/questions/56446898/nmap-could-not-locate-nse-main-lua
RUN apk update && apk add nmap nmap-scripts git && rm -rf /var/cache/apk/*

RUN git clone https://github.com/scipag/vulscan.git /usr/share/nmap/scripts/vulscan
RUN rm -r /usr/share/nmap/scripts/vulscan/.git

RUN git clone https://github.com/vulnersCom/nmap-vulners.git /usr/share/nmap/scripts/nmap-vulners
RUN rm -r /usr/share/nmap/scripts/nmap-vulners/.git

# note, we don't set workdirs, github will do that on the checked out code
# we just create the directory to be able to mount anything here
RUN mkdir -p /app

COPY --from=BUILDER /${PROJECT}/bin/${PROJECT} /usr/local/bin/${PROJECT}

ENTRYPOINT ["sh", "-c", "\"$PROJECT\" \"$@\"", "--"]
