FROM ubuntu:18.04

RUN apt-get update && \
apt-get -y upgrade
RUN apt-get -y install \
net-tools \
golang \
nginx \
git

# コンテナ内で必要なスクリプトを実行
COPY docker-entrypoint.sh /usr/local/bin
RUN chmod 777 /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]

RUN go get -u github.com/beego/bee && \
go get -u github.com/astaxie/beego

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
# TODO:バージョンの固定