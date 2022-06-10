FROM ubuntu

RUN apt-get update

# Rebol 3 alpha
#RUN apt-get install -y wget
#RUN wget https://rebolsource.net/downloads/experimental/r3-linux-x64-gbf237fc-static -O r3-alpha
#RUN chmod +x r3-alpha

# Ren-C
WORKDIR /home
RUN apt-get install -y git
RUN git clone https://github.com/metaeducation/ren-c --no-checkout

WORKDIR /home/ren-c/make
RUN git checkout 368bd5a724b7e84bf82cc0c588c56e634dd7e8cf
RUN cp r3-linux-x64-gbf237fc-static /usr/local/bin
RUN ln -s r3-linux-x64-gbf237fc-static r3-make
RUN apt-get install -y make gcc
RUN make -f makefile.boot
RUN mv r3 /usr/local/bin/r3-368bd5a7-from-gbf237fc

RUN ln -sf /usr/local/bin/r3-368bd5a7-from-gbf237fc r3-make
RUN make -f makefile.boot
RUN mv r3 /usr/local/bin/r3-368bd5a7-from-self

RUN git clean -dxf ..
RUN git checkout 8994d234e8a8be2962a4045dc5b4ff4805d8ad61
RUN r3-368bd5a7-from-self make.r

RUN git checkout 9403833acfc469432045408de532d5da88f6301f
RUN wget https://r3bootstraps.s3.amazonaws.com/r3-linux-x64-8994d23 -P prebuilt
RUN chmod +x prebuilt/r3-linux-x64-8994d23
