FROM ubuntu:19.10
WORKDIR /usr/src/app
ADD server /usr/src/app
CMD ./server