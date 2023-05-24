FROM ubuntu:23.10 

RUN \
    apt update && \
    apt install -y apache2

EXPOSE 80

ENTRYPOINT ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]