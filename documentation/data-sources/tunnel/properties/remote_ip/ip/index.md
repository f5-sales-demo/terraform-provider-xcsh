---
page_title: "remote_ip.ip"
subcategory: ""
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["remote ip ip"], "body_bytes": 1910, "body_sha256": "sha256:e6c15bd1a0fdd5fd2ff625d3680e611f3ea6e652dd726655d570865ed1f99dfe", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv4", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv6"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "path": "documentation/data-sources/tunnel/properties/remote_ip/ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "ip"], "schema_version": 1, "sections": [{"aliases": ["remote ip ip dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip", "dual_stack"], "syntax": "attribute", "type": "object"}, {"aliases": ["remote ip ip ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["remote ip ip ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- remote_ip.ip

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

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/ipv6/): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/dual_stack/)
- [remote_ip.ip.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/ipv4/)
- [remote_ip.ip.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/ipv6/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
