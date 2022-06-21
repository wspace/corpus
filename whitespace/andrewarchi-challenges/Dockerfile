FROM wspace-corpus/rust/censoredusername-whitespacers AS whitespacers

FROM alpine AS builder

RUN apk add git make bash coreutils
COPY --from=whitespacers /wsc /usr/local/bin/
RUN git clone https://github.com/andrewarchi/wslib
WORKDIR /ws-challenges
COPY ws-challenges .
RUN make -k COMPILED_PROGRAMS= all run_tests || :
RUN ./test.bash
