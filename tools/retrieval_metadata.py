# Schema dictionaries intentionally use dynamic types, as in generate-doc-collections.py.
# ruff: noqa: INP001, ANN001, ANN201, ANN204, D101, D102, D107, EM101, TRY003
"""Reviewed, versioned retrieval hints; source Markdown remains authoritative."""

import hashlib
import json
import re
from pathlib import Path

VERSION = 1
SUMMARY_LIMIT = 320
TASKS = {"configuration", "troubleshooting", "import", "authentication", "lifecycle"}
CATEGORIES = {
    "load-balancing",
    "security",
    "networking",
    "infrastructure",
    "container",
    "monitoring",
    "dns",
    "api-management",
    "identity",
    "cdn",
    "administration",
}


def summary(description, fallback):
    """Keep useful prose, excluding validator annotations and markup."""
    prose = re.sub(r"\[[^\]]*\]", "", description)
    prose = re.sub(r"\s+", " ", prose).strip()
    text = prose or fallback
    return (
        text if len(text) <= SUMMARY_LIMIT else text[:SUMMARY_LIMIT].rsplit(" ", 1)[0]
    )


class RetrievalRules:
    def __init__(self, data):
        if data.get("version") != VERSION:
            raise ValueError("unsupported retrieval rules version")
        self.rules = data.get("rules", [])
        self.terms = data.get("terms", [])
        self.summaries = data.get("summaries", [])
        if any(
            term.get("match_source", "description")
            not in {"identifier", "description", "schema_path"}
            for term in self.terms
        ):
            raise ValueError("invalid retrieval term match source")
        summaries_seen: dict[tuple, dict] = {}
        for rule in self.summaries:
            if (
                not isinstance(rule, dict)
                or not isinstance(rule.get("schema_path"), list)
                or not all(isinstance(part, str) for part in rule["schema_path"])
                or not all(
                    isinstance(rule.get(key), str) and rule[key].strip()
                    for key in (
                        "provider_type",
                        "collection",
                        "summary",
                        "review_basis",
                    )
                )
            ):
                raise ValueError("invalid reviewed retrieval summary")
            if not isinstance(rule.get("aliases", []), list) or not all(
                isinstance(alias, str) and alias.strip()
                for alias in rule.get("aliases", [])
            ):
                raise ValueError("invalid reviewed retrieval aliases")
            key = (
                rule["provider_type"],
                rule["collection"],
                tuple(rule["schema_path"]),
            )
            if key in summaries_seen and summaries_seen[key] != rule:
                raise ValueError("conflicting reviewed retrieval summaries")
            summaries_seen[key] = rule
        self.digest = (
            "sha256:"
            + hashlib.sha256(
                json.dumps(data, sort_keys=True, separators=(",", ":")).encode()
            ).hexdigest()
        )
        seen: dict[tuple, dict] = {}
        for rule in self.rules:
            key = (
                rule.get("provider_type", "*"),
                rule.get("collection", "*"),
                tuple(rule.get("schema_prefix", [])),
            )
            if key in seen and seen[key] != rule:
                raise ValueError("conflicting retrieval rules: " + str(key))
            seen[key] = rule
            if rule.get("category") and rule["category"] not in CATEGORIES:
                raise ValueError("invalid retrieval category")
            if any(t not in TASKS for t in rule.get("tasks", [])):
                raise ValueError("invalid retrieval task")
            if any(
                not re.fullmatch(r"[a-z][a-z0-9-]*(?:\.[a-z][a-z0-9-]*)*", c)
                for c in rule.get("capabilities", [])
            ):
                raise ValueError("invalid retrieval capability")

    @classmethod
    def default(cls) -> "RetrievalRules":
        return cls(
            json.loads(Path(__file__).with_name("retrieval-rules.json").read_text())
        )

    def classify(
        self, provider_type, collection, schema_path, role, upstream_category=""
    ):
        category = (
            re.sub(r"[^a-z0-9]+", "-", upstream_category.lower()).strip("-") or None
        )
        if category == "platform":
            category = "administration"
        elif category == "operations":
            category = "monitoring"
        if category not in CATEGORIES:
            category = None
        capabilities, tasks, sources = set(), set(), []
        if category:
            sources.append("receipt-pinned-upstream")
        matches = []
        for rule in self.rules:
            prefix = rule.get("schema_prefix", [])
            if (
                rule.get("provider_type", "*") in ("*", provider_type)
                and rule.get("collection", "*") in ("*", collection)
                and schema_path[: len(prefix)] == prefix
            ):
                matches.append(rule)
        matches.sort(
            key=lambda r: (
                len(r.get("schema_prefix", [])),
                r.get("collection", "*") != "*",
                r.get("provider_type", "*") != "*",
            )
        )
        for rule in matches:
            category = rule.get("category", category)
            capabilities.update(rule.get("capabilities", []))
            tasks.update(rule.get("tasks", []))
        if matches:
            sources.append("reviewed-rule")
        if category:
            capabilities.add(category)
        if role in ("import", "timeouts", "lifecycle"):
            tasks.add("import" if role == "import" else "lifecycle")
        elif role != "navigation":
            tasks.add("configuration")
        if provider_type == "provider":
            tasks.add("authentication")
        return {
            "retrieval_version": VERSION,
            "product": "distributed-cloud",
            "category": category,
            "capabilities": sorted(capabilities),
            "tasks": sorted(tasks),
            "classification": {
                "status": "resolved" if category else "unresolved",
                "sources": sources,
                "rules_sha256": self.digest,
            },
        }

    def reviewed_rule(self, provider_type, collection, schema_path):
        matches = [
            rule
            for rule in self.summaries
            if rule["provider_type"] == provider_type
            and rule["collection"] in ("*", collection)
            and rule["schema_path"] == schema_path
        ]
        if not matches:
            return None
        specific = [rule for rule in matches if rule["collection"] == collection]
        preferred = specific or matches
        if any(rule != preferred[0] for rule in preferred[1:]):
            raise ValueError("conflicting reviewed retrieval summaries")
        return preferred[0]

    def reviewed_summary(self, provider_type, collection, schema_path):
        rule = self.reviewed_rule(provider_type, collection, schema_path)
        return rule["summary"] if rule else None

    def reviewed_aliases(self, provider_type, collection, schema_path):
        rule = self.reviewed_rule(provider_type, collection, schema_path)
        return sorted(set(rule.get("aliases", []))) if rule else []

    def aliases(self, identifier, description=""):
        prose = re.sub(r"[_\.]+", " ", identifier).lower()
        values = {prose.strip()} if prose.strip() else set()
        text = prose + " " + description.lower()
        local_identifier = identifier.rsplit(".", 1)[-1].replace("_", " ").lower()
        for term in self.terms:
            if term.get("context_pattern") and not re.search(
                term["context_pattern"], prose
            ):
                continue
            if re.search(
                term["pattern"],
                local_identifier
                if term.get("match_source") == "identifier"
                else prose
                if term.get("match_source") == "schema_path"
                else text,
            ):
                values.update(term["aliases"])
        return sorted(values)


def constraint_relationships(schema_path, constraints, coverage):
    """Only AST-extracted provider validators establish enforced relationships."""
    relationships = []
    text = constraints.get("Validators", "")
    for validator, relation in [
        ("RequiredObjectAttributes", "requires"),
        ("RequiredListObjectAttributes", "requires"),
        ("RequiredOneOfListObjectAttributes", "choice"),
        ("ConflictingListObjectAttributes", "conflicts"),
        ("ConflictingObjectAttributes", "conflicts"),
        ("ExactlyOneOf", "choice"),
        ("AtLeastOneOf", "choice"),
        ("ConflictsWith", "conflicts"),
        ("AlsoRequires", "requires"),
    ]:
        for match in re.finditer(r"\b" + validator + r"\(([^()]*)\)", text):
            names = re.findall(r'"([A-Za-z0-9_.]+)"', match.group(1))
            if not names:
                continue
            group = ".".join(schema_path) + ":" + validator + ":" + ",".join(names)
            for name in names:
                target = coverage.get(".".join([*schema_path, name]))
                if target is None:
                    target = coverage.get(name)
                if target:
                    relationships.append(
                        {
                            "type": relation,
                            "target_id": target["document_id"],
                            "anchor": target["anchor"],
                            "enforcement": "provider-schema",
                            "source": "ast-validator:" + validator,
                            "group": group,
                        }
                    )
    return sorted(
        relationships,
        key=lambda r: (r["type"], r["target_id"], r["anchor"], r["group"]),
    )
