FROM progrium/busybox
MAINTAINER Matt Aitchison <matt@lanciv.com>

ADD ./stage/apiserver /bin/apiserver

EXPOSE 5005

ENTRYPOINT ["/bin/apiserver"]
# CMD ["-staticDir=/opt/www"]
