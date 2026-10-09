---
page_title: "allowlist_policy_content.ip_range_allowlist"
subcategory: ""
description: "IP Range. Allowlist or permitted items"
xcsh_docs: {"aliases": ["allowlist policy content ip range allowlist"], "body_bytes": 1230, "body_sha256": "sha256:57d122c434dd8ad999cac9799f523f7591ed7985122e12c500a5d3f81649bd56", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0111122020022203-1232122233101210-1033333222323312-2133030202122102-3211130120212000-1221203332300202-1231210012013213-1211130121211230", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content", "ip_range_allowlist"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip range allowlist end with"], "anchor": "schema-allowlist_policy_content--ip_range_allowlist--end_with", "description": "End With. IP range end with.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist", "end_with"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowlist policy content ip range allowlist ip description"], "anchor": "schema-allowlist_policy_content--ip_range_allowlist--ip_description", "description": "Description. The description for IP range.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist", "ip_description"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowlist policy content ip range allowlist start with"], "anchor": "schema-allowlist_policy_content--ip_range_allowlist--start_with", "description": "Start With. IP range start with.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist", "start_with"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "IP Range. Allowlist or permitted items", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
