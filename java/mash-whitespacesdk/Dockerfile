FROM eclipse-temurin:20-jdk-jammy AS builder

WORKDIR /whitespacesdk
COPY whitespacesdk .
RUN javac -d bin -cp lib/commons-lang-2.4.jar $(find src -name '*.java')

FROM eclipse-temurin:20-jre-jammy

WORKDIR /whitespacesdk
COPY --from=builder /whitespacesdk/bin/ bin
COPY --from=builder /whitespacesdk/lib/ lib
COPY --from=builder /whitespacesdk/src/uk/co/mash/whitespace/example/ example
ENTRYPOINT ["java", "-cp", "bin:lib/commons-lang-2.4.jar", "uk.co.mash.whitespace.executable.WhitespaceProgramRunner"]
# ENTRYPOINT ["java", "-cp", "bin:lib/commons-lang-2.4.jar", "uk.co.mash.whitespace.executable.ListingFileProgramRunner"]
# ENTRYPOINT ["java", "-cp", "bin:lib/commons-lang-2.4.jar", "uk.co.mash.whitespace.executable.WhitespaceCompilerDebugger"]
# ENTRYPOINT ["java", "-cp", "bin:lib/commons-lang-2.4.jar", "uk.co.mash.whitespace.executable.CreateSimpleOutputFromString"]
