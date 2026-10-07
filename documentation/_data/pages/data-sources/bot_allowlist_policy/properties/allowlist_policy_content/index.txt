---
page_title: "allowlist_policy_content"
subcategory: ""
description: "IP Allowlist. Allowlist Policy Content."
xcsh_docs: {"aliases": ["allowlist policy content"], "body_bytes": 924, "body_sha256": "sha256:f6d094f2e6dc0d4905299d5d9be51bfa3fb7b7853bb73a087c1a1e601084c12b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0131022110231201-0200202321231133-2200222130022320-2033331112100012-0212011010313101-0123010203230113-2210012232233302-1302000003301102", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip allowlist"], "anchor": "section", "description": "IP & IP Prefix. Allowlist or permitted items", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist"], "syntax": "attribute", "type": "object"}, {"aliases": ["allowlist policy content ip range allowlist"], "anchor": "section", "description": "IP Range. Allowlist or permitted items", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "IP Allowlist. Allowlist Policy Content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowlist_policy_content

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- allowlist_policy_content

<a id="section"></a>

Type: `"single"`. Computed.

IP Allowlist. Allowlist Policy Content.

## Direct properties

- [ip_allowlist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/): complete subsection reference.

- [ip_range_allowlist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/): complete subsection reference.
