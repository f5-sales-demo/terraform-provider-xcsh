---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": ["infraprotect mitigation ips"], "body_bytes": 1784, "body_sha256": "sha256:7029c9a93a4b19293fee8923bfb06a3c89dbb8e49546a16605ee71cfbf6c72e9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "path": "documentation/data-sources/infraprotect_mitigation_ips/properties/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3211003101332320-2011321310020021-1301100203211023-1330101230021031-1102102332023332-0031020312103202-2022313110031312-3331302220003310", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["ips"], "anchor": "section", "description": "List of IP addresses along with log count.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ips"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation id"], "anchor": "schema-mitigation_id", "description": "Mitigation ID ID of the mitigation we want to GET the IPs for.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace This request is supported only in system namespace.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/properties/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Property reference for xcsh_infraprotect_mitigation_ips.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
- Property reference

## Direct properties

- [ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/ips/): complete subsection reference.

<a id="schema-mitigation_id"></a>

### mitigation_id property

Type: `"string"`. Required.

Mitigation ID ID of the mitigation we want to GET the IPs for.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace This request is supported only in system namespace.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `ips` | [ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/ips/#section) |
| `ips.ip` | [ips.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/ips/#schema-ips--ip) |
| `ips.log_count` | [ips.log_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/ips/#schema-ips--log_count) |
| `mitigation_id` | [mitigation_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/#schema-mitigation_id) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/#schema-namespace) |
