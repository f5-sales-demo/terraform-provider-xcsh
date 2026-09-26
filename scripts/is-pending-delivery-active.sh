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

# File existence is authoritative. A pending delivery must resume generation
# even when an unrelated commit did not modify the pending marker; otherwise
# the workflow can incorrectly authorize a direct release that the release
# validator must reject for lacking a regeneration receipt.
[ -e "$pending_file" ]
