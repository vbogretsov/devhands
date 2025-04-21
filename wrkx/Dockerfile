# FROM ubuntu:latest as build
# RUN apt update
# RUN apt install -y gcc make git libssl-dev zlib1g-dev
# RUN git clone https://github.com/vbogretsov/wrkx.git

# WORKDIR /wrkx
# RUN make

# FROM ubuntu:latest
# COPY --from=build /wrkx/wrk /usr/local/bin/wrk
# ENTRYPOINT [ "wrk" ]

FROM ubuntu:latest as build
RUN apt update
RUN apt install -y gcc make git libssl-dev zlib1g-dev
RUN git clone https://github.com/vbogretsov/wrkx.git

WORKDIR /wrkx
RUN git checkout luajit-2.1
RUN make

FROM ubuntu:latest
COPY --from=build /wrkx/wrk /usr/local/bin/wrk
ENTRYPOINT [ "wrk" ]

# FROM alpine:3.21 AS build
# RUN echo https://www.mirror.yandex.ru/mirrors/alpine/v3.21/main > /etc/apk/repositories
# RUN echo https://www.mirror.yandex.ru/mirrors/alpine/v3.21/community >> /etc/apk/repositories
# RUN apk update
# RUN apk add \
#     make \
#     gcc \
#     musl-dev \
#     bsd-compat-headers \
#     openssl-dev \
#     zlib-dev \
#     build-base \
#     git
# RUN git clone https://github.com/vbogretsov/wrkx.git
# WORKDIR /wrkx
# RUN git checkout luajit-2.1
# RUN make
# # RUN cp wrk /usr/local/bin/wrk

# FROM alpine:3.21
# COPY --from=build /wrkx/wrk /usr/local/bin/wrk
# RUN apk add --no-cache zlib openssl
# ENTRYPOINT [ "wrk" ]
