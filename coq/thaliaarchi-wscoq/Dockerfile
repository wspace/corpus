FROM coqorg/coq

USER coq
COPY --chown=coq wscoq /home/coq/wscoq
WORKDIR /home/coq/wscoq
RUN make
