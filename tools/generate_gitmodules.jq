# Generate .gitmodules

def escape_space:
  if . | contains(" ") then "\(.)" else . end;

.[] |
.id as $id |
.submodules[]? |
"\($id)/\(.path)" as $path |

"[submodule \"\($path)\"]
\tpath = \($path | escape_space)
\turl = \(.url)",
if .branch != null then "\tbranch = \(.branch)" else empty end
