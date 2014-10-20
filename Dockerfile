FROM progrium/busybox
MAINTAINER Matt Aitchison <matt@lanciv.com>

ADD ./stage/GoGradeMeAPI /bin/GoGradeMeAPI

EXPOSE 5005

ENTRYPOINT ["/bin/GoGradeMeAPI"]
# CMD ["-staticDir=/opt/www"]
