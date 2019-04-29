FROM ubuntu:18.04

RUN apt-get update && \
apt-get -y upgrade
RUN apt-get -y install \
net-tools \
golang \
git \
vim \
mysql-client

# コンテナ内で必要なスクリプトを実行
COPY docker-entrypoint.sh /usr/local/bin
RUN chmod 777 /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]

RUN go get -u github.com/beego/bee && \
go get -u github.com/astaxie/beego && \
go get -u github.com/go-sql-driver/mysql && \
go get -u github.com/astaxie/beego/logs

CMD ["tail", "-f", "/dev/null"]
# TODO:バージョンの固定