# ruff: noqa: INP001, EM101, TRY003
# pylint: disable=invalid-name
"""Audit visible Markdown prose in canonical and Registry documentation.

The scanner does not infer prose quality from heading adjacency alone. A section
is empty only when its entire heading subtree lacks visible content.
"""

import argparse
import hashlib
import json
import re
import shutil
import subprocess
import sys
from collections import Counter
from pathlib import Path

SITE = "https://f5-sales-demo.github.io/terraform-provider-xcsh"
REGISTRY = "https://registry.terraform.io/providers/f5-sales-demo/xcsh/latest/docs"
HEADING = re.compile(r"^(#{1,6})\s+(.+?)\s*#*\s*$")
FENCE = re.compile(r"^\s{0,3}(`{3,}|~{3,})")
ANCHOR = re.compile(r'^\s*<a\s+id="[^"]+"\s*></a>\s*$')
NUMERIC_ARTIFACT = re.compile(r"\b[0-3]{12}\b|\bcanonical-[0-3]{16}(?:-[0-3]+)+\b")
GRAMMAR = re.compile(r"\b[aA] (?:HTTP|API|AWS|IKE|SSH|SSL)\b")
EDITORIAL = re.compile(r"\b(?:[aA]) (?:Enhanced|Origin|Advertise|Authentication)\b")
ENTRY_PAGE = re.compile(
    r"^documentation/(?:resources|data-sources|actions|ephemeral-resources)/[^/]+/index\.md$"
)
SUBSTANTIVE_LENGTH = 80
GIT_HEADER_FIELDS = 3


def route(relative: str) -> str:
    """Map a generated source path to its published route."""
    if relative.startswith("documentation/"):
        return (
            SITE
            + "/"
            + relative.removeprefix("documentation/").removesuffix("index.md")
        )
    path = Path(relative)
    if path.parts[1] == "guides":
        return REGISTRY + "/guides/" + path.stem
    return REGISTRY + "/" + "/".join(path.parts[1:]).removesuffix(".md")


def scan(path: Path, relative: str) -> list[dict]:
    """Classify visible content in one Markdown file."""
    return scan_text(path.read_text(encoding="utf-8"), relative)


def scan_text(source: str, relative: str) -> list[dict]:
    """Classify exact Markdown bytes supplied from a file or immutable Git tree."""
    # pylint: disable=too-many-locals,too-many-branches,too-many-statements
    lines = source.splitlines()
    start = 0
    if lines and lines[0] == "---":
        start = next(
            (i + 1 for i in range(1, len(lines)) if lines[i] == "---"), len(lines)
        )
    headings: list[tuple[int, int, str]] = []
    content: list[bool] = [False] * len(lines)
    findings: list[dict] = []
    seen_headings: dict[str, int] = {}
    seen_paragraphs: dict[tuple[int, str], int] = {}
    paragraph: list[str] = []
    paragraph_line = 0
    section_line = 0
    fence = None
    entry_intro = False
    entry_reported = False

    def add(category: str, severity: str, line: int, detail: str) -> None:
        findings.append(
            {
                "category": category,
                "severity": severity,
                "source": relative,
                "published_location": route(relative),
                "line": line,
                "detail": detail,
            }
        )

    def flush() -> None:
        nonlocal paragraph
        value = " ".join(" ".join(paragraph).split())
        if len(value) >= SUBSTANTIVE_LENGTH and not value.startswith(("|", "- [")):
            key = (section_line, value.casefold())
            if key in seen_paragraphs:
                add("repeated_description", "error", paragraph_line, value[:180])
            else:
                seen_paragraphs[key] = paragraph_line
        paragraph = []

    for index in range(start, len(lines)):
        line = lines[index]
        line_number = index + 1
        marker = FENCE.match(line)
        if marker:
            flush()
            run = marker.group(1)
            if fence is None:
                fence = run
            elif run[0] == fence[0] and len(run) >= len(fence):
                fence = None
            content[index] = True
            continue
        if fence:
            content[index] = True
            continue
        match = HEADING.match(line)
        if match:
            flush()
            level, label = len(match.group(1)), match.group(2).strip()
            if ENTRY_PAGE.fullmatch(relative):
                entry_intro = level == 1
            headings.append((index, level, label))
            section_line = line_number
            if NUMERIC_ARTIFACT.search(label):
                add("numeric_heading_artifact", "error", line_number, label)
            key = re.sub(r"\s+", " ", label).casefold()
            if key in seen_headings:
                add("duplicate_visible_heading", "error", line_number, label)
            else:
                seen_headings[key] = line_number
            continue
        if not line.strip() or line.lstrip().startswith("<!--") or ANCHOR.match(line):
            flush()
            continue
        content[index] = True
        if GRAMMAR.search(line):
            add("known_grammar_defect", "error", line_number, line.strip()[:180])
        elif EDITORIAL.search(line):
            add("editorial_candidate", "review", line_number, line.strip()[:180])
        if (
            entry_intro
            and not entry_reported
            and ("configuration." in line.casefold() or "metadata.namespace" in line)
        ):
            add("editorial_candidate", "review", line_number, line.strip()[:180])
            entry_reported = True
        if not paragraph:
            paragraph_line = line_number
        paragraph.append(line)
    flush()
    for position, (index, level, label) in enumerate(headings):
        end = len(lines)
        for later, next_level, _ in headings[position + 1 :]:
            if next_level <= level:
                end = later
                break
        if not any(content[index + 1 : end]):
            add("empty_section", "error", index + 1, label)
    return findings


def summarize(paths: list[str], findings: list[dict]) -> dict:
    """Count both source trees and all classified findings."""
    return {
        "schema_version": 1,
        "files": {
            "canonical": sum(path.startswith("documentation/") for path in paths),
            "registry": sum(path.startswith("docs/") for path in paths),
        },
        "findings": dict(
            sorted(Counter(item["category"] for item in findings).items())
        ),
        "errors": sum(item["severity"] == "error" for item in findings),
        "review_candidates": sum(item["severity"] == "review" for item in findings),
    }


def audit(root: Path) -> tuple[dict, list[dict]]:
    """Scan both documentation projections in the working tree."""
    files = sorted((root / "documentation").rglob("*.md"))
    files = [file for file in files if "_data" not in file.relative_to(root).parts]
    files += sorted(
        file
        for file in (root / "docs").rglob("*.md")
        if "specifications" not in file.relative_to(root / "docs").parts
    )
    paths = [file.relative_to(root).as_posix() for file in files]
    findings = [
        item
        for file, relative in zip(files, paths, strict=True)
        for item in scan(file, relative)
    ]
    return summarize(paths, findings), findings


def audit_git(root: Path, ref: str) -> tuple[dict, list[dict]]:
    """Read a published Git tree through one Git object stream without extraction."""
    git = shutil.which("git")
    if git is None:
        raise RuntimeError("git is required for --git-ref")
    names = subprocess.check_output(  # noqa: S603
        [git, "ls-tree", "-r", "--name-only", ref, "--", "docs", "documentation"],
        cwd=root,
        text=True,
    ).splitlines()
    paths = [
        name
        for name in names
        if name.endswith(".md")
        and "_data" not in Path(name).parts
        and "specifications" not in Path(name).parts
    ]
    findings = []
    with subprocess.Popen(  # noqa: S603
        [git, "cat-file", "--batch"],
        cwd=root,
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
    ) as process:
        if process.stdin is None or process.stdout is None:
            raise RuntimeError("Git batch pipes are unavailable")
        for relative in paths:
            process.stdin.write(f"{ref}:{relative}\n".encode())
            process.stdin.flush()
            header = process.stdout.readline().decode().split()
            if len(header) != GIT_HEADER_FIELDS or header[1] != "blob":
                raise RuntimeError("cannot read Git Markdown blob: " + relative)
            size = int(header[2])
            source = process.stdout.read(size).decode("utf-8")
            if process.stdout.read(1) != b"\n":
                raise RuntimeError("invalid Git batch delimiter")
            findings.extend(scan_text(source, relative))
        process.stdin.close()
        if process.wait() != 0:
            raise RuntimeError("Git batch reader failed")
    return summarize(paths, findings), findings


def redact_detail(finding: dict) -> dict:
    """Retain finding identity without republishing upstream prose."""
    record = dict(finding)
    record["detail_sha256"] = (
        "sha256:" + hashlib.sha256(record.pop("detail").encode()).hexdigest()
    )
    return record


def main() -> int:
    """Run the audit CLI and optionally fail on objective findings."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--root", type=Path, default=Path(__file__).resolve().parents[1]
    )
    parser.add_argument("--git-ref", help="audit an immutable Git tree")
    parser.add_argument("--ledger", type=Path)
    parser.add_argument("--redact-detail", action="store_true")
    parser.add_argument("--summary", type=Path)
    parser.add_argument("--strict", action="store_true")
    args = parser.parse_args()
    summary, findings = (
        audit_git(args.root, args.git_ref) if args.git_ref else audit(args.root)
    )
    if args.ledger:
        args.ledger.parent.mkdir(parents=True, exist_ok=True)
        with args.ledger.open("w", encoding="utf-8") as output:
            for finding in findings:
                record = redact_detail(finding) if args.redact_detail else finding
                output.write(
                    json.dumps(record, sort_keys=True, ensure_ascii=False) + "\n"
                )
    if args.summary:
        args.summary.parent.mkdir(parents=True, exist_ok=True)
        args.summary.write_text(json.dumps(summary, indent=2, sort_keys=True) + "\n")
    sys.stdout.write(json.dumps(summary, sort_keys=True) + "\n")
    return int(args.strict and summary["errors"] > 0)


if __name__ == "__main__":
    raise SystemExit(main())
