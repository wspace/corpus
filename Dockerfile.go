FROM wspace-corpus/go/135yshr-wspacego AS go_135yshr-wspacego
FROM wspace-corpus/go/kinu AS go_kinu
FROM wspace-corpus/go/makiuchid-whitenote AS go_makiuchid-whitenote
FROM wspace-corpus/go/pocke-gows AS go_pocke-gows
FROM wspace-corpus/go/qeedquan AS go_qeedquan
FROM wspace-corpus/go/samuelpratt AS go_samuelpratt
FROM wspace-corpus/go/simomu AS go_simomu
FROM wspace-corpus/go/technohippy AS go_technohippy
FROM wspace-corpus/go/tempxla-go-wspace AS go_tempxla-go-wspace
# FROM wspace-corpus/go/thaliaarchi-nebula AS go_thaliaarchi-nebula
FROM wspace-corpus/go/zorchenhimer AS go_zorchenhimer

FROM scratch

WORKDIR /wspace
COPY --from=go_135yshr-wspacego     /wspacego      go/135yshr-wspacego/
COPY --from=go_kinu                 /whitespace    go/kinu/
COPY --from=go_makiuchid-whitenote  /wspace        go/makiuchid-whitenote/
COPY --from=go_pocke-gows           /gows          go/pocke-gows/
COPY --from=go_qeedquan             /whitespace    go/qeedquan/
COPY --from=go_samuelpratt          /whitespace-go go/samuelpratt/
COPY --from=go_simomu               /ws            go/simomu/
COPY --from=go_technohippy          /go-whitespace go/technohippy/
COPY --from=go_tempxla-go-wspace    /go-wspace     go/tempxla-go-wspace/
# COPY --from=go_thaliaarchi-nebula   /nebula        go/thaliaarchi-nebula/
COPY --from=go_zorchenhimer         /wi            go/zorchenhimer/
