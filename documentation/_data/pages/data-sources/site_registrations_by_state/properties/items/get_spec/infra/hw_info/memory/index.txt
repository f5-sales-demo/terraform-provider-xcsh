---
page_title: "items.get_spec.infra.hw_info.memory"
subcategory: ""
description: "Memory Information. Memory information."
xcsh_docs: {"aliases": ["items get spec infra hw info memory"], "body_bytes": 1610, "body_sha256": "sha256:4b15047e66a5f4c9ccbdcbd14b0f980137997aa05b44687b1e492da432159bdc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:memory", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/memory/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0320021223023322-2211221200033132-0223111232231321-3100100311210120-3122221233211010-3023100313001032-3332012213031313-1302111032131300", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "memory"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info memory size mb"], "anchor": "schema-items--get_spec--infra--hw_info--memory--size_mb", "description": "RAM. RAM size in MB.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory", "size_mb"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info memory speed"], "anchor": "schema-items--get_spec--infra--hw_info--memory--speed", "description": "Speed. RAM data rate in MT/s.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info memory type"], "anchor": "schema-items--get_spec--infra--hw_info--memory--type", "description": "Type. Type of memory, eg. DDR4.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/memory/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Memory Information. Memory information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.memory

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
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
