# ruff: noqa: INP001
"""Deterministic, complete Registry projections of canonical Markdown sections."""

import hashlib
import json
import os
import re
from collections import defaultdict
from collections.abc import Callable
from pathlib import PurePosixPath
from typing import Any

TARGET = 450_000
LIMIT = 500_000
LINK = re.compile(r"\[([^\]\n]+)\]\(([^)\n]+)\)")
ANCHOR = re.compile(r'<a id="([^"]+)"')


def sha(text: str) -> str:
    """Digest exact UTF-8 section bytes."""
    return "sha256:" + hashlib.sha256(text.encode()).hexdigest()


def token(identifier: str) -> str:
    """Derive collision-safe explicit anchors from canonical identities."""
    value = int.from_bytes(hashlib.sha256(identifier.encode()).digest(), "big")
    digits = "".join(str((value >> shift) & 3) for shift in range(254, -1, -2))
    return "canonical-" + "-".join(
        digits[index : index + 16] for index in range(0, 128, 16)
    )


def pack(
    parts: list[str], overhead: str, target: int = TARGET, limit: int = LIMIT
) -> list[list[str]]:
    """Pack indivisible parts, counting exact UTF-8 wrapper bytes."""
    groups: list[list[str]] = []
    group: list[str] = []
    for part in parts:
        if len((overhead + part).encode()) > limit:
            msg = "indivisible section exceeds Registry byte limit"
            raise ValueError(msg)
        if group and len((overhead + "".join(group) + part).encode()) > target:
            groups.append(group)
            group = []
        group.append(part)
    if group:
        groups.append(group)
    return groups


def sections(body: str, budget: int) -> list[str]:
    """Split headings outside fences and oversized tables between complete rows."""
    blocks: list[str] = []
    current: list[str] = []
    fence = None
    for line in body.splitlines(keepends=True):
        marker = re.match(r"^\s{0,3}(`{3,}|~{3,})", line)
        if marker:
            run = marker.group(1)
            if fence is None:
                fence = run
            elif run[0] == fence[0] and len(run) >= len(fence):
                fence = None
        if fence is None and re.match(r"^#{1,3} ", line) and current:
            blocks.append("".join(current))
            current = []
        current.append(line)
    if current:
        blocks.append("".join(current))
    result = []
    for block in blocks:
        if len(block.encode()) <= budget:
            result.append(block)
            continue
        lines = block.splitlines(keepends=True)
        table = next(
            (
                i
                for i in range(len(lines) - 1)
                if lines[i].startswith("|")
                and re.match(r"^\|[ :|-]+\|\s*$", lines[i + 1])
            ),
            None,
        )
        if table is None:
            result.append(block)
            continue
        end = table + 2
        while end < len(lines) and lines[end].startswith("|"):
            end += 1
        header = "".join(lines[table : table + 2])
        prefix = "".join(lines[:table])
        chunk = prefix + header
        for row in lines[table + 2 : end]:
            if len((chunk + row).encode()) > budget and chunk != prefix + header:
                result.append(chunk)
                chunk = prefix + header
            chunk += row
        result.append(chunk)
        if end < len(lines):
            result.extend(sections("".join(lines[end:]), budget))
    return result


def canonical_url(page: dict[str, Any], site: str) -> str:
    """Return the original canonical section route."""
    return (
        site
        + "/"
        + page["path"].removeprefix("documentation/").removesuffix("index.md")
    )


def wrapper(title: str, category: str) -> str:
    """Build the byte-counted Registry document wrapper."""
    # JSON string syntax is valid YAML and preserves arbitrary source titles.

    return (
        "---\npage_title: "
        + json.dumps(title)
        + "\nsubcategory: "
        + json.dumps(category)
        + "\ndescription: "
        + json.dumps("Complete grouped canonical reference for " + title + ".")
        + "\n---\n\n# "
        + title
        + "\n\n"
    )


def rewrite_prose(body: str, callback: Callable) -> str:
    """Leave fenced and inline code bytes intact while rewriting prose links."""
    pieces = re.split(
        r"(?ms)(^\s{0,3}`{3,}[^\n]*\n.*?^\s{0,3}`{3,}\s*$|^\s{0,3}~{3,}[^\n]*\n.*?^\s{0,3}~{3,}\s*$|`[^`\n]*`)",
        body,
    )
    return "".join(
        piece if index % 2 else LINK.sub(callback, piece)
        for index, piece in enumerate(pieces)
    )


def collect_sections(
    pages: list[dict[str, Any]], site: str, version: str | None
) -> tuple[dict, list]:
    """Resolve coherent canonical sections and explicit oversized exceptions."""
    buckets, records = defaultdict(list), []
    for page in pages:
        role = page["role"]
        family = (
            "landing"
            if role == "fundamentals"
            else "reference"
            if role in ("reference", "properties")
            else "examples"
            if role in ("examples", "example", "negative-example")
            else "lifecycle"
        )
        bucket = (page["collection_id"], family)
        # Reserve space for navigation, wrappers and expanded destination links.
        for number, original in enumerate(sections(page["body"], TARGET - 16_384)):
            identifier = page["id"] + ":section:" + str(number)
            anchor = token(identifier)
            anchor_map = {
                old: token(page["id"] + "#" + old) for old in ANCHOR.findall(original)
            }

            def replace_anchor(
                match: re.Match, mapping: dict[str, str] = anchor_map
            ) -> str:
                return '<a id="' + mapping[match.group(1)] + '"'

            body = ANCHOR.sub(replace_anchor, original)
            canonical = canonical_url(page, site)
            if version:
                canonical = canonical.replace(
                    site + "/", site + "/versions/" + version + "/", 1
                )
            record = {
                "canonical_id": page["id"],
                "section_id": identifier,
                "source_sha256": sha(original),
                "source_bytes": len(original.encode()),
                "canonical_source_sha256": sha(page["body"]),
                "canonical_url": canonical,
                "anchor": anchor,
                "anchor_map": anchor_map,
                "mode": "embedded",
            }
            if len(body.encode()) > LIMIT - 16_384:
                if not version:
                    msg = "oversized exceptions require an immutable canonical version"
                    raise ValueError(msg)
                record.update(
                    mode="canonical-link",
                    reason="Indivisible coherent section exceeds the 500,000-byte document limit.",
                )
                source_anchor = next(iter(anchor_map), "")
                record["canonical_url"] += (
                    ("#" + source_anchor) if source_anchor else ""
                )
                body = (
                    "## Oversized canonical section\n\nThis indivisible section exceeds the Registry document limit. "
                    "[Complete versioned canonical section]("
                    + record["canonical_url"]
                    + ").\n\n"
                )
            buckets[bucket].append(
                {
                    "page": page,
                    "record": record,
                    "text": publication_body(
                        '<a id="' + anchor + '"></a>\n\n' + body + "\n\n",
                        page["title"].split(".")[-1]
                        + " / "
                        + token(page["id"])[-12:]
                        + " / "
                        + str(number + 1),
                    ),
                }
            )
            records.append(record)
    return buckets, records


def assign_groups(buckets: dict, categories: dict[str, str]) -> tuple[dict, dict]:
    """Assign deterministic grouped files and every fragment destination."""
    outputs: dict[str, Any] = {}
    destinations: dict[Any, tuple[str, str]] = {}
    for (_, family), items in sorted(buckets.items()):
        page = items[0]["page"]
        title = "xcsh_" + page["provider_name"] + " " + family
        header = wrapper(title, categories[page["collection_id"]])
        groups: list[list[dict]] = []
        group: list[dict] = []
        size = len(header.encode())
        for item in items:
            # Bound link expansion conservatively before destinations are known.
            item_size = len(item["text"].encode()) + 512 * len(
                LINK.findall(item["text"])
            )
            if group and size + item_size > TARGET:
                groups.append(group)
                group, size = [], len(header.encode())
            group.append(item)
            size += item_size
        if group:
            groups.append(group)
        for number, group in enumerate(groups, 1):
            path = (
                f"docs/{page['provider_type']}/{page['provider_name']}.md"
                if family == "landing"
                else f"docs/guides/{page['provider_type']}--{page['provider_name']}--{family}--group-{number:03}.md"
            )
            if family == "landing" and number > 1:
                msg = "landing page exceeds one Registry document"
                raise ValueError(msg)
            for item in group:
                record = item["record"]
                record["registry_path"] = path
                destinations.setdefault(item["page"]["id"], (path, record["anchor"]))
                for old, new in record["anchor_map"].items():
                    destinations[(item["page"]["id"], old)] = (path, new)
            outputs[path] = (header, group)
    return outputs, destinations


def publication_body(text: str, context: str) -> str:
    """Give embedded headings context and normalize only prose spacing."""
    parts = re.split(r"(?ms)(^```[^\n]*\n.*?^```\s*$|^~~~[^\n]*\n.*?^~~~\s*$)", text)
    for index in range(0, len(parts), 2):
        part = re.sub(
            r"(?m)^#{1,3} (.+)$",
            lambda match: "## " + match.group(1) + " — " + context,
            parts[index],
        )
        terms = {
            "key value": "key-value",
            "name space": "namespace",
            "id": "ID",
            "ipv4": "IPv4",
            "ipv6": "IPv6",
            "indexes": "indices",
            "cloudflare": "Cloudflare",
            "azure": "Azure",
            "url": "URL",
            "javascript": "JavaScript",
        }
        segments = re.split(r"(`[^`\n]*`|\[[^]\n]+\]\([^ )\n]+\)|<[^>]*>)", part)
        for offset in range(0, len(segments), 2):
            for before, after in terms.items():
                segments[offset] = re.sub(
                    r"\b" + re.escape(before) + r"\b", after, segments[offset]
                )
        part = "".join(segments).replace("[id](", "[ID](")
        parts[index] = re.sub(r"\n{3,}", "\n\n", part)
    return "\n\n".join(part.strip() for part in parts if part.strip()) + "\n\n"


def project(
    pages: list[dict[str, Any]],
    categories: dict[str, str],
    site: str,
    version: str | None = None,
) -> tuple[dict[str, str], dict[str, Any]]:
    """Emit grouped pages and a complete section/destination publication receipt."""
    by_url = {canonical_url(page, site): page["id"] for page in pages}
    buckets, records = collect_sections(pages, site, version)
    outputs, destinations = assign_groups(buckets, categories)
    for path, (header, items) in list(outputs.items()):

        def rewrite(match: re.Match, current_path: str = path) -> str:
            label, href = match.groups()
            url, _, anchor = href.partition("#")
            identifier = by_url.get(url)
            if identifier is None:
                return match.group(0)
            destination = (
                destinations.get((identifier, anchor))
                if anchor
                else destinations.get(identifier)
            )
            if destination is None:
                raise ValueError("unknown canonical fragment: " + href)
            target, fragment = destination
            relative = os.path.relpath(target, str(PurePosixPath(current_path).parent))
            return f"[{label}]({relative}#{fragment})"

        text = header + "".join(rewrite_prose(item["text"], rewrite) for item in items)
        text = text.rstrip() + "\n"
        if len(text.encode()) > LIMIT:
            raise ValueError("Registry document exceeds 500,000 bytes: " + path)
        outputs[path] = text
    for page in pages:
        page["registry_path"], page["registry_anchor"] = destinations[page["id"]]
    for record in records:
        record["anchor_map"] = list(record["anchor_map"].items())
    manifest = {
        "schema_version": 1,
        "target_bytes": TARGET,
        "limit_bytes": LIMIT,
        "sections": records,
        "files": {
            path: {"bytes": len(text.encode()), "sha256": sha(text)}
            for path, text in sorted(outputs.items())
        },
    }
    return outputs, manifest
