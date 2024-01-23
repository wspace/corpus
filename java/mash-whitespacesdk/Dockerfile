FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /whitespacesdk
COPY whitespacesdk .
RUN <<EOF
set -e

javac -d bin -cp lib/commons-lang-2.4.jar $(find src -name '*.java')

# Merge whitespacesdk and commons-lang jars:
mkdir tmp
cd tmp
jar -xf ../lib/commons-lang-2.4.jar
rm -rf META-INF
cp -r ../bin/* .
jar -cfe ../whitespacesdk.jar uk.co.mash.whitespace.executable.WhitespaceProgramRunner .
cd ..
rm -rf tmp

EOF

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespacesdk
COPY --from=builder /whitespacesdk/whitespacesdk.jar .
COPY whitespacesdk/src/uk/co/mash/whitespace/example/ example
COPY run.sh .
ENTRYPOINT ["./run.sh"]
