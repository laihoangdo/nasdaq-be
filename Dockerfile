# syntax=docker/dockerfile:1

###
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20.2-alpine as builder

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY ./go.mod .
COPY ./go.sum .

# copy directory files i.e all files ending with .go
COPY . .

# download Go modules and dependencies
RUN go mod download

# compile application
# /api: directory stores binaries file
RUN go build -o /api ./cmd/api/main.go

## Add Wkhtmltopdf
FROM alpine:3.14

RUN apk add --no-cache \
            xvfb \
            # Additionnal dependencies for better rendering
            ttf-freefont \
            fontconfig \
            dbus \
    && \
    # Install wkhtmltopdf from `testing` repository
    apk add qt5-qtbase-dev \
            wkhtmltopdf \
            --no-cache \
            --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ \
            --allow-untrusted \
    && \
    # Wrapper for xvfb
    mv /usr/bin/wkhtmltopdf /usr/bin/wkhtmltopdf-origin && \
    echo $'#!/usr/bin/env sh\n\
Xvfb :0 -screen 0 1024x768x24 -ac +extension GLX +render -noreset & \n\
DISPLAY=:0.0 wkhtmltopdf-origin $@ \n\
killall Xvfb\
' > /usr/bin/wkhtmltopdf && \
    chmod +x /usr/bin/wkhtmltopdf

RUN mv /usr/bin/wkhtmltoimage /usr/local/bin/
# RUN chmod 700 /app

WORKDIR /app

COPY --from=builder /api api
COPY --from=builder /app/start.sh ./

RUN apk add --no-cache tzdata 
ENV TZ=Asia/Ho_Chi_Minh

RUN chmod +x start.sh

CMD ["./start.sh"]



## STEP 2 - DEPLOY
##
#FROM scratch
#WORKDIR /app
#COPY --from=builder /api api

#ENTRYPOINT ["./api"]
