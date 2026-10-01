# ruff: noqa: INP001
"""Resolve documentation import identities from generated importer declarations."""

import re

FORMATS = {
    "name",
    "namespace/name",
    "namespace/name/allowed_domain",
    "namespace/name/mitigated_domain",
    "namespace/name/protected_domain",
}


def resolve_import_contract(text: str, name: str) -> dict[str, str] | None:
    """Fail closed for missing or unsupported importer metadata."""
    if not re.search(r"func\s+\([^)]*\)\s+ImportState\s*\(", text):
        return None
    matches = re.findall(r"// Import ID format: ([^\n]+)", text)
    if len(matches) != 1:
        message = f"import supported without unique documented source syntax: {name}"
        raise ValueError(message)
    syntax = matches[0].strip().split(" ", 1)[0]
    if syntax not in FORMATS:
        message = f"unrecognized import syntax: {name}: {syntax}"
        raise ValueError(message)
    if syntax == "name":
        identity = "example-" + name.replace("_", "-")
        address = f"xcsh_{name}.this"
        guidance = "This tenant-level resource uses its bare name. The `namespace` argument is omitted."
    else:
        identity = "/".join(
            "system" if part == "namespace" else "example" for part in syntax.split("/")
        )
        address = f"xcsh_{name}.example"
        guidance = f"Import using the `{syntax}` identifier format."
    command = f"terraform import {address} {identity}\n"
    return {"syntax": syntax, "command": command, "guidance": guidance}
