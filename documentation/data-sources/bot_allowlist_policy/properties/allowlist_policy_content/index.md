---
page_title: "allowlist_policy_content"
subcategory: ""
description: "IP Allowlist. Allowlist Policy Content."
xcsh_docs: {"aliases": ["allowlist policy content"], "body_bytes": 924, "body_sha256": "sha256:f6d094f2e6dc0d4905299d5d9be51bfa3fb7b7853bb73a087c1a1e601084c12b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "path": "documentation/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0131022110231201-0200202321231133-2200222130022320-2033331112100012-0212011010313101-0123010203230113-2210012232233302-1302000003301102", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowlist_policy_content"], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content ip allowlist"], "anchor": "section", "description": "IP & IP Prefix. Allowlist or permitted items", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_allowlist", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allowlist_policy_content", "ip_allowlist"], "syntax": "attribute", "type": "object"}, {"aliases": ["allowlist policy content ip range allowlist"], "anchor": "section", "description": "IP Range. Allowlist or permitted items", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content:ip_range_allowlist", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allowlist_policy_content", "ip_range_allowlist"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "IP Allowlist. Allowlist Policy Content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
