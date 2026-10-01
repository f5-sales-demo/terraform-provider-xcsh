# pylint: disable=invalid-name
# ruff: noqa: INP001
"""Generate resource import scripts from the shared importer contract."""

import json
from pathlib import Path

from import_contract import resolve_import_contract


def generate(root: Path) -> None:
    """Resolve every supported resource before writing any import output."""
    surface = json.loads((root / "provider-release-surface.json").read_text())
    contracts = {}
    for name in surface["resources"]:
        source = root / "internal/provider" / f"{name}_resource.go"
        contracts[name] = resolve_import_contract(source.read_text(), name)
    for name, contract in contracts.items():
        target = root / "examples/resources" / f"xcsh_{name}" / "import.sh"
        if contract is None:
            target.unlink(missing_ok=True)
            continue
        target.parent.mkdir(parents=True, exist_ok=True)
        guidance = contract["guidance"].replace("`", "")
        target.write_text(
            "#!/usr/bin/env bash\n# " + guidance + "\n" + contract["command"]
        )


if __name__ == "__main__":
    generate(Path.cwd())
