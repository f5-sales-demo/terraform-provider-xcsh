---
page_title: "items.get_spec.infra.hw_info.memory"
subcategory: ""
description: "Memory Information. Memory information."
xcsh_docs: {"aliases": ["items get spec infra hw info memory"], "body_bytes": 1603, "body_sha256": "sha256:98cde8c16b26c177e6694c36b5402153875e976cfcc846f87f4debd5673370eb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:memory", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/memory/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1023301320233031-1211311133101303-1030312210022212-0332330111312113-2202303133100223-1302013232123310-1102001012002013-3331323002333202", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "memory"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info memory size mb"], "anchor": "schema-items--get_spec--infra--hw_info--memory--size_mb", "description": "RAM. RAM size in MB.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory", "size_mb"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info memory speed"], "anchor": "schema-items--get_spec--infra--hw_info--memory--speed", "description": "Speed. RAM data rate in MT/s.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info memory type"], "anchor": "schema-items--get_spec--infra--hw_info--memory--type", "description": "Type. Type of memory, eg. DDR4.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/memory/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Memory Information. Memory information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.memory

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.memory

<a id="section"></a>

Type: `"single"`. Computed.

Memory Information. Memory information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--memory--size_mb"></a>

### size_mb property

Type: `"number"`. Computed.

RAM. RAM size in MB.

<a id="schema-items--get_spec--infra--hw_info--memory--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. RAM data rate in MT/s.

<a id="schema-items--get_spec--infra--hw_info--memory--type"></a>

### type property

Type: `"string"`. Computed.

Type. Type of memory, eg. DDR4.
