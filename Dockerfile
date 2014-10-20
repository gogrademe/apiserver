FROM progrium/busybox
MAINTAINER Matt Aitchison <matt@lanciv.com>

ADD ./stage/gogradeapi /bin/gogradeapi

EXPOSE 5005

ENTRYPOINT ["/bin/gogradeapi"]
# CMD ["-staticDir=/opt/www"]
