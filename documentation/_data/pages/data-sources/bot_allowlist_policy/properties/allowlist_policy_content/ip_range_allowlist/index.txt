---
page_title: "allowlist_policy_content.ip_range_allowlist"
subcategory: ""
description: "IP Range. Allowlist or permitted items"
xcsh_docs: {"aliases": ["allowlist policy content ip range allowlist"], "body_bytes": 1230, "body_sha256": "sha256:57d122c434dd8ad999cac9799f523f7591ed7985122e12c500a5d3f81649bd56", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0111122020022203-1232122233101210-1033333222323312-2133030202122102-3211130120212000-1221203332300202-1231210012013213-1211130121211230", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content", "ip_range_allowlist"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip range allowlist end with"], "anchor": "schema-allowlist_policy_content--ip_range_allowlist--end_with", "description": "End With. IP range end with.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist", "end_with"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowlist policy content ip range allowlist ip description"], "anchor": "schema-allowlist_policy_content--ip_range_allowlist--ip_description", "description": "Description. The description for IP range.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist", "ip_description"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowlist policy content ip range allowlist start with"], "anchor": "schema-allowlist_policy_content--ip_range_allowlist--start_with", "description": "Start With. IP range start with.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist", "start_with"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "IP Range. Allowlist or permitted items", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowlist_policy_content.ip_range_allowlist

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/)
- allowlist_policy_content.ip_range_allowlist

<a id="section"></a>

Type: `"list"`. Computed.

IP Range. Allowlist or permitted items

## Direct properties

<a id="schema-allowlist_policy_content--ip_range_allowlist--end_with"></a>

### end_with property

Type: `"string"`. Computed.

End With. IP range end with.

<a id="schema-allowlist_policy_content--ip_range_allowlist--ip_description"></a>

### ip_description property

Type: `"string"`. Computed.

Description. The description for IP range.

<a id="schema-allowlist_policy_content--ip_range_allowlist--start_with"></a>

### start_with property

Type: `"string"`. Computed.

Start With. IP range start with.
