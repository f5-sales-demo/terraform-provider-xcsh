#!/usr/bin/env python3
# CLI diagnostics and schema dictionaries intentionally use dynamic types.
# ruff: noqa: ANN001, ANN201, ANN202, ANN204, ANN205, D101, D102, D103, D107, EM101, EM102, TRY003, T201
# pylint: disable=invalid-name,too-many-instance-attributes,too-many-arguments,too-many-locals,too-many-branches,too-many-lines,too-many-statements
"""Complete schema collections and destination-specific Registry projections.

The installed schema is the reference authority. Enrichment and AST-extracted
validators supplement it; transformed Markdown is never used as schema input.
"""

import argparse
import hashlib
import html
import json
import os
import re
import shutil
import subprocess
import sys
import textwrap
from pathlib import Path, PurePosixPath
from urllib.parse import quote

sys.path.insert(0, str(Path(__file__).resolve().parent))
from import_contract import resolve_import_contract
from registry_projection import project as grouped_project
from retrieval_metadata import (
    RetrievalRules,
    constraint_relationships,
    summary as retrieval_summary,
)

PROVIDER = "registry.terraform.io/f5-sales-demo/xcsh"
SITE = "https://f5-sales-demo.github.io/terraform-provider-xcsh"
REGISTRY_FILENAME_BUDGET = 240
REGISTRY_LIMIT = 500_000  # HashiCorp documents 500KB, including frontmatter.
PROJECTION_RECEIPT_BUDGET = 500_000
TYPES = {
    "resources": ("resource_schemas", "resources", "resource.tf", "resource"),
    "data-sources": (
        "data_source_schemas",
        "data_sources",
        "data-source.tf",
        "data_source",
    ),
    "actions": ("action_schemas", "actions", "action.tf", "action"),
    "ephemeral-resources": (
        "ephemeral_resource_schemas",
        "ephemeral_resources",
        "ephemeral.tf",
        "ephemeral_resource",
    ),
}
LINK = re.compile(r"\[([^\]\n]+)\]\(([^)\n]+)\)")


def prose_links(text):
    """Parse generated navigation only, leaving code examples and regexes intact."""
    without_fences = re.sub(r"(?ms)^```[^\n]*\n.*?^```\s*$", "", text)
    without_code = re.sub(r"`[^`\n]*`", "", without_fences)
    return LINK.findall(re.sub(r"\\([\[\]])", "", without_code))


def description_markdown(value):
    """Publish full descriptions as plain prose without interpreting API markup."""
    escaped = html.escape(value, quote=False)
    escaped = re.sub(r"([\\`*_{}\[\]<>#])", r"\\\1", escaped)
    escaped = re.sub(
        r"(?i)(?<![A-Za-z0-9])www\.([A-Za-z0-9.-]+)",
        lambda match: "www&#46;" + match.group(1),
        escaped,
    )
    escaped = re.sub(
        r"https?://", lambda match: match.group(0).replace(":", "&#58;"), escaped
    )
    escaped = re.sub(
        r"(?m)^([ ]*)(?=[+.-] |[0-9]+[.)] )",
        lambda match: match.group(1) + "&#8203;",
        escaped,
    )
    escaped = re.sub(
        r"(?i)\bnotin\b",
        lambda match: match.group(0)[0] + "&#111;" + match.group(0)[2:],
        escaped,
    )
    paragraphs = re.split(r"\n\s*\n", escaped)
    result = "\n\n".join(
        textwrap.fill(
            " ".join(part.split()),
            width=100,
            break_long_words=False,
            break_on_hyphens=False,
        )
        for part in paragraphs
    )
    return re.sub(r"(?m)^(?=[+.-] |[0-9]+[.)] )", "&#8203;", result)


def normalize_body(value):
    """Normalize whitespace while preserving complete fenced examples."""
    value = value.expandtabs(2)
    return re.sub(r"\n{3,}", "\n\n", value).rstrip() + "\n"


def digest(data):
    return "sha256:" + hashlib.sha256(data).hexdigest()


def read_json(path):
    return json.loads(Path(path).read_text(encoding="utf-8"))


def stable_id(kind, name, role, path=()):
    # Exact identifiers, with reversible segment encoding, independent of content.
    return ":".join(["xcsh-docs", kind, name, role, *[quote(s, safe="") for s in path]])


def json_text(value):
    return json.dumps(value, ensure_ascii=False, sort_keys=True)


def object_attributes(typ):
    if not isinstance(typ, list):
        return None
    if typ[0] == "object":
        return {key: {"type": value} for key, value in typ[1].items()}
    if typ[0] in ("list", "set", "map"):
        return object_attributes(typ[1])
    if typ[0] == "tuple":
        # Tuples are scalar references unless all entries share an object shape.
        shapes = [object_attributes(t) for t in typ[1]]
        if shapes and all(s == shapes[0] for s in shapes):
            return shapes[0]
    return None


def children(block):
    fields = {}
    for name, attribute in block.get("attributes", {}).items():
        nested = attribute.get("nested_type")
        shape = {"attributes": nested.get("attributes", {})} if nested else None
        if shape is None:
            attributes = object_attributes(attribute.get("type"))
            if attributes is not None:
                shape = {"attributes": attributes}
        if shape is not None and "type" in attribute:
            # Child flags of a plain object inherit the enclosing value flags.
            for child in shape["attributes"].values():
                for flag in ("required", "optional", "computed", "sensitive"):
                    if attribute.get(flag):
                        child.setdefault(flag, True)
        fields[name] = (attribute, shape, "attribute")
    for name, block_type in block.get("block_types", {}).items():
        if name in fields:
            raise ValueError(f"competing attribute and block: {name}")
        fields[name] = (block_type, block_type["block"], "block")
    return dict(sorted(fields.items()))


def schema_root(schema):
    # Actions have a config_schema rather than a resource block.
    return schema.get("block", schema.get("config_schema", {}))


class Specs:
    def __init__(self, root):
        self.root = root
        pin_file = root / "tools/spec-release.json"
        self.pin_digest = digest(pin_file.read_bytes())
        self.pin = read_json(pin_file)
        spec_dir = root / "docs/specifications/api"
        index_path = spec_dir / "index.json"
        if digest(index_path.read_bytes()) != self.pin["assets"]["index.json"]:
            raise ValueError("spec index differs from receipt pin")
        self.index = read_json(index_path)
        self.resources = {}
        # Consume the exact receipt-pinned aggregate; domain copies can rename
        # shared schemas to avoid collisions and are not a byte equality input.
        aggregate = read_json(spec_dir / "openapi.json")
        if (
            digest((spec_dir / "openapi.json").read_bytes())
            != self.pin["assets"]["openapi.json"]
        ):
            raise ValueError("aggregate specs differ from receipt pin")
        self.schemas = aggregate.get("components", {}).get("schemas", {})
        for entry in self.index["specifications"]:
            for resource in entry.get("x-f5xc-primary-resources", []):
                self.resources[resource["name"]] = (entry, resource)

    @staticmethod
    def resolve(value, schemas, seen=()):
        if not isinstance(value, dict):
            return {}
        result = {}
        ref = value.get("$ref")
        if ref:
            key = ref.rsplit("/", 1)[-1]
            if key not in seen:
                result.update(
                    Specs.resolve(schemas.get(key, {}), schemas, (*seen, key))
                )
        for member in value.get("allOf", []):
            result.update(Specs.resolve(member, schemas, seen))
        result.update({k: v for k, v in value.items() if k not in ("$ref", "allOf")})
        return result

    def resource(self, name):
        entry, metadata = self.resources.get(name, ({}, {}))
        metadata = dict(metadata)
        if not metadata.get("category") or metadata["category"] == "Other":
            metadata["category"] = entry.get("x-f5xc-category", "")
            metadata["category_source"] = "receipt-pinned-domain"
        else:
            metadata["category_source"] = "receipt-pinned-resource"
        schemas = self.schemas
        candidates = []
        for key, value in schemas.items():
            identity = value.get("x-ves-proto-message", "")
            if (
                identity.endswith(f".{name}.CreateRequest")
                or key == f"{name}CreateRequest"
            ):
                candidates.append((key, self.resolve(value, schemas)))
        roots = {}
        identity = []
        for key, candidate in candidates:
            identity.append(key)
            for envelope in ("metadata", "spec"):
                value = self.resolve(
                    candidate.get("properties", {}).get(envelope, {}), schemas
                )
                roots.update(value.get("properties", {}))
        return metadata, roots, schemas, identity

    def immutable_oneof_groups(self, name):
        groups = {}
        for key, value in self.schemas.items():
            if key.endswith(f"{name}CreateSpecType"):
                groups.update(
                    self.resolve(value, self.schemas).get(
                        "x-f5xc-immutable-oneof-groups", {}
                    )
                )
        return groups

    def field(self, roots, schemas, path):
        value = {"properties": roots}
        for segment in path:
            value = self.resolve(value, schemas)
            if value.get("type") == "array":
                value = self.resolve(value.get("items", {}), schemas)
            value = value.get("properties", {}).get(segment, {})
        return self.resolve(value, schemas)


class Collection:
    def __init__(self, kind, name, schema, specs, constraints, root, schema_digest):
        self.kind, self.name, self.schema = kind, name, schema
        self.root = root
        self.base = f"documentation/{kind}/{name}"
        self.pages = {}
        self.coverage = {}
        self.metadata, self.spec_roots, self.spec_schemas, identity = specs.resource(
            name
        )
        self.specs = specs
        self.constraints = constraints
        self.provenance = {
            "upstream_identity": {
                "release_tag": specs.pin["release_tag"],
                "target_commit": specs.pin["target_commit"],
                "schema_components": identity,
            },
            "spec_pin_digest": specs.pin_digest,
            "provider_schema_digest": schema_digest,
        }
        self.description = schema_root(schema).get(
            "description", ""
        ) or self.metadata.get(
            "description", f"Terraform {kind} configuration for xcsh_{name}."
        )
        self.category = self.metadata.get("category", "")
        self.retrieval_rules = RetrievalRules.default()
        self.add("fundamentals", (), "index.md", f"xcsh_{name}")
        self.add(
            "reference", (), "properties/index.md", "Property reference", "fundamentals"
        )
        self.walk(schema_root(schema), ())
        self.add_examples()
        self.add_lifecycle()
        self.enrich_retrieval()

    def enrich_retrieval(self):
        for page in self.pages.values():
            schema_path = page["schema_path"]
            page.update(
                self.retrieval_rules.classify(
                    self.kind, self.name, schema_path, page["role"], self.category
                )
            )
            description = self.description if page["role"] == "fundamentals" else ""
            if page.get("section"):
                field, _ = page["section"]
                description = field.get(
                    "description", field.get("block", {}).get("description", "")
                )
                upstream = self.specs.field(
                    self.spec_roots, self.spec_schemas, schema_path
                ).get("description", "")
                if upstream:
                    description = upstream
            page["summary"] = retrieval_summary(description, page["summary"])
            page["aliases"] = self.retrieval_rules.aliases(
                ".".join(schema_path) or self.name, description
            )
            page["relationships"] = constraint_relationships(
                schema_path,
                self.constraints.get(".".join(schema_path), {}),
                self.coverage,
            )
            page["classification"]["upstream_category_source"] = self.metadata.get(
                "category_source", "unresolved"
            )
            page["sections"] = []
            for name, field, _subsection, syntax in page.get("fields", []):
                exact = [*schema_path, name]
                target = self.coverage[".".join(exact)]
                description = field.get(
                    "description", field.get("block", {}).get("description", "")
                )
                upstream = self.specs.field(
                    self.spec_roots, self.spec_schemas, exact
                ).get("description", "")
                prose = retrieval_summary(
                    upstream or description, name.replace("_", " ")
                )
                page["sections"].append(
                    {
                        "schema_path": exact,
                        "document_id": target["document_id"],
                        "anchor": target["anchor"],
                        "description": prose,
                        "aliases": self.retrieval_rules.aliases(
                            name, upstream or description
                        ),
                        "flags": [
                            flag
                            for flag in (
                                "required",
                                "optional",
                                "computed",
                                "sensitive",
                                "deprecated",
                                "write_only",
                            )
                            if field.get(flag)
                        ],
                        "nesting": field.get(
                            "nesting_mode",
                            field.get("nested_type", {}).get("nesting_mode"),
                        ),
                        "min_items": field.get("min_items"),
                        "max_items": field.get("max_items"),
                        "syntax": syntax,
                        "type": field.get("type", "object")[0]
                        if isinstance(field.get("type"), list)
                        else field.get("type", "object"),
                        "relationships": constraint_relationships(
                            exact,
                            self.constraints.get(".".join(exact), {}),
                            self.coverage,
                        ),
                    }
                )

        self.enrich_immutable_choices()

    def enrich_immutable_choices(self):
        """Expose receipt-pinned type choices without claiming schema prerequisites."""
        if self.kind != "resources":
            return
        groups = getattr(self.specs, "immutable_oneof_groups", lambda _name: {})(
            self.name
        )
        for group, members in sorted(groups.items()):
            destinations = [self.coverage.get(member) for member in members]
            if any(target is None for target in destinations):
                raise ValueError("missing immutable choice destination: " + group)
            for page in self.pages.values():
                if page["role"] != "reference" and page["schema_path"] not in [
                    [member] for member in members
                ]:
                    continue
                for target in destinations:
                    if target["document_id"] == page["id"]:
                        continue
                    page["relationships"].append(
                        {
                            "type": "choice",
                            "target_id": target["document_id"],
                            "anchor": target["anchor"],
                            "enforcement": "provider-choice",
                            "source": "receipt-pinned-immutable-oneof",
                            "group": group,
                        }
                    )

    def add(self, role, path, filename, title, parent_role=None, parent_path=()):
        identifier = stable_id(self.kind, self.name, role, path)
        parent = (
            stable_id(self.kind, self.name, parent_role, parent_path)
            if parent_role
            else None
        )
        page = {
            "id": identifier,
            "collection_id": stable_id(self.kind, self.name, "collection"),
            "provider_type": self.kind,
            "provider_name": self.name,
            "role": role,
            "schema_path": list(path),
            "parent_id": parent,
            "child_ids": [],
            "summary": f"{title} for xcsh_{self.name}.",
            "aliases": [],
            "completeness": "complete",
            "path": f"{self.base}/{filename}",
            "title": title,
            **self.provenance,
        }
        if identifier in self.pages:
            raise ValueError(f"duplicate document ID: {identifier}")
        self.pages[identifier] = page
        if parent:
            self.pages[parent]["child_ids"].append(identifier)
        return page

    def walk(self, block, path, inherited=None):
        page = self.pages[
            stable_id(
                self.kind, self.name, "reference" if not path else "properties", path
            )
        ]
        page["fields"] = []
        for name, (field, shape, syntax) in children(block).items():
            exact = (*path, name)
            inherited_flags = dict(inherited or {})
            effective = dict(field)
            for flag in ("sensitive", "computed"):
                if inherited_flags.get(flag):
                    effective[flag] = True
            if shape is not None:
                leaf = self.add(
                    "properties",
                    exact,
                    f"properties/{'/'.join(exact)}/index.md",
                    ".".join(exact),
                    "reference" if not path else "properties",
                    path,
                )
                leaf["section"] = (effective, syntax)
                self.coverage[".".join(exact)] = {
                    "document_id": leaf["id"],
                    "anchor": "section",
                    "schema_path": list(exact),
                }
                self.walk(shape, exact, effective)
            else:
                self.coverage[".".join(exact)] = {
                    "document_id": page["id"],
                    "anchor": "schema-" + "--".join(exact),
                    "schema_path": list(exact),
                }
            page["fields"].append((name, effective, shape is not None, syntax))

    def add_examples(self):
        files = sorted(
            (self.root / "examples" / self.kind / f"xcsh_{self.name}").glob("*.tf")
        )
        canonical = TYPES[self.kind][2]
        if not any(p.name == canonical for p in files):
            raise ValueError(
                f"missing validated minimal example: {self.kind}/{self.name}/{canonical}"
            )
        index = self.add(
            "examples", (), "examples/index.md", "Examples", "fundamentals"
        )
        index["examples"] = []
        for file in files:
            negative = "conflict" in file.stem
            role = "negative-example" if negative else "example"
            page = self.add(
                role,
                (file.stem,),
                f"examples/{file.stem}/index.md",
                file.stem.replace("-", " ").capitalize(),
                "examples",
            )
            page["config"] = file.read_text(encoding="utf-8")
            page["evidence"] = {
                "source_path": file.relative_to(self.root).as_posix(),
                "sha256": digest(file.read_bytes()),
                "validation": "terraform validate",
                "outcome": "expected conflict" if negative else "valid configuration",
                "attribution": "Acceptance-test-derived fixture; no new live API execution is claimed."
                if file.name != canonical
                else "Schema-derived minimal configuration validated with the checked-out provider.",
            }
            index["examples"].append(page["id"])
            if file.name == canonical:
                self.minimal = page["id"]

    def add_lifecycle(self):
        if self.kind == "resources":
            source = self.root / "internal/provider" / f"{self.name}_resource.go"
            text = source.read_text() if source.exists() else ""
            contract = resolve_import_contract(text, self.name)
            if contract:
                page = self.add(
                    "import", (), "lifecycle/import/index.md", "Import", "fundamentals"
                )
                page["import"] = contract["command"]
                page["import_guidance"] = contract["guidance"]
        if "timeouts" in schema_root(self.schema).get("block_types", {}):
            self.add(
                "timeouts",
                (),
                "lifecycle/timeouts/index.md",
                "Timeouts",
                "fundamentals",
            )
        if self.kind == "resources" and getattr(
            self.specs, "immutable_oneof_groups", lambda _name: {}
        )(self.name):
            self.add("lifecycle", (), "lifecycle/index.md", "Lifecycle", "fundamentals")
        if self.kind in ("actions", "ephemeral-resources"):
            self.add("lifecycle", (), "lifecycle/index.md", "Lifecycle", "fundamentals")

    def link(self, identifier, label=None, anchor=None):
        target = self.pages[identifier]
        fragment = "#" + anchor if anchor else ""
        return f"[{label or target['title']}]({SITE}/{target['path'].removeprefix('documentation/').removesuffix('index.md')}{fragment})"

    def field_text(self, path, field, syntax, heading=True):
        name = path[-1]
        lines = (
            [f'<a id="schema-{"--".join(path)}"></a>\n\n### {name} property\n']
            if heading
            else []
        )
        block = field.get("block", {})
        flags = [
            flag.capitalize()
            for flag in (
                "required",
                "optional",
                "computed",
                "sensitive",
                "deprecated",
                "write_only",
            )
            if field.get(flag)
        ]
        if syntax == "block":
            flags.append(f"{field.get('nesting_mode', 'single')} nested block")
            if field.get("min_items"):
                flags.append("Required")
            elif not field.get("computed"):
                flags.append("Optional")
        typ = field.get(
            "type", field.get("nested_type", {}).get("nesting_mode", "object")
        )
        lines.append(f"Type: `{json_text(typ)}`. {', '.join(flags)}.\n")
        desc = field.get("description", block.get("description", ""))
        if desc:
            lines.append(description_markdown(desc) + "\n")
        spec = self.specs.field(self.spec_roots, self.spec_schemas, path)
        long_description = spec.get("description", "")
        if long_description and long_description != desc:
            lines.append(
                "Upstream description:\n\n"
                + description_markdown(long_description)
                + "\n"
            )
        lines.extend(
            f"{key}: `{field[key]}`.\n"
            for key in ("min_items", "max_items")
            if key in field
        )
        source = self.constraints.get(".".join(path), {})
        if source:
            lines.append(
                "Provider validators and defaults (from schema source):\n\n```go\n"
                + "\n".join(f"{k}: {v}" for k, v in sorted(source.items()))
                .replace(", validators.", ",\n  validators.")
                .replace(', "', ',\n    "')
                .replace(", validators.", ",\n  validators.")
                .replace(', "', ',\n    "')
                + "\n```\n"
            )
        enriched = {
            key: value
            for key, value in spec.items()
            if key
            in (
                "default",
                "enum",
                "minimum",
                "maximum",
                "exclusiveMinimum",
                "exclusiveMaximum",
                "minLength",
                "maxLength",
                "pattern",
                "minItems",
                "maxItems",
                "uniqueItems",
                "oneOf",
            )
            or key.startswith(
                (
                    "x-f5xc-constraints",
                    "x-ves-validation",
                    "x-validation",
                    "x-f5xc-default",
                    "x-f5xc-server-default",
                    "x-f5xc-oneof",
                    "x-ves-oneof",
                    "x-f5xc-required",
                    "x-f5xc-sensitive",
                )
            )
        }
        if enriched:
            lines.append(
                "Receipt-pinned upstream constraints:\n\n```json\n"
                + json.dumps(enriched, indent=2, sort_keys=True, ensure_ascii=False)
                + "\n```\n"
            )
        choice = re.search(r"\[OneOf: ([^;\]]+)", desc)
        if choice:
            members = [s.strip() for s in choice.group(1).split(",")]
            parent = path[:-1]
            mapped = [self.coverage.get(".".join((*parent, s))) for s in members]
            links = [
                self.link(m["document_id"], s, m["anchor"]) if m else f"`{s}`"
                for s, m in zip(members, mapped, strict=True)
            ]
            lines.append("OneOf alternatives in this subsection:\n")
            lines.extend("- " + link for link in links)
            lines.append(
                "\nSelect alternatives according to the provider validators above.\n"
            )
        if not field.get("computed") or field.get("optional") or field.get("required"):
            attrs = object_attributes(typ)
            empty = attrs == {} or (syntax == "block" and not children(block))
            if syntax == "block":
                example = (
                    f"{name} {{\n  # Configure direct properties listed below.\n}}"
                    if not empty
                    else f"{name} {{}}"
                )
            elif empty:
                example = f"{name} = {{}}"
            elif (
                isinstance(typ, list)
                and typ[0] in ("list", "set")
                and attrs is not None
            ):
                example = f"{name} = [{{\n  # Configure object properties.\n}}]"
            elif attrs is not None:
                example = f"{name} = {{\n  # Configure object properties.\n}}"
            else:
                example = None
            if example:
                lines.append(
                    "Terraform syntax:\n\n```terraform\n" + example + "\n```\n"
                )
        return "\n".join(lines)

    def body(self, page):
        role, path = page["role"], tuple(page["schema_path"])
        ancestors = []
        current = page
        while current.get("parent_id"):
            current = self.pages[current["parent_id"]]
            ancestors.append(self.link(current["id"]))
        lines = [
            f"# {page['title']}\n",
            "Breadcrumbs:\n\n"
            + "\n".join("- " + item for item in [*reversed(ancestors), page["title"]])
            + "\n",
        ]
        if role == "fundamentals":
            lines += [
                description_markdown(self.description) + "\n",
                "## Prerequisites\n",
                "Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.\n",
            ]
            if self.metadata.get("tier"):
                lines.append(f"Required service tier: {self.metadata['tier']}.\n")
            dependencies = self.metadata.get("dependencies", {})
            for label, key in (
                ("Required dependencies", "required"),
                ("Optional integrations", "optional"),
            ):
                if dependencies.get(key):
                    lines.append(
                        label
                        + ": "
                        + ", ".join(f"`{name}`" for name in dependencies[key])
                        + ".\n"
                    )
            for hint in self.metadata.get("relationship_hints", []):
                lines.append("- " + hint + "\n")
            lines += [
                "## Minimal configuration\n",
                "Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.\n",
                "```terraform\n"
                + self.pages[self.minimal]["config"].rstrip()
                + "\n```\n",
                "## Root configuration\n",
            ]
            root_fields = children(schema_root(self.schema))
            required = [
                f"`{name}`"
                for name, (field, _, _) in root_fields.items()
                if field.get("required") or field.get("min_items", 0) > 0
            ]
            lines.append(
                "Required root properties: "
                + (", ".join(required) or "none")
                + ". Full root flags and choices appear in the property reference.\n"
            )
        elif role in ("reference", "properties"):
            if path:
                lines += [
                    '<a id="section"></a>\n',
                    self.field_text(
                        path, page["section"][0], page["section"][1], heading=False
                    ),
                ]
            lines.append("## Direct properties\n")
            for name, field, subsection, syntax in page["fields"]:
                exact = (*path, name)
                if subsection:
                    target = self.coverage[".".join(exact)]
                    lines.append(
                        "- "
                        + self.link(target["document_id"], name)
                        + ": complete subsection reference.\n"
                    )
                else:
                    lines.append(self.field_text(exact, field, syntax))
            if not page["fields"]:
                lines.append(
                    "This is an empty object or choice marker. It has no direct properties.\n"
                )
            if role == "reference":
                lines += [
                    "## All schema paths\n",
                    "Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.\n",
                    "| Schema path | Complete reference |\n| --- | --- |",
                ]
                for exact, target in sorted(self.coverage.items()):
                    lines.append(
                        f"| `{exact}` | {self.link(target['document_id'], exact, target['anchor'])} |"
                    )
                lines.append("")
        elif role == "examples":
            lines.append("## Complete configurations\n")
            for identifier in page["examples"]:
                ex = self.pages[identifier]
                lines.append(
                    "- "
                    + self.link(identifier)
                    + ": "
                    + ex["evidence"]["outcome"]
                    + ".\n"
                )
        elif role in ("example", "negative-example"):
            ev = page["evidence"]
            lines += [
                ev["attribution"] + "\n",
                "Expected outcome: **" + ev["outcome"] + "**.\n",
                f"Source: `{ev['source_path']}`; digest `{ev['sha256']}`.\n",
                "```terraform\n" + page["config"].rstrip() + "\n```\n",
            ]
        elif role == "import":
            lines += [
                page["import_guidance"] + "\n",
                "```shell\n" + page["import"].rstrip() + "\n```\n",
            ]
        elif role == "timeouts":
            lines.append(
                "Configure the supported operation timeouts in the "
                + self.link(
                    stable_id(self.kind, self.name, "properties", ("timeouts",))
                )
                + ". Use Terraform duration strings such as `30m`.\n"
            )
        elif role == "lifecycle":
            if self.kind == "resources":
                for group, members in sorted(
                    self.specs.immutable_oneof_groups(self.name).items()
                ):
                    links = ",\n".join(
                        self.link(
                            stable_id(self.kind, self.name, "properties", (member,))
                        )
                        for member in members
                    )
                    lines.append(
                        f"Changing the {group} selection between\n{links},\n"
                        "including an omitted selection, requires recreation and may interrupt service.\n"
                        "F5 Distributed Cloud cannot change this type selection in place.\n"
                        "Certificate rotation and other supported settings within the same selected type remain updates.\n"
                    )
                lines.append(
                    "Terraform identifies the affected type blocks as **must be replaced**. Unknown block presence requires replacement when unchanged selection cannot be proven; unknown child settings alone do not. Recommendations do not establish API defaults.\n"
                )
                lines.append(
                    "Use `lifecycle { prevent_destroy = true }` to reject replacement before remote writes. Terraform controls replacement ordering: its default destroys before creating; `create_before_destroy` requests creation first, which may fail if XC requires a unique name. Plan a maintenance window or use a distinct name for a staged migration.\n"
                )
            elif self.kind == "actions":
                lines.append(
                    "Invoke this action using Terraform action triggers or `terraform apply -invoke`. The action executes its documented operation; it does not maintain a resource lifecycle. Inspect asynchronous operations separately where described by the API.\n"
                )
            else:
                lines.append(
                    "Terraform opens this ephemeral resource during evaluation and closes it when its lifecycle ends. Ephemeral values are not stored in plans or state. Use returned credentials only in contexts that permit ephemeral values. Sensitive flags remain visible in the property reference.\n"
                )
        next_ids = list(page["child_ids"])
        if page.get("parent_id"):
            next_ids.append(page["parent_id"])
        if role != "fundamentals":
            next_ids.append(stable_id(self.kind, self.name, "fundamentals"))
        lines.append("## Next pages\n")
        lines.extend(
            "- " + self.link(identifier) for identifier in dict.fromkeys(next_ids)
        )
        return (
            "<!-- Exact provider and upstream contract identifiers. -->\n\n<!-- textlint-disable terminology -->\n\n"
            + normalize_body("\n".join(lines).rstrip() + "\n")
        )


def frontmatter(page, body, category=""):
    metadata = {
        k: v
        for k, v in page.items()
        if k
        not in ("fields", "section", "config", "import", "examples", "title", "body")
    }
    metadata.update(
        schema_version=1,
        body_bytes=len(body.encode()),
        body_sha256=digest(body.encode()),
    )
    return (
        "---\npage_title: "
        + json_text(page["title"])
        + "\nsubcategory: "
        + json_text(category)
        + "\ndescription: "
        + json_text(page["summary"])
        + "\nxcsh_docs: "
        + json_text(metadata)
        + "\n---\n\n"
        + body
    )


def projection_name(page):
    name = "--".join(
        [
            page["provider_type"],
            page["provider_name"],
            page["role"],
            *page["schema_path"],
        ]
    )
    if len((name + ".md").encode()) > REGISTRY_FILENAME_BUDGET:
        suffix = hashlib.sha256(page["id"].encode()).hexdigest()[:32]
        name = (
            name.encode()[:190].decode("utf-8", errors="ignore").rstrip("-")
            + "--"
            + suffix
        )
    return name


def registry_project(pages, categories):
    outputs, _ = grouped_project(pages, categories, SITE)
    return outputs


def validate(pages, collections, outputs):
    identifiers = {p["id"]: p for p in pages}
    if len(identifiers) != len(pages):
        raise ValueError("duplicate page IDs")
    for page in pages:
        for child in page["child_ids"]:
            if identifiers[child]["parent_id"] != page["id"]:
                raise ValueError(f"broken relationship: {child}")
        if (
            page["parent_id"]
            and page["id"] not in identifiers[page["parent_id"]]["child_ids"]
        ):
            raise ValueError("missing parent edge")
    for page in pages:
        for section in page.get("sections", []):
            target = identifiers.get(section["document_id"])
            if not target or f'id="{section["anchor"]}"' not in outputs[target["path"]]:
                raise ValueError("invalid retrieval section destination")
        relations = [
            *page.get("relationships", []),
            *[
                r
                for section in page.get("sections", [])
                for r in section.get("relationships", [])
            ],
        ]
        for relation in relations:
            target = identifiers.get(relation["target_id"])
            if not target or (
                relation["anchor"]
                and f'id="{relation["anchor"]}"' not in outputs[target["path"]]
            ):
                raise ValueError("invalid retrieval relationship destination")
    urls = {
        SITE
        + "/"
        + p["path"].removeprefix("documentation/").removesuffix("index.md"): p
        for p in pages
    }
    canonical_bodies = {url: outputs[page["path"]] for url, page in urls.items()}
    for output, text in outputs.items():
        if not output.endswith(".md"):
            continue
        if output.startswith("docs/") and len(text.encode()) > REGISTRY_LIMIT:
            raise ValueError(f"Registry storage ceiling exceeded: {output}")
        for _, href in prose_links(text):
            target, _, anchor = href.partition("#")
            if target in (SITE + "/llms.txt", SITE + "/terraform-llms-index.json"):
                continue
            if target.startswith(SITE + "/"):
                if target not in urls:
                    raise ValueError(f"unresolved canonical link: {target}")
                if anchor and f'id="{anchor}"' not in canonical_bodies[target]:
                    raise ValueError(f"unresolved canonical anchor: {target}#{anchor}")
            elif not re.match(r"[a-z]+:", target) and target:
                resolved = os.path.normpath(str(PurePosixPath(output).parent / target))
                if resolved not in outputs:
                    raise ValueError(f"unresolved projected link: {output} -> {target}")
                if anchor and f'id="{anchor}"' not in outputs[resolved]:
                    raise ValueError(f"unresolved projected anchor: {output} -> {href}")
    for collection in collections:
        expected: dict[str, list[str]] = {}

        def walk(block, path, destinations, recurse):
            for name, (_, shape, _) in children(block).items():
                exact = (*path, name)
                destinations[".".join(exact)] = list(exact)
                if shape is not None:
                    recurse(shape, exact, destinations, recurse)

        walk(schema_root(collection.schema), (), expected, walk)
        if set(expected) != set(collection.coverage):
            raise ValueError(
                f"schema coverage mismatch: {collection.kind}/{collection.name}"
            )
        for path, target in collection.coverage.items():
            if (
                target["schema_path"] != expected[path]
                or target["document_id"] not in identifiers
            ):
                raise ValueError(f"invalid property destination: {path}")


def registry_navigation(surface, outputs):
    """Own navigation indexes and validate every exact release-surface target."""
    for kind, (_, key, _, _) in TYPES.items():
        names = sorted(surface[key])
        title = kind.replace("-", " ").title()
        links = []
        for name in names:
            target = f"docs/{kind}/{name}.md"
            if target not in outputs:
                raise ValueError(f"missing navigation target: {target}")
            links.append(f"- [xcsh_{name}]({name}.md)")
        related = [
            f"- [{other.replace(chr(45), chr(32)).title()}](../{other}/index.md)"
            for other in TYPES
            if other != kind
        ]
        outputs[f"docs/{kind}/index.md"] = (
            f"# {title}\n\nThis provider includes {len(names)} {kind.replace(chr(45), chr(32))}.\n\n"
            + "\n".join(links)
            + "\n\n## Related Documentation\n\n"
            + "\n".join(related)
            + "\n"
        )


def generate(root, schema_path, constraints_path):
    schema_bytes = Path(schema_path).read_bytes()
    provider = json.loads(schema_bytes)["provider_schemas"][PROVIDER]
    # Canonical JSON digest ignores Terraform map iteration ordering.
    schema_digest = digest(
        json.dumps(provider, sort_keys=True, separators=(",", ":")).encode()
    )
    constraints = read_json(constraints_path)
    specs = Specs(root)
    surface = read_json(root / "provider-release-surface.json")
    collections, pages, outputs, categories = [], [], {}, {}
    for kind, (schema_key, surface_key, _, source_suffix) in TYPES.items():
        schemas = provider.get(schema_key, {})
        expected = {"xcsh_" + name for name in surface[surface_key]}
        if set(schemas) != expected:
            raise ValueError(f"installed {kind} do not match release surface")
        for full_name, schema in sorted(schemas.items()):
            name = full_name.removeprefix("xcsh_")
            source_fields = constraints.get(f"{name}_{source_suffix}.go", {})
            collection = Collection(
                kind, name, schema, specs, source_fields, root, schema_digest
            )
            collections.append(collection)
            categories[stable_id(kind, name, "collection")] = collection.category
            for page in collection.pages.values():
                page["source_url"] = (
                    SITE
                    + "/_data/pages/"
                    + page["path"].removeprefix("documentation/").removesuffix(".md")
                    + ".txt"
                )
                body = collection.body(page)
                page.update(
                    body=body,
                    body_bytes=len(body.encode()),
                    body_sha256=digest(body.encode()),
                )
                outputs[page["path"]] = frontmatter(page, body, collection.category)
                pages.append(page)
    registry_outputs, projection = grouped_project(
        pages, categories, SITE, version=os.environ.get("DOCUMENTATION_VERSION")
    )
    outputs.update(registry_outputs)
    outputs["documentation/registry-projection-manifest.json"] = (
        json.dumps(projection, indent=2, sort_keys=True) + "\n"
    )
    registry_navigation(surface, outputs)
    provider_index = (root / "templates/index.md.tmpl").read_text(encoding="utf-8")
    provider_example = (root / "examples/provider/provider.tf").read_text(
        encoding="utf-8"
    )
    provider_index = provider_index.replace(
        '{{ tffile "examples/provider/provider.tf" }}',
        "```terraform\n" + provider_example.rstrip() + "\n```",
    )
    provider_index = provider_index.replace("(llms.txt)", "(" + SITE + "/llms.txt)")
    outputs["docs/index.md"] = provider_index
    auxiliary = [("provider", "setup", provider_index)]
    for guide in sorted((root / "templates/guides").glob("*.md")):
        text = guide.read_text(encoding="utf-8")
        text = text.replace(
            "../../examples/",
            "https://github.com/f5-sales-demo/terraform-provider-xcsh/blob/"
            + os.environ.get("DOCUMENTATION_VERSION", "main")
            + "/examples/",
        )
        outputs["docs/guides/" + guide.name] = text
        auxiliary.append(("guides", guide.stem, text))
    for kind, name, text in auxiliary:
        body = text.split("---\n", 2)[-1].lstrip() if text.startswith("---\n") else text
        identifier = stable_id(kind, name, "overview")
        page = {
            "id": identifier,
            "collection_id": stable_id(kind, name, "collection"),
            "provider_type": kind,
            "provider_name": name,
            "role": "overview",
            "schema_path": [],
            "parent_id": None,
            "child_ids": [],
            "title": "Provider setup and authentication"
            if kind == "provider"
            else name,
            "summary": "Complete provider setup and authentication."
            if kind == "provider"
            else "Maintained " + name + " guide.",
            "aliases": [],
            "completeness": "complete",
            "path": f"documentation/{kind}/{name}/index.md",
            "body": body,
            "body_bytes": len(body.encode()),
            "body_sha256": digest(body.encode()),
            "provider_schema_digest": schema_digest,
            "spec_pin_digest": specs.pin_digest,
            "registry_path": "docs/index.md"
            if kind == "provider"
            else f"docs/guides/{name}.md",
        }
        rules = RetrievalRules.default()
        page.update(rules.classify(kind, name, [], "overview"))
        page["aliases"] = rules.aliases(name, body)
        page["sections"] = []
        page["relationships"] = []
        pages.append(page)
        outputs[page["path"]] = frontmatter(page, body)
        projection["sections"].append(
            {
                "canonical_id": identifier,
                "section_id": identifier,
                "source_sha256": digest(body.encode()),
                "source_bytes": len(body.encode()),
                "canonical_source_sha256": digest(body.encode()),
                "registry_path": page["registry_path"],
                "mode": "embedded",
                "canonical_url": SITE + f"/{kind}/{name}/",
                "anchor_map": {},
                "anchor": "",
            }
        )
    projection["provider_schema_digest"] = schema_digest
    projection["spec_pin_digest"] = specs.pin_digest
    projection["files"] = {
        path: {"bytes": len(text.encode()), "sha256": digest(text.encode())}
        for path, text in sorted(outputs.items())
        if path.startswith("docs/")
    }
    outputs["documentation/registry-projection-manifest.json"] = (
        json.dumps(projection, indent=2, sort_keys=True) + "\n"
    )
    shards: list[list[dict]] = []
    pending: list[dict] = []
    pending_bytes = 0
    for section in projection.pop("sections"):
        size = len(json.dumps(section, ensure_ascii=False).encode()) + 1
        if pending and pending_bytes + size > PROJECTION_RECEIPT_BUDGET:
            shards.append(pending)
            pending, pending_bytes = [], 0
        pending.append(section)
        pending_bytes += size
    if pending:
        shards.append(pending)
    projection["section_manifests"] = []
    for number, records in enumerate(shards, 1):
        path = f"documentation/registry-projection/sections-{number:04}.json"
        text = (
            json.dumps(
                {"schema_version": 1, "sections": records},
                separators=(",", ":"),
                ensure_ascii=False,
            )
            + "\n"
        )
        outputs[path] = text
        projection["section_manifests"].append(
            {"path": path, "section_count": len(records)}
        )
    outputs["documentation/registry-projection-manifest.json"] = (
        json.dumps(projection, sort_keys=True) + "\n"
    )
    root_id = stable_id("provider", "xcsh", "navigation")
    for kind in ("provider", *TYPES):
        identifier = (
            root_id if kind == "provider" else stable_id(kind, "xcsh", "navigation")
        )
        owned = [
            p for p in pages if p["provider_type"] == kind and p["parent_id"] is None
        ]
        if kind == "provider":
            owned += [
                p
                for p in pages
                if p["provider_type"] == "guides" and p["parent_id"] is None
            ]
        children_ids = [p["id"] for p in owned]
        if kind == "provider":
            children_ids += [stable_id(t, "xcsh", "navigation") for t in TYPES]
        for child in owned:
            child["parent_id"] = identifier
        body = (
            "# "
            + ("xcsh provider documentation" if kind == "provider" else kind)
            + "\n\n"
            + "\n".join(
                "- ["
                + p["title"]
                + "]("
                + SITE
                + "/"
                + p["path"].removeprefix("documentation/").removesuffix("index.md")
                + ")"
                for p in owned
            )
            + "\n"
        )
        if kind == "provider":
            body += (
                "\n"
                + "\n".join("- [" + t + "](" + SITE + "/" + t + "/)" for t in TYPES)
                + "\n"
            )
        page = {
            "id": identifier,
            "collection_id": stable_id(kind, "xcsh", "collection"),
            "provider_type": kind,
            "provider_name": "xcsh",
            "role": "navigation",
            "schema_path": [],
            "parent_id": None if kind == "provider" else root_id,
            "child_ids": children_ids,
            "title": kind,
            "summary": "Navigate " + kind + " documentation collections.",
            "aliases": [],
            "completeness": "complete",
            "path": "documentation/index.md"
            if kind == "provider"
            else f"documentation/{kind}/index.md",
            "body": body,
            "body_bytes": len(body.encode()),
            "body_sha256": digest(body.encode()),
            "provider_schema_digest": schema_digest,
            "spec_pin_digest": specs.pin_digest,
            "sections": [],
            "relationships": [],
            **RetrievalRules.default().classify(kind, "xcsh", [], "navigation"),
        }
        pages.append(page)
        outputs[page["path"]] = frontmatter(page, body)
    by_collection = {(c.kind, c.name): c for c in collections}
    for collection in collections:
        page = collection.pages[
            stable_id(collection.kind, collection.name, "fundamentals")
        ]
        page["unresolved_relationships"] = []
        for kind, names in collection.metadata.get("dependencies", {}).items():
            if kind not in ("required", "optional"):
                continue
            for name in names:
                target_collection = by_collection.get(("resources", name))
                if target_collection:
                    target_id = stable_id("resources", name, "fundamentals")
                    page["relationships"].append(
                        {
                            "type": "advisory",
                            "target_id": target_id,
                            "anchor": "",
                            "enforcement": "upstream-advisory",
                            "source": "receipt-pinned-dependency:" + kind,
                        }
                    )
                else:
                    page["unresolved_relationships"].append(
                        {"name": name, "source": "receipt-pinned-dependency:" + kind}
                    )
    # Serialize after adding navigation parents, preserving the body bytes.
    for page in pages:
        outputs[page["path"]] = frontmatter(
            page, page["body"], categories.get(page["collection_id"], "")
        )
    validate(pages, collections, outputs)
    entries = [
        {
            k: v
            for k, v in page.items()
            if k not in ("fields", "section", "config", "import", "examples", "body")
        }
        for page in pages
    ]
    index = {
        "schema_version": 1,
        "provider": PROVIDER,
        "provider_schema_digest": schema_digest,
        "spec_pin_digest": specs.pin_digest,
        "collections": [
            {
                "id": stable_id(c.kind, c.name, "collection"),
                "provider_type": c.kind,
                "provider_name": c.name,
                "fundamentals_id": stable_id(c.kind, c.name, "fundamentals"),
                "reference_id": stable_id(c.kind, c.name, "reference"),
                "properties": c.coverage,
            }
            for c in collections
        ],
        "pages": entries,
    }
    outputs["documentation/terraform-llms-index.json"] = (
        json.dumps(index, sort_keys=True, separators=(",", ":"), ensure_ascii=False)
        + "\n"
    )
    outputs["documentation/llms.txt"] = (
        "# xcsh provider documentation\n\nLoad fundamentals first, then the property reference or immediate children. Fetch complete relevant leaves. The machine index supplies exact paths, relationships and digests.\n\n"
        f"- [Machine index]({SITE}/terraform-llms-index.json)\n\n"
        + "\n".join(
            f"- [{c.kind}/{c.name}]({SITE}/{c.kind}/{c.name}/): {c.description.splitlines()[0]}"
            for c in collections
        )
        + "\n- [Provider setup and authentication]("
        + SITE
        + "/provider/setup/)\n"
        + "".join(
            f"- [{name} guide]({SITE}/guides/{name}/)\n"
            for kind, name, _ in auxiliary
            if kind == "guides"
        )
    )
    # The Pages builder copies only _data static assets; publish full Markdown
    # bytes there for exact AI retrieval without using transformed text.
    for page in pages:
        raw_path = (
            "documentation/_data/pages/"
            + page["path"].removeprefix("documentation/").removesuffix(".md")
            + ".txt"
        )
        outputs[raw_path] = outputs[page["path"]]
        page["source_url"] = (
            SITE
            + "/_data/pages/"
            + page["path"].removeprefix("documentation/").removesuffix(".md")
            + ".txt"
        )
    for entry, page in zip(entries, pages, strict=True):
        entry["source_url"] = page["source_url"]
    for collection in index["collections"]:
        collection_id = collection["id"]
        own_pages = [page for page in entries if page["collection_id"] == collection_id]
        collection_index = {
            "schema_version": 1,
            "provider": PROVIDER,
            "collection": collection,
            "pages": own_pages,
        }
        relative = f"{collection['provider_type']}/{collection['provider_name']}/collection.json"
        outputs["documentation/_data/collections/" + relative] = (
            json.dumps(
                collection_index,
                sort_keys=True,
                separators=(",", ":"),
                ensure_ascii=False,
            )
            + "\n"
        )
        collection["index_url"] = SITE + "/_data/collections/" + relative
    index_text = (
        json.dumps(index, sort_keys=True, separators=(",", ":"), ensure_ascii=False)
        + "\n"
    )
    outputs["documentation/terraform-llms-index.json"] = index_text
    outputs["documentation/_data/terraform-llms-index.json"] = index_text
    outputs["documentation/_data/llms.txt"] = outputs["documentation/llms.txt"]
    manifest_path = root / "documentation/generated-manifest.json"
    old = read_json(manifest_path).get("files", {}) if manifest_path.exists() else {}
    for relative in old:
        candidate = PurePosixPath(relative)
        if (
            candidate.is_absolute()
            or ".." in candidate.parts
            or candidate.parts[0] not in ("documentation", "docs")
        ):
            raise ValueError(f"unsafe manifest-owned path: {relative}")
        owned = root / relative
        if owned.is_symlink() or root.resolve() not in owned.resolve().parents:
            raise ValueError(f"manifest path escapes output: {relative}")
        if relative not in outputs and owned.is_file():
            if digest(owned.read_bytes()) != old[relative]["sha256"]:
                raise ValueError(
                    f"obsolete generated page has unowned edits: {relative}"
                )
            owned.unlink()
    legacy = read_json(root / "tools/legacy-docs-manifest.json")
    for relative, evidence in legacy["files"].items():
        candidate = PurePosixPath(relative)
        if (
            candidate.is_absolute()
            or ".." in candidate.parts
            or not relative.startswith("docs/")
        ):
            raise ValueError(f"unsafe legacy output path: {relative}")
        owned = root / relative
        if owned.exists():
            if owned.is_symlink() or digest(owned.read_bytes()) != evidence["sha256"]:
                raise ValueError(f"legacy output contains unowned edits: {relative}")
            owned.unlink()
    formatter = [shutil.which("biome") or "/usr/bin/biome"]
    if os.environ.get("PROVIDER_FORK_ISOLATION") == "true":
        formatter = [
            shutil.which("npx") or "/usr/bin/npx",
            "--yes",
            "@biomejs/biome@2.5.6",
        ]
    else:
        version = subprocess.run(  # noqa: S603 - verified local tool, fixed arguments
            [formatter[0], "--version"], check=True, capture_output=True, text=True
        ).stdout.strip()
        if version not in ("Version: 2.5.6", "2.5.6"):
            raise ValueError("documentation formatter must be Biome 2.5.6")
    for relative, text in list(outputs.items()):
        # Large machine catalogs retain compact deterministic JSON, avoiding
        # repeated indentation bytes and GitHub single-file storage limits.
        if relative in (
            "documentation/terraform-llms-index.json",
            "documentation/_data/terraform-llms-index.json",
        ):
            outputs[relative] = (
                json.dumps(
                    json.loads(text),
                    sort_keys=True,
                    separators=(",", ":"),
                    ensure_ascii=False,
                )
                + "\n"
            )
            continue
        if relative.endswith(".json"):
            formatted = subprocess.run(  # noqa: S603 - generated path and verified formatter
                [
                    *formatter,
                    "format",
                    "--stdin-file-path=" + relative,
                    "--files-max-size=100000000",
                ],
                input=text,
                text=True,
                capture_output=True,
                check=True,
            )
            outputs[relative] = formatted.stdout
    for relative, text in outputs.items():
        path = root / relative
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text, encoding="utf-8")
    manifest = {
        "schema_version": 1,
        "spec_pin_digest": specs.pin_digest,
        "provider_schema_digest": schema_digest,
        "files": {
            path: {"bytes": len(text.encode()), "sha256": digest(text.encode())}
            for path, text in sorted(outputs.items())
        },
    }
    manifest_path.write_text(json.dumps(manifest, indent=2, sort_keys=True) + "\n")
    http = next(
        c
        for c in collections
        if c.kind == "resources" and c.name == "http_loadbalancer"
    )
    print(
        f"Generated {len(collections)} collections, {len(pages)} pages, {sum(len(c.coverage) for c in collections)} exact property destinations."
    )
    print(
        f"Registry: {len(projection['files'])} documents, maximum {max(item['bytes'] for item in projection['files'].values())} bytes; {sum(item['mode'] == 'canonical-link' for shard in shards for item in shard)} oversized exceptions."
    )
    print(
        f"HTTP load balancer: {len(http.coverage)} properties; fundamentals {http.pages[stable_id(http.kind, http.name, 'fundamentals')]['body_bytes']} bytes."
    )


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--root", type=Path, default=Path.cwd())
    parser.add_argument("--schema", type=Path, required=True)
    parser.add_argument("--constraints", type=Path, required=True)
    args = parser.parse_args()
    generate(args.root, args.schema, args.constraints)


if __name__ == "__main__":
    try:
        main()
    except (ValueError, KeyError, OSError) as error:
        print(f"documentation generation failed: {error}", file=sys.stderr)
        sys.exit(1)
