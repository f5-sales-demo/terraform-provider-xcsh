---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": ["infraprotect mitigation ips"], "body_bytes": 2058, "body_sha256": "sha256:37d6d84be4a7e4af1415040d33b7d7dd43f953e21460f93c068a8711805c1c7b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "path": "documentation/data-sources/infraprotect_mitigation_ips/properties/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3211003101332320-2011321310020021-1301100203211023-1330101230021031-1102102332023332-0031020312103202-2022313110031312-3331302220003310", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["ips"], "anchor": "section", "description": "List of IP addresses along with log count.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ips"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation id"], "anchor": "schema-mitigation_id", "description": "Mitigation ID ID of the mitigation we want to GET the IPs for.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mitigation_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace This request is supported only in system namespace.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_infraprotect_mitigation_ips.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/ips/)
- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
