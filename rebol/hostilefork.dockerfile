FROM ubuntu

RUN apt-get update
RUN apt-get install -y wget git make gcc

# Pre-built Rebol 3 alpha bf237fc
RUN wget https://rebolsource.net/downloads/experimental/r3-linux-x64-gbf237fc-static -P /usr/local/bin

# Pre-built Ren-C 8994d23
RUN wget https://r3bootstraps.s3.amazonaws.com/r3-linux-x64-8994d23 -P /usr/local/bin
RUN wget http://hostilefork.com/media/shared/github/r3-linux-8994d23-patched -P /usr/local/bin
RUN chmod +x /usr/local/bin/r3-*

# Ren-C
WORKDIR /home
RUN git clone https://github.com/metaeducation/ren-c --no-checkout

# Build 368bd5a from pre-built bf237fc
WORKDIR /home/ren-c/make
RUN git checkout 368bd5a
RUN cmp r3-linux-x64-gbf237fc-static /usr/local/bin/r3-linux-x64-gbf237fc-static
RUN ln -s r3-linux-x64-gbf237fc-static r3-make
RUN make -f makefile.boot
RUN mv r3 /usr/local/bin/r3-368bd5a-from-bf237fc

# Build 368bd5a from itself
RUN ln -sf /usr/local/bin/r3-368bd5a-from-bf237fc r3-make
RUN make -f makefile.boot
RUN mv r3 /usr/local/bin/r3-368bd5a-from-self

# Build 8994d23 from pre-built bf237fc
RUN git clean -dxf ..
RUN git checkout 8994d23
RUN r3-linux-x64-gbf237fc-static make.r
RUN mv r3 /usr/local/bin/r3-8994d23-from-bf237fc

# Build 8994d23 from itself
RUN r3-8994d23-from-bf237fc make.r
RUN mv r3 /usr/local/bin/r3-8994d23-from-self

# Build 8994d23 from 368bd5a
# RUN git clean -dxf ..
# RUN git checkout 8994d23
# RUN r3-368bd5a-from-self make.r
# RUN mv r3 /usr/local/bin/r3-8994d23-from-368bd5a

# Build 8994d23 from itself
# RUN r3-8994d23-from-368bd5a make.r
# RUN mv r3 /usr/local/bin/r3-8994d23-from-self

# Build 9403833 from pre-built bf237fc
RUN git clean -dxf ..
RUN git checkout 9403833
RUN r3-linux-x64-gbf237fc-static make.r
