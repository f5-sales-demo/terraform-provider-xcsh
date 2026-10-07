---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_cloud_init."
xcsh_docs: {"aliases": ["site cloud init"], "body_bytes": 1910, "body_sha256": "sha256:ed13bcbd32951bbdad5188b17f986c0bde2c9ca6e854b123d5171091efae7380", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_cloud_init:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_cloud_init:reference", "parent_id": "xcsh-docs:data-sources:site_cloud_init:fundamentals", "path": "documentation/data-sources/site_cloud_init/properties/index.md", "product": "distributed-cloud", "provider_name": "site_cloud_init", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1010210110123221-2301023030310012-3131213301230313-0212223132222233-1001200020011312-1122011223101313-1011020230110010-0030103221003122", "registry_path": "docs/guides/data-sources--site_cloud_init--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["cloud init config"], "anchor": "schema-cloud_init_config", "description": "Cloud-init template with an unresolved token placeholder; substitute a separately issued site-bound JWT before deployment. This sensitive value is stored in Terraform state; protect state access accordingly.", "document_id": "xcsh-docs:data-sources:site_cloud_init:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloud_init_config"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable management network"], "anchor": "schema-enable_management_network", "description": "Management network choice for this cloud-init config.", "document_id": "xcsh-docs:data-sources:site_cloud_init:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_management_network"], "syntax": "attribute", "type": "bool"}, {"aliases": ["provider ref"], "anchor": "schema-provider_ref", "description": "Provider for that cloud-init config.", "document_id": "xcsh-docs:data-sources:site_cloud_init:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["site name"], "anchor": "schema-site_name", "description": "Site name for this cloud-init config.", "document_id": "xcsh-docs:data-sources:site_cloud_init:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_cloud_init/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_site_cloud_init.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_cloud_init](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/)
- Property reference

## Direct properties

<a id="schema-cloud_init_config"></a>

### cloud_init_config property

Type: `"string"`. Computed, Sensitive.

Cloud-init template with an unresolved token placeholder; substitute a separately issued site-bound
JWT before deployment. This sensitive value is stored in Terraform state; protect state access
accordingly.

<a id="schema-enable_management_network"></a>

### enable_management_network property

Type: `"bool"`. Optional.

Management network choice for this cloud-init config.

<a id="schema-provider_ref"></a>

### provider_ref property

Type: `"string"`. Required.

Provider for that cloud-init config.

<a id="schema-site_name"></a>

### site_name property

Type: `"string"`. Required.

Site name for this cloud-init config.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cloud_init_config` | [cloud_init_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-cloud_init_config) |
| `enable_management_network` | [enable_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-enable_management_network) |
| `provider_ref` | [provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-provider_ref) |
| `site_name` | [site_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-site_name) |
