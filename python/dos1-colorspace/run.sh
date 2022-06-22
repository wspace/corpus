#!/bin/sh

usage() {
  echo 'Usage: <command> [...]

Commands:

    colorspace           -- Colorspace interpreter
    colorspace-decompile -- Whitespace to Colorspace translator
    colorspace-convert   -- Whitespace to Colorspace image converter' >&2

  exit 2
}

if [ $# = 0 ]; then
  usage
fi

command="$1"
shift
case "$command" in
colorspace)
  if [ $# != 1 ]; then
    echo 'Usage: colorspace <image>' >&2
    exit 2
  fi
  exec python colorspace "$@"
  ;;
colorspace-decompile)
  if [ $# != 2 ]; then
    echo 'Usage: colorspace-decompile <file> <output_image>' >&2
    exit 2
  fi
  exec python colorspace-decompile "$@"
  ;;
colorspace-convert)
  if [ $# != 3 ]; then
    echo 'Usage: colorspace-convert <file> <source_image> <output_image>' >&2
    exit 2
  fi
  exec python colorspace-convert "$@"
  ;;
*)
  usage
  ;;
esac
