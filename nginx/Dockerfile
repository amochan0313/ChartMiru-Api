FROM ubuntu:18.04

RUN apt-get update && \
apt-get -y upgrade
RUN apt-get -y install \
nginx \
net-tools \
vim \
iputils-ping

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]