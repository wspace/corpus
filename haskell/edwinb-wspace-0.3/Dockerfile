FROM i386/ubuntu

ADD https://web.archive.org/web/20150717214201id_/http://compsoc.dur.ac.uk:80/whitespace/downloads/wspace .
ADD https://web.archive.org/web/20160512184850id_/http://compsoc.dur.ac.uk/whitespace/downloads/wspace-0.3.tgz .
RUN chmod +x wspace
RUN tar xf wspace-0.3.tgz
WORKDIR /WSpace/examples
ENTRYPOINT ["/wspace"]
