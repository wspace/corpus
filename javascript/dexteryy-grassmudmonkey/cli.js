var fs = require('fs');

function usage() {
  process.stdout.write("usage: grassmudmonkey [-w | --whitespace] <file>\n");
  process.stdout.write("\n");
  process.stdout.write("    -w, --whitespace\n");
  process.stdout.write("        Use Whitespace syntax instead of GrassMudMonkey syntax\n");
  process.stdout.write("    <file>\n");
  process.stdout.write("        Filename of the program\n");
  process.exit(2);
}

let type = 'GrassMudHorse';
let filename = null;
for (let i = 2; i < process.argv.length; i++) {
  const arg = process.argv[i];
  if (arg === "--whitespace" || arg == '-w') {
    type = 'Whitespace';
  } else {
    if (filename !== null) {
      usage();
    }
    filename = arg;
  }
}
if (filename === null) {
  usage();
}

let program;
try {
  program = fs.readFileSync(filename, "utf-8");
} catch (e) {
  process.stdout.write(e.toString());
  process.stdout.write("\n");
  process.exit(1);
}

GrassMudMonkey.type = type;
GrassMudMonkey.print = s => process.stdout.write(s);
GrassMudMonkey.eval(program)
