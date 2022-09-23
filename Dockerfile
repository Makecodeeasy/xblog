FROM ubuntu:14.04

WORKDIR /workspace

COPY xblog /workspace/xblog

CMD [ "./xblog" ]