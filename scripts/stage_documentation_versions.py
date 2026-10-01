#!/usr/bin/env python3
"""Stage immutable canonical versions and isolated main preview for Pages."""

import argparse
import hashlib
import io
import json
import re
import shutil
import subprocess
import tarfile
from pathlib import Path

SITE = "https://f5-sales-demo.github.io/terraform-provider-xcsh"
VERSION = re.compile(r"v(\d+)\.(\d+)\.(\d+)$")


def run(*args: str) -> bytes:
    """Execute fixed Git/GitHub argv without shell interpolation."""
    return subprocess.run(  # noqa: S603 - fixed Git/GitHub argv
        [shutil.which(args[0]) or args[0], *args[1:]], check=True, capture_output=True
    ).stdout


def select(releases: list[dict]) -> list[str]:
    """Select stable immutable providers whose immutable docs snapshot exists."""
    published = {
        item["tag_name"]
        for item in releases
        if not item["draft"] and not item["prerelease"] and item.get("immutable")
    }
    return sorted(
        (
            tag
            for tag in published
            if VERSION.fullmatch(tag)
            and tuple(map(int, VERSION.fullmatch(tag).groups())) >= (12, 0, 6)
            and "docs-" + tag in published
        ),
        key=lambda tag: tuple(map(int, VERSION.fullmatch(tag).groups())),
    )


def transform(text: str, prefix: str) -> str:
    """Isolate rendered links while preserving original receipts separately."""
    return text.replace(SITE + "/", SITE + "/" + prefix + "/") if prefix else text


def digest(data: bytes) -> str:
    """Digest exact source or rendered bytes."""
    return "sha256:" + hashlib.sha256(data).hexdigest()


def transform_files(destination: Path, manifest: dict, prefix: str) -> dict:
    """Validate source receipts and transform version-local content."""
    updated = {}
    for path in sorted(manifest["files"]):
        if not path.startswith("documentation/"):
            continue
        relative = path.removeprefix("documentation/")
        file = destination / relative
        data = file.read_bytes()
        if (
            len(data) != manifest["files"][path]["bytes"]
            or digest(data) != manifest["files"][path]["sha256"]
        ):
            raise ValueError("tagged source differs from canonical receipt: " + path)
        if file.suffix in (".md", ".txt", ".json"):
            text = transform(data.decode(), prefix)
            if text.startswith("---\n") and "\nxcsh_docs: " in text:
                frontmatter, body = text.split("---\n\n", 1)
                lines = frontmatter.splitlines()
                for index, line in enumerate(lines):
                    if line.startswith("xcsh_docs: "):
                        meta = json.loads(line.removeprefix("xcsh_docs: "))
                        meta.update(
                            body_bytes=len(body.encode()),
                            body_sha256=digest(body.encode()),
                        )
                        updated[meta["id"]] = meta
                        lines[index] = "xcsh_docs: " + json.dumps(
                            meta, separators=(",", ":")
                        )
                text = "\n".join(lines) + "\n---\n\n" + body
            file.write_text(text)
    return updated


def stage(ref: str, destination: Path, prefix: str) -> None:
    """Extract exact commit content and build independent rendered receipts."""
    source = run("git", "archive", "--format=tar", ref, "documentation")
    destination.mkdir(parents=True, exist_ok=True)
    with tarfile.open(fileobj=io.BytesIO(source)) as archive:
        for member in archive:
            if not member.isfile():
                continue
            relative = Path(member.name).relative_to("documentation")
            if ".." in relative.parts:
                msg = "unsafe documentation member"
                raise ValueError(msg)
            extracted = archive.extractfile(member)
            if extracted is None:
                msg = "missing documentation member"
                raise ValueError(msg)
            target = destination / relative
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_bytes(extracted.read())
    manifest_path = destination / "generated-manifest.json"
    manifest = json.loads(manifest_path.read_text())
    source_receipts = destination / "_data/source-receipts.json"
    source_receipts.parent.mkdir(parents=True, exist_ok=True)
    source_receipts.write_text(json.dumps(manifest, sort_keys=True) + "\n")
    updated = transform_files(destination, manifest, prefix)
    for file in destination.rglob("*.json"):
        if file in (source_receipts, manifest_path):
            continue
        value = json.loads(file.read_text())
        if isinstance(value, dict) and "pages" in value:
            for page in value["pages"]:
                if page["id"] in updated:
                    page.update(
                        {
                            key: updated[page["id"]][key]
                            for key in ("body_bytes", "body_sha256")
                        }
                    )
            file.write_text(json.dumps(value, separators=(",", ":")) + "\n")
    manifest["source_commit"] = (
        run("git", "rev-parse", ref + "^{commit}").decode().strip()
    )
    manifest["publication_prefix"] = prefix
    manifest["files"] = {
        "documentation/" + str(file.relative_to(destination)): {
            "bytes": len(file.read_bytes()),
            "sha256": digest(file.read_bytes()),
        }
        for file in sorted(destination.rglob("*"))
        if file.is_file() and file != manifest_path
    }
    manifest_path.write_text(json.dumps(manifest, sort_keys=True) + "\n")


def main() -> None:
    """Stage latest successfully published stable version plus immutable history."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--output", type=Path, required=True)
    args = parser.parse_args()
    releases = [
        item
        for page in json.loads(
            run(
                "gh",
                "api",
                "--paginate",
                "--slurp",
                "repos/f5-sales-demo/terraform-provider-xcsh/releases?per_page=100",
            )
        )
        for item in page
    ]
    versions = select(releases)
    if not versions:
        msg = "no successfully published stable documentation version"
        raise ValueError(msg)
    run("git", "fetch", "--tags", "origin")
    stage(versions[-1], args.output, "")
    for version in versions:
        stage(version, args.output / "versions" / version, "versions/" + version)
    stage("HEAD", args.output / "preview/main", "preview/main")
    (args.output / "documentation-versions.json").write_text(
        json.dumps(
            {
                "schema_version": 1,
                "latest": versions[-1],
                "versions": versions,
                "preview": "preview/main",
            }
        )
        + "\n"
    )


if __name__ == "__main__":
    main()
