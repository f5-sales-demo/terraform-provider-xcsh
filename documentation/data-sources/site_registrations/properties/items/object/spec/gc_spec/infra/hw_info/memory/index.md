---
page_title: "items.object.spec.gc_spec.infra.hw_info.memory"
subcategory: ""
description: "Memory Information. Memory information."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info memory"], "body_bytes": 1946, "body_sha256": "sha256:79fb3b91da82acdec27819142cc4c0bfbeb15c90cd8945cf3b4b1ed95c355e10", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:memory", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/memory/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3222133202212201-2100233211220010-1201111332302213-1133003122120111-2103121310012201-0210222203032313-3131013331100330-1122200113131021", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec infra hw info memory size mb"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--memory--size_mb", "description": "RAM. RAM size in MB.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory", "size_mb"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object spec gc spec infra hw info memory speed"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--memory--speed", "description": "Speed. RAM data rate in MT/s.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object spec gc spec infra hw info memory type"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--memory--type", "description": "Type. Type of memory, eg. DDR4.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/memory/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Memory Information. Memory information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.memory

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.memory

<a id="section"></a>

Type: `"single"`. Computed.

Memory Information. Memory information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--memory--size_mb"></a>

### size_mb property

Type: `"number"`. Computed.

RAM. RAM size in MB.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--memory--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. RAM data rate in MT/s.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--memory--type"></a>

### type property

Type: `"string"`. Computed.

Type. Type of memory, eg. DDR4.
