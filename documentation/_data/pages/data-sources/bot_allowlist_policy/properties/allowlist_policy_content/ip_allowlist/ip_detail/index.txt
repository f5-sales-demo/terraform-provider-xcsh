---
page_title: "allowlist_policy_content.ip_allowlist.ip_detail"
subcategory: ""
description: "IP Detail. Support the single IP value."
xcsh_docs: {"aliases": ["allowlist policy content ip allowlist ip detail"], "body_bytes": 1269, "body_sha256": "sha256:2fb8eff9ffdb1dc93e9baeb3e7b245b82eeabc5b574bfa1cce1c134b74220677", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_detail", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_detail/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-1313123213301330-1200212333303012-3222311022203302-3313103113203012-3020311211331023-0033233200133231-0010010331103133-2102110310303332", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_detail"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip allowlist ip detail ip description"], "anchor": "schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_description", "description": "Description. The description for IP address.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_detail", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_detail", "ip_description"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowlist policy content ip allowlist ip detail ip value"], "anchor": "schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_value", "description": "Value. A single IP address.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_detail", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_detail", "ip_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_detail/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "IP Detail. Support the single IP value.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowlist_policy_content.ip_allowlist.ip_detail

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/)
- [allowlist_policy_content.ip_allowlist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/)
- allowlist_policy_content.ip_allowlist.ip_detail

<a id="section"></a>

Type: `"single"`. Computed.

IP Detail. Support the single IP value.

## Direct properties

<a id="schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_description"></a>

### ip_description property

Type: `"string"`. Computed.

Description. The description for IP address.

<a id="schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_value"></a>

### ip_value property

Type: `"string"`. Computed.

Value. A single IP address.
