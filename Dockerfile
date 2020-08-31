FROM ubuntu:19.10
WORKDIR /usr/src/app
ADD ./source/server /usr/src/app
CMD ./server

FROM registry.cn-hangzhou.aliyuncs.com/michaelrepo/common/inbound-agent:v2
ENTRYPOINT "jenkins-agent"