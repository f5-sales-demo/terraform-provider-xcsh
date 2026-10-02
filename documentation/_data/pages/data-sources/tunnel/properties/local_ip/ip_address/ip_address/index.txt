---
page_title: "local_ip.ip_address.ip_address"
subcategory: ""
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["local ip ip address ip address"], "body_bytes": 2259, "body_sha256": "sha256:f51702ca30970c41446596f2fc80f1cd9b07cb2f9878fbbca5f99eeb70e280f8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "path": "documentation/data-sources/tunnel/properties/local_ip/ip_address/ip_address/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address"], "schema_version": 1, "sections": [{"aliases": ["dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/)
- local_ip.ip_address.ip_address

<a id="section"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

## Direct properties

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/ipv6/): complete subsection reference.

## Next pages

- [local_ip.ip_address.ip_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/)
- [local_ip.ip_address.ip_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/ipv4/)
- [local_ip.ip_address.ip_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/ipv6/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
