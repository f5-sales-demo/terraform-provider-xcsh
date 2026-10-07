---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_signatures_update."
xcsh_docs: {"aliases": ["site signatures update"], "body_bytes": 1121, "body_sha256": "sha256:9c105a480af8c19d415cca776bd99bf3def7dfe693d3e7f440985c631d8290c5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_signatures_update:reference", "parent_id": "xcsh-docs:actions:site_signatures_update:fundamentals", "path": "documentation/actions/site_signatures_update/properties/index.md", "product": "distributed-cloud", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "actions", "registry_anchor": "canonical-2103311202213331-3322033031022121-3121001201130323-0111123021202301-2313002222033000-2333220012121133-2230132001320330-3233213301303202", "registry_path": "docs/guides/actions--site_signatures_update--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace of the site to update signatures for.", "document_id": "xcsh-docs:actions:site_signatures_update:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["site names"], "anchor": "schema-site_names", "description": "List of site names to update. If empty, no sites will be updated.", "document_id": "xcsh-docs:actions:site_signatures_update:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_site_signatures_update.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_signatures_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/)
- Property reference

## Direct properties

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace of the site to update signatures for.

<a id="schema-site_names"></a>

### site_names property

Type: `["list", "string"]`. Optional.

List of site names to update. If empty, no sites will be updated.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/properties/#schema-namespace) |
| `site_names` | [site_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/properties/#schema-site_names) |
