#!/bin/bash
# Convert a Dockerfile to an Earthfile \
if (( "$#" != 1 )); then echo "Usage: dockerfile_to_earthfile.jq [Dockerfile]" >&2; fi; exec jq -Rrsf "$0" -- "${@:1}"

(input_filename // "File") as $filename |
[scan("(?:\\\\*[^\\\\\n]|\\\\\n)*(?:\n|$)") | sub("\n$"; "")] |
. as $lines |
# Desugar heredocs
(reduce range(0; length) as $i ([];
  def prev: .[-1];
  $lines[$i] as $line |
  if prev | type == "object" then
    if $line == prev.heredoc then
      # End of heredoc
      prev |= (.data | "RUN " + sub("(?:\\s*\\\\\n)+$"; ""; ""))
    else
      # Middle of heredoc
      prev.data +=
        if prev.data == "" then "" else "    " end
        + $line
        + if $line == "" then "" else " " end
        + "\\\n"
    end
  elif ($line | test("^\\s*#")) and (prev | . != null and test("^\\s*#")) then
    # Join adjacent comments
    prev += "\n" + $line
  else
    # Start of heredoc
    ($line | first(capture("^RUN <<(?<term>.+)$").term, null)) as $heredoc |
    . + [if $heredoc then {$heredoc, data: ""} else $line end]
  end
)) as $lines |
[
  # Group lines into FROM stages
  range(0; $lines | length) as $i |
  if $lines[$i] | test("^FROM ") then
    $i | if $i > 0 and ($lines[.-1] | test("^\\s*#")) then .-1 end
  else empty end
] as $ranges |
$ranges |
if $ranges[0] != 0 then
  "Error: \($filename) does not start with FROM\n" | halt_error(1)
end |
[
  range(0; $ranges | length) as $i |
  $lines[$ranges[$i] : $ranges[$i+1]] as $stage | $stage |
  # Strip empty lines after FROM and at the end
  first(range(0; length) | select($stage[.] | test("^FROM "))) as $from |
  first((.[$from] | capture("^FROM [^ ]+ AS (?<name>[^ ]+)$").name), null) as $name |
  .[$from] |= sub(" AS [^ ]+$"; "") |
  if .[$from+1] == "" then .[:$from+1] + .[$from+2:] end |
  if .[-1] == "" then .[:-1] end |
  {stage: $name, commands: .}
] |
map(.stage |= if . == "builder" then "build" end) |
# Add implicit stage names
if length == 1 and .[0].stage == null then
  .[0].stage =
    if isempty(.[0].commands[] | select(test("^ENTRYPOINT "))) then "build"
    else "docker" end
elif length == 2 and .[1].stage == null and .[0].stage == "build" then
  .[1].stage = "docker"
elif length > 2 and .[-1].stage == null then
  .[-1].stage = "docker"
end |
# Convert COPY to SAVE ARTIFACT
if .[-2].stage? == "build" and .[-1].stage? == "docker" then
  .[-2].commands += [
    .[-1].commands[] |
    select(test("^COPY ")) |
    gsub("\\s*\\\\\n\\s*"; " ") |
    first(
      (capture("^COPY(?: --from=(?<from>builder))? (?<paths>\\[.+\\])$") |
        .paths | fromjson | {src: .[:-1], dest: .[-1]}),
      (capture("^COPY(?: --from=(?<from>builder))? (?<src>[^ ]+(?: [^ ]+)*) (?<dest>[^ ]+)$") |
        {src: .src | split(" "), dest}),
      ("Error: \($filename): Unrecognized COPY: \(.)\n" | halt_error(1))
    ) |
    .src = (.src[] | sub("^/?[^/]+/"; "")) |
    .dest |= if . == "." or . == "./" then "/" end |
    # Classify some files
    if .dest == "/" then
      .dest =
        if .src | test("(?:^|/)(?:programs|samples|examples|tests?|resources)$") then "/programs"
        elif .src | test("\\.(?:ws|wsa|txt|asm|wsp|whitespace)$") then "/programs/"
        elif .src | test("(?:^|/)(bin|target|out)$") then "/bin"
        elif .src | test("\\.(?:out|sh|bash|jar|exe|dll)$|(?:/|^)[^.*]+$") then "/bin/"
        else .dest end
    elif (.dest | test("^/[^/]+$")) and (.dest == .src or .dest == "/\(.src)") then
      .dest =
        if .src | test("(?:^|/)(?:programs|samples|examples|tests?|resources)$") then "/programs"
        else .dest end
    end |
    if (.src | contains(" ")) or (.dest | contains(" ")) then
      "SAVE ARTIFACT [\(.src | tojson), \(.dest | tojson)]"
    else
      "SAVE ARTIFACT \(.src) \(.dest)"
    end
  ] |
  (if isempty(.[-1].commands[] | select(test("^WORKDIR ")))
    then "/" else "." end) as $dest |
  .[-1].commands |= (
    first(first(.[] | select(test("^COPY "))) = "COPY +build/ \($dest)", .) |
    map(select(test("^COPY [^+]") | not)) |
    # Move entrypoint into bin/
    map(
      . as $line |
      sub("^ENTRYPOINT (?<args>\\[.+\\])$";
        .args | fromjson | . as $args |
        (if .[0] | contains("/") then 0
          else
            first(
              (range(1; length) | select($args[.] | test("^-") | not)),
              ("Error: \($filename): Unable to identify entrypoint: \($line)\n" | halt_error(1))
            )
          end) as $bin_i |
        .[$bin_i] |=
          if test("^\\.?/") then
            sub("^(?<dot>\\.?)(?<root>(?:/[^/]+)?)/";
              if .dot == "." and .root == "" then "bin/"
              else "\(.dot)\(.root)/bin/" end
            )
          else "bin/" + . end |
        "ENTRYPOINT [\(map(tojson) | join(", "))]"
      )
    )
  )
end |
if length == 1 and .[0].stage == "build" then
  . + [{stage: "docker", commands: []}]
end |
.[-1].commands +=
  ["SAVE IMAGE wspace-corpus/" + ($filename | sub("/(?:Dockerfile|Earthfile)$"; ""))] |
map(
  [
    "\(.stage):",
    (.commands[] | split("\n")[] | if . != "" then "    " + . end)
  ] |
  join("\n")
) |
"VERSION 0.8\n\n" + join("\n\n")
