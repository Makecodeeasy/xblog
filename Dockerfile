FROM ubuntu:14.04

WORKDIR /workspace

COPY xblog /workspace/xblog

RUN mkdir -p /workspace/config

CMD [ "./xblog" ]