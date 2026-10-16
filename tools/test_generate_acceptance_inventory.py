#!/usr/bin/env python3
"""Prove native ownership and fail-closed inventory generation without dependencies."""

import json
import pathlib
import tempfile

from generate_acceptance_inventory import EXPLICIT_EXCLUSIONS, generate


def _rejects(root: pathlib.Path, message: str) -> None:
    try:
        generate(root)
    except RuntimeError as error:
        if message not in str(error):
            raise AssertionError from error
    else:
        raise AssertionError


def verify_contracts() -> None:
    """Exercise the native client and reject missing or changed contracts."""
    with tempfile.TemporaryDirectory() as directory:
        root = pathlib.Path(directory)
        (root / "internal/client").mkdir(parents=True)
        (root / "internal/acctest").mkdir()
        surface = root / "provider-release-surface.json"
        native = root / "internal/client/secret_policy.go"
        source = (
            pathlib.Path(__file__).resolve().parents[1]
            / "internal/client/secret_policy.go"
        )
        original = source.read_text()
        surface.write_text(
            json.dumps({"resources": [*EXPLICIT_EXCLUSIONS, "secret_policy"]})
        )
        native.write_text(original)
        generate(root)
        inventory = json.loads(
            (root / "internal/acctest/ownership_handler_exclusions.json").read_text()
        )
        if not {
            entry["type"]: entry["handler"] for entry in inventory["resources"]
        } == {name: contract[0] for name, contract in EXPLICIT_EXCLUSIONS.items()}:
            raise AssertionError
        handlers = (
            root / "internal/acctest/ownership_handlers_generated.go"
        ).read_text()
        if "/api/secret_management/namespaces/%s/secret_policys/%s" not in handlers:
            raise AssertionError
        if "c.GetSecretPolicy(ctx, ns, name)" not in handlers:
            raise AssertionError
        if "c.DeleteSecretPolicy(ctx, ns, name)" not in handlers:
            raise AssertionError
        if "readOwnershipProjection" not in handlers:
            raise AssertionError
        if not "/api/config/" not in handlers:
            raise AssertionError
        surface.write_text(json.dumps({"resources": ["unreviewed"]}))
        _rejects(root, "unreviewed ownership lifecycle")
        (root / "internal/client/unreviewed_types.go").write_text("package client\n")
        _rejects(root, "unreviewed ownership signatures")
        surface.write_text(json.dumps({"resources": ["secret_policy"]}))
        native.unlink()
        _rejects(root, "unreviewed ownership lifecycle")
        native.write_text(
            original.replace("DeleteSecretPolicy(ctx", "DeleteDifferentPolicy(ctx")
        )
        _rejects(root, "native ownership methods changed")
        native.write_text(original.replace('/secret_policys"', '/other_policys"'))
        _rejects(root, "native ownership collection changed")
        native.write_text(original)
        generate(root)


if __name__ == "__main__":
    verify_contracts()
