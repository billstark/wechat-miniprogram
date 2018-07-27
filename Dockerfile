FROM golang

EXPOSE 8080
CMD /usr/local/bin/run
RUN apt-get update && apt-get install -y \
  supervisor \
  jq \
  nginx-extras \
  nginx \
  && ln -s /go/src/wechat-miniprogram/deploy/run /usr/local/bin/run
ADD . /go/src/wechat-miniprogram
WORKDIR /go/src/wechat-miniprogram
RUN cd /go/src/wechat-miniprogram \
  && go build -o /opt/wechat-miniprogram/wechat-miniprogram