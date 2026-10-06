---
page_title: "allowlist_policy_content.ip_allowlist"
subcategory: ""
description: "IP & IP Prefix. Allowlist or permitted items"
xcsh_docs: {"aliases": ["allowlist policy content ip allowlist"], "body_bytes": 1130, "body_sha256": "sha256:810fb9dbf34e63202ae4965fc0e5c6c77e2dca4a8aa17ed8e582b5461a92c87e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_detail", "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_prefix_detail"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2202221230200011-3210122211122323-1200023002310231-2232102113323101-1011001131211012-1332002010332313-0232211332020233-0222132000002132", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content", "ip_allowlist"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip allowlist ip detail"], "anchor": "section", "description": "IP Detail. Support the single IP value.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_detail", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_detail"], "syntax": "attribute", "type": "object"}, {"aliases": ["allowlist policy content ip allowlist ip prefix detail"], "anchor": "section", "description": "IP Prefix Detail. Support the IP prefix value.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist:ip_prefix_detail", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist", "ip_prefix_detail"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "IP & IP Prefix. Allowlist or permitted items", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowlist_policy_content.ip_allowlist

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/)
- allowlist_policy_content.ip_allowlist

<a id="section"></a>

Type: `"list"`. Computed.

IP &amp; IP Prefix. Allowlist or permitted items

## Direct properties

- [ip_detail](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_detail/): complete subsection reference.

- [ip_prefix_detail](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_prefix_detail/): complete subsection reference.
