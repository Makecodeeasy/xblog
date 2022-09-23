FROM ubuntu:14.04

WORKDIR /workspace

COPY main /workspace/main

CMD [ "./main" ]