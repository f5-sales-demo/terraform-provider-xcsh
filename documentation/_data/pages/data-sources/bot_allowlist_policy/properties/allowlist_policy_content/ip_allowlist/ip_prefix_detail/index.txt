---
page_title: "allowlist_policy_content.ip_allowlist.ip_prefix_detail"
subcategory: ""
description: "IP Prefix Detail. Support the IP prefix value."
xcsh_docs: {"aliases": ["allowlist policy content ip allowlist ip prefix detail"], "body_bytes": 1311, "body_sha256": "sha256:312f340a1c4232d0c225d76f6399c067fd35384dabcffd34fee4f3658b8e34df", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_prefix_detail", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_prefix_detail/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3233212033332213-1100132320320220-0102103020220003-3111330022230012-0222333033121331-0132310130222122-1213212000230113-3203220101030103", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_prefix_detail"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip allowlist ip prefix detail ip description"], "anchor": "schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_description", "description": "Description. The description for IP prefix.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_prefix_detail", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_prefix_detail", "ip_description"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowlist policy content ip allowlist ip prefix detail ip value"], "anchor": "schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_value", "description": "Value. IP prefix e.g. 192.0.2.0/24.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_prefix_detail", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_prefix_detail", "ip_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_prefix_detail/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "IP Prefix Detail. Support the IP prefix value.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowlist_policy_content.ip_allowlist.ip_prefix_detail

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/)
- [allowlist_policy_content.ip_allowlist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/)
- allowlist_policy_content.ip_allowlist.ip_prefix_detail

<a id="section"></a>

Type: `"single"`. Computed.

IP Prefix Detail. Support the IP prefix value.

## Direct properties

<a id="schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_description"></a>

### ip_description property

Type: `"string"`. Computed.

Description. The description for IP prefix.

<a id="schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_value"></a>

### ip_value property

Type: `"string"`. Computed.

Value. IP prefix e.g. 192.0.2.0/24.
