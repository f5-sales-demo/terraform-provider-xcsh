#!/usr/bin/env python3
"""Publish exact tagged canonical documentation Markdown without rebuilding provider schemas."""

import argparse
import gzip
import hashlib
import io
import json
import os
import re
import shutil
import subprocess
import tarfile
import time
from pathlib import Path, PurePosixPath
from typing import Any

REPOSITORY = "f5-sales-demo/terraform-provider-xcsh"
VERSION = re.compile(r"v[0-9]+\.[0-9]+\.[0-9]+\Z")
CREATE_ATTEMPTS = 3


def run(*args: str) -> bytes:
    """Execute fixed CLI arguments without a shell and check their status."""
    result = subprocess.run(  # noqa: S603 - fixed CLI commands with separate arguments
        [shutil.which(args[0]) or args[0], *args[1:]], check=False, capture_output=True
    )
    if result.returncode:
        detail = result.stderr.decode("utf-8", errors="replace").strip()
        message = f"{args[0]} exited {result.returncode}: {detail or 'no stderr'}"
        raise RuntimeError(message)
    return result.stdout


def sha(data: bytes) -> str:
    """Return a plain SHA-256 digest."""
    return hashlib.sha256(data).hexdigest()


def encoded(value: Any) -> bytes:
    """Serialize canonical UTF-8 JSON."""
    return (
        json.dumps(value, sort_keys=True, separators=(",", ":"), ensure_ascii=False)
        + "\n"
    ).encode()


def metadata(relative: str, data: bytes) -> tuple[str, dict]:
    """Validate enriched bodies or supply path-based metadata."""
    text = data.decode("utf-8")
    body = text
    enriched = None
    if text.startswith("---\n"):
        header, separator, body = text[4:].partition("\n---\n")
        if not separator:
            message = f"unterminated frontmatter: {relative}"
            raise ValueError(message)
        for line in header.splitlines():
            if line.startswith("xcsh_docs: "):
                enriched = json.loads(line.removeprefix("xcsh_docs: "))
        # Generated body hashes exclude the blank line after frontmatter.
        body = body.removeprefix("\n")
    if enriched is None:
        title = next(
            (line[2:] for line in body.splitlines() if line.startswith("# ")), relative
        )
        enriched = {
            "schema_version": 1,
            "id": "xcsh-docs:path:" + relative,
            "canonical_id": "xcsh-docs:path:" + relative,
            "path": relative,
            "provider_type": "provider"
            if relative == "documentation/index.md"
            else "guides",
            "provider_name": "xcsh",
            "role": "overview",
            "schema_path": [],
            "summary": title,
            "aliases": [],
            "parent_id": None,
            "child_ids": [],
        }
    else:
        if enriched.get("path") != relative:
            message = f"metadata path mismatch: {relative}"
            raise ValueError(message)
        if enriched.get("body_sha256") != "sha256:" + sha(
            body.encode()
        ) or enriched.get("body_bytes") != len(body.encode()):
            message = f"metadata body mismatch: {relative}"
            raise ValueError(message)
    enriched = {
        **enriched,
        "canonical_id": enriched.get("canonical_id", enriched["id"]),
    }
    return body, enriched


def canonical_assets(
    tag: str, source_root: str = "documentation"
) -> tuple[dict[str, bytes], dict]:
    """Archive the complete canonical corpus, preserving exact source receipts."""
    source = run("git", "archive", "--format=tar", tag, source_root)
    archive = io.BytesIO()
    files = []
    with (
        tarfile.open(fileobj=io.BytesIO(source)) as tagged,
        tarfile.open(fileobj=archive, mode="w", format=tarfile.PAX_FORMAT) as target,
    ):
        for member in sorted(tagged, key=lambda item: item.name):
            if not member.isfile():
                continue
            path = PurePosixPath(member.name)
            if path.is_absolute() or ".." in path.parts or path.parts[0] != source_root:
                msg = "unsafe canonical archive member"
                raise ValueError(msg)
            stream = tagged.extractfile(member)
            if stream is None:
                msg = "missing canonical archive content"
                raise ValueError(msg)
            data = stream.read()
            info = tarfile.TarInfo(member.name)
            info.size, info.mode, info.mtime = len(data), 0o644, 0
            target.addfile(info, io.BytesIO(data))
            files.append(
                {"path": member.name, "size_bytes": len(data), "sha256": sha(data)}
            )
    generated = json.loads(
        run("git", "show", tag + ":documentation/generated-manifest.json")
    )
    manifest = {
        "schema_version": 1,
        "source_root": source_root,
        "provider_version": tag,
        "source_commit": run("git", "rev-parse", tag + "^{commit}").decode().strip(),
        "provider_schema_digest": generated["provider_schema_digest"],
        "spec_pin_digest": generated["spec_pin_digest"],
        "files": files,
    }
    compressed = bytearray(gzip.compress(archive.getvalue(), compresslevel=9, mtime=0))
    compressed[9] = 255
    archive_name = (
        "canonical-documentation.tar.gz"
        if source_root == "documentation"
        else "registry-documentation.tar.gz"
    )
    manifest_name = (
        "canonical-manifest.json"
        if source_root == "documentation"
        else "registry-manifest.json"
    )
    return {archive_name: bytes(compressed), manifest_name: encoded(manifest)}, manifest


def snapshot(tag: str, output: Path) -> dict:  # pylint: disable=too-many-locals
    """Package every Markdown file from one exact stable Git tag."""
    if not VERSION.fullmatch(tag):
        message = "only stable vN.N.N provider tags are allowed"
        raise ValueError(message)
    commit = run("git", "rev-parse", tag + "^{commit}").decode().strip()
    source = run("git", "archive", "--format=tar", tag, "documentation")
    documents: list[dict[str, Any]] = []
    archive = io.BytesIO()
    with (
        tarfile.open(fileobj=io.BytesIO(source)) as tagged,
        tarfile.open(fileobj=archive, mode="w", format=tarfile.PAX_FORMAT) as target,
    ):
        members = sorted(
            (member for member in tagged if member.name.endswith(".md")),
            key=lambda member: member.name,
        )
        for member in members:
            relative = PurePosixPath(member.name)
            if (
                not member.isfile()
                or relative.is_absolute()
                or ".." in relative.parts
                or relative.parts[0] != "documentation"
            ):
                message = f"unsafe tagged document: {member.name}"
                raise ValueError(message)
            extracted = tagged.extractfile(member)
            if extracted is None:
                message = "tagged Markdown member has no file content"
                raise ValueError(message)
            data = extracted.read()
            body, meta = metadata(member.name, data)
            info = tarfile.TarInfo(member.name)
            info.size, info.mode, info.mtime = len(data), 0o644, 0
            target.addfile(info, io.BytesIO(data))
            documents.append(
                {
                    "path": member.name,
                    "size_bytes": len(data),
                    "sha256": sha(data),
                    "body_sha256": sha(body.encode()),
                    "metadata": meta,
                }
            )
    ids = {d["metadata"]["id"] for d in documents} | {
        d["metadata"]["canonical_id"] for d in documents
    }
    for document in documents:
        meta = document["metadata"]
        for related in [meta.get("parent_id"), *meta.get("child_ids", [])]:
            if related and related not in ids:
                message = (
                    f"missing relationship target: {document['path']} -> {related}"
                )
                raise ValueError(message)
    versions = {d["metadata"].get("retrieval_version") for d in documents}
    if versions not in ({1}, {None}):
        message = "snapshot requires complete retrieval metadata version 1"
        raise ValueError(message)
    generated = json.loads(
        run("git", "show", tag + ":documentation/generated-manifest.json")
    )
    manifest = {
        **({"retrieval_metadata_version": 1} if versions == {1} else {}),
        "schema_version": 2,
        "source_root": "documentation",
        "source_repository": REPOSITORY,
        "provider_version": tag,
        "source_commit": commit,
        "document_count": len(documents),
        "documents": documents,
        "provider_schema_digest": generated["provider_schema_digest"],
        "spec_pin_digest": generated["spec_pin_digest"],
    }
    output.mkdir(parents=True, exist_ok=True)
    assets = {
        "terraform-docs.tar.gz": gzip.compress(
            archive.getvalue(), compresslevel=9, mtime=0
        ),
        "manifest.json": encoded(manifest),
    }
    compressed = bytearray(assets["terraform-docs.tar.gz"])
    compressed[9] = 255
    assets["terraform-docs.tar.gz"] = bytes(compressed)
    if "documentation/registry-projection-manifest.json" in generated["files"]:
        canonical, _ = canonical_assets(tag)
        assets.update(canonical)
        registry, _ = canonical_assets(tag, "docs")
        assets.update(registry)
        assets["registry-projection-manifest.json"] = run(
            "git", "show", tag + ":documentation/registry-projection-manifest.json"
        )
    assets["publication.json"] = encoded(
        {
            "schema_version": 2,
            **({"retrieval_metadata_version": 1} if versions == {1} else {}),
            "source_root": "documentation",
            "source_repository": REPOSITORY,
            "release_tag": "documentation-" + tag,
            "provider_version": tag,
            "source_commit": commit,
            "provider_schema_digest": manifest["provider_schema_digest"],
            "spec_pin_digest": manifest["spec_pin_digest"],
            "document_count": len(documents),
            "assets": {
                name: {"sha256": sha(data), "size_bytes": len(data)}
                for name, data in assets.items()
            },
        }
    )
    assets["SHA256SUMS"] = "".join(
        f"{sha(data)}  {name}\n" for name, data in sorted(assets.items())
    ).encode()
    for name, data in assets.items():
        (output / name).write_bytes(data)
    return manifest


def assert_no_partial_snapshot(tag: str) -> None:
    """Permit a create retry only when GitHub has no tag or draft for it."""
    tag_ref = subprocess.run(  # noqa: S603 - fixed endpoint from validated tag
        [
            shutil.which("gh") or "/usr/bin/gh",
            "api",
            f"repos/{REPOSITORY}/git/ref/tags/{tag}",
        ],
        check=False,
        capture_output=True,
    )
    if tag_ref.returncode == 0:
        message = "snapshot creation left a tag; refusing a duplicate"
        raise ValueError(message)
    if b"404" not in tag_ref.stderr:
        message = "cannot verify snapshot tag absence after failed create"
        raise RuntimeError(message)
    pages = json.loads(
        run(
            "gh",
            "api",
            "--paginate",
            "--slurp",
            f"repos/{REPOSITORY}/releases?per_page=100",
        )
    )
    if any(release.get("tag_name") == tag for page in pages for release in page):
        message = "snapshot creation left a draft; refusing a duplicate"
        raise ValueError(message)


def create_snapshot_draft(tag: str, arguments: list[str]) -> None:
    """Retry transient create failures only after proving no remote state exists."""
    for attempt in range(CREATE_ATTEMPTS):
        try:
            run("gh", "release", "create", tag, *arguments)
        except RuntimeError:
            assert_no_partial_snapshot(tag)
            if attempt == CREATE_ATTEMPTS - 1:
                raise
            time.sleep(attempt + 1)
        else:
            return


def publish(tag: str, output: Path) -> None:  # pylint: disable=too-many-locals
    """Publish assets atomically under repository immutable-release policy."""
    if not VERSION.fullmatch(tag):
        message = "only stable vN.N.N provider tags are allowed"
        raise ValueError(message)
    release = json.loads(run("gh", "api", f"repos/{REPOSITORY}/releases/tags/{tag}"))
    if (
        release["draft"]
        or release["prerelease"]
        or not release.get("immutable")
        or release["tag_name"] != tag
    ):
        message = "provider release must be published, stable, and immutable"
        raise ValueError(message)
    generated = json.loads(
        run("git", "show", tag + ":documentation/generated-manifest.json")
    )
    for kind in ("resources", "data-sources", "actions", "ephemeral-resources"):
        relative = f"documentation/{kind}/index.md"
        data = run("git", "show", tag + ":" + relative)
        evidence = generated["files"].get(relative)
        if (
            not evidence
            or evidence["sha256"] != "sha256:" + sha(data)
            or b"functions/index.md" in data
        ):
            message = (
                "provider release must contain manifest-owned corrected navigation"
            )
            raise ValueError(message)
    docs_tag = "documentation-" + tag
    existing = subprocess.run(  # noqa: S603 - fixed endpoint from validated tag
        [
            shutil.which("gh") or "/usr/bin/gh",
            "api",
            f"repos/{REPOSITORY}/releases/tags/{docs_tag}",
        ],
        check=False,
        capture_output=True,
    )
    if existing.returncode == 0:
        previous = json.loads(existing.stdout)
        if previous["draft"] or previous["prerelease"] or not previous.get("immutable"):
            message = "existing snapshot is not published and immutable"
            raise ValueError(message)
        snapshot(tag, output)
        names = {file.name for file in output.iterdir() if file.is_file()}
        published = {asset["name"]: asset for asset in previous["assets"]}
        if set(published) != names:
            message = "existing snapshot asset membership mismatch"
            raise ValueError(message)
        for name in sorted(names):
            asset = published[name]
            local = (output / name).read_bytes()
            if asset["size"] != len(local) or asset.get("digest") != "sha256:" + sha(
                local
            ):
                message = "existing immutable snapshot differs from exact tagged files"
                raise ValueError(message)
        print(f"Already published: {docs_tag}")
        return
    if b"404" not in existing.stderr:
        message = "snapshot release lookup failed"
        raise RuntimeError(message)
    policy_environment = dict(os.environ)
    if os.environ.get("POLICY_READ_TOKEN"):
        policy_environment["GH_TOKEN"] = os.environ["POLICY_READ_TOKEN"]
    policy = json.loads(
        subprocess.run(  # noqa: S603 - fixed administration read endpoint
            [
                shutil.which("gh") or "/usr/bin/gh",
                "api",
                f"repos/{REPOSITORY}/immutable-releases",
            ],
            env=policy_environment,
            check=True,
            capture_output=True,
        ).stdout
    )
    if not policy.get("enabled"):
        message = "immutable releases must be enabled"
        raise ValueError(message)
    latest_before = json.loads(run("gh", "api", f"repos/{REPOSITORY}/releases/latest"))[
        "tag_name"
    ]
    manifest = snapshot(tag, output)
    create_snapshot_draft(
        docs_tag,
        [
            "--repo",
            REPOSITORY,
            "--target",
            manifest["source_commit"],
            "--draft",
            "--latest=false",
            "--title",
            f"Terraform documentation {tag}",
            "--notes",
            f"Exact documentation/ Markdown from provider {tag} at {manifest['source_commit']}.",
            *(
                str(output / name)
                for name in sorted(
                    file.name for file in output.iterdir() if file.is_file()
                )
            ),
        ],
    )
    run(
        "gh",
        "release",
        "edit",
        docs_tag,
        "--repo",
        REPOSITORY,
        "--draft=false",
        "--latest=false",
    )

    published = json.loads(
        run("gh", "api", f"repos/{REPOSITORY}/releases/tags/{docs_tag}")
    )
    latest_after = json.loads(run("gh", "api", f"repos/{REPOSITORY}/releases/latest"))[
        "tag_name"
    ]
    if not published.get("immutable") or latest_after != latest_before:
        message = "snapshot publication failed immutability or provider latest check"
        raise ValueError(message)


def main() -> None:
    """Run explicit publication or reconcile the newest stable release."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--tag")
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--publish", action="store_true")
    args = parser.parse_args()
    tags = (
        [args.tag]
        if args.tag
        else [
            r["tag_name"]
            for page in json.loads(
                run(
                    "gh",
                    "api",
                    "--paginate",
                    "--slurp",
                    f"repos/{REPOSITORY}/releases?per_page=100",
                )
            )
            for r in page
            if not r["draft"]
            and not r["prerelease"]
            and VERSION.fullmatch(r["tag_name"])
        ]
    )
    if not args.tag:
        tags = [max(tags, key=lambda value: tuple(map(int, value[1:].split("."))))]
    for tag in sorted(tags, key=lambda value: tuple(map(int, value[1:].split(".")))):
        (publish if args.publish else snapshot)(tag, args.output / tag)


if __name__ == "__main__":
    main()
