FROM alpine

RUN apk add --no-cache lua5.4 lua5.4-dev luarocks5.4 gcc musl-dev wget
RUN luarocks-5.4 install luacov
RUN luarocks-5.4 install busted
WORKDIR /callisto
COPY callisto .
# Fix the require search paths by adding a missing semicolon between the
# prepended and existing paths
RUN find . -name '*.lua' -exec sed -i "s/'\.\.package\.path/;'..package.path/" \{} \;
RUN busted ./test
ENTRYPOINT ["lua5.4", "callisto.lua"]
