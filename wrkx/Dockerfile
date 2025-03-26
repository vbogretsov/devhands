FROM ubuntu:latest as build
RUN apt update
RUN apt install -y gcc make git libssl-dev zlib1g-dev
RUN git clone https://github.com/devhands-io/wrkx.git
WORKDIR /wrkx
RUN make

FROM ubuntu:latest
COPY --from=build /wrkx/wrk /usr/local/bin/wrk
ENTRYPOINT [ "wrk" ]
