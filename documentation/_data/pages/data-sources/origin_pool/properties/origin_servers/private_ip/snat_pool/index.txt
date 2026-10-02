---
page_title: "origin_servers.private_ip.snat_pool"
subcategory: "Load Balancing"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers private ip snat pool", "upstream servers"], "body_bytes": 2135, "body_sha256": "sha256:fae8fc2ec4cf02a230e451b61e18fa537850687e45424831e5c526bb67aaf877", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip", "path": "documentation/data-sources/origin_pool/properties/origin_servers/private_ip/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip:snat_pool:no_snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "private_ip", "snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip:snat_pool:snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_ip", "snat_pool", "snat_pool"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.private_ip.snat_pool

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/)
- [origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/)
- origin_servers.private_ip.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

## Direct properties

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/snat_pool/snat_pool/): complete subsection reference.

## Next pages

- [origin_servers.private_ip.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/snat_pool/no_snat_pool/)
- [origin_servers.private_ip.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/snat_pool/snat_pool/)
- [origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
