#!/bin/bash

set -euo pipefail

if [ "$#" -ne 2 ]; then
  echo "usage: $0 <pending-delivery-file> <resume-pending>" >&2
  exit 2
fi

pending_file=$1
resume_pending=$2

case "$resume_pending" in
true) exit 0 ;;
false) ;;
*)
  echo "resume-pending must be true or false" >&2
  exit 2
  ;;
esac

# A changed path alone is insufficient: the durable receipt merge deletes the
# pending file. Only an extant pending delivery may resume normal generation.
[ -e "$pending_file" ] || exit 1
grep -Fxq -- "$pending_file"
