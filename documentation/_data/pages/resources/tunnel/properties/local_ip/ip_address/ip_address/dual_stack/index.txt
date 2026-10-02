---
page_title: "local_ip.ip_address.ip_address.dual_stack"
subcategory: ""
description: "DualStackAddressType represents both IPv4 and IPv6 together."
xcsh_docs: {"aliases": ["local ip ip address ip address dual stack"], "body_bytes": 2198, "body_sha256": "sha256:497b3c289a88230d620b1942488cabd0c87c99ab0fa9fc54fa916f27c16a9814", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3013212302112132-1030133231022030-3222012011111002-1112033112113322-3030322233312100-1312312201211031-1020310211101203-1321022203113030", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack", "ipv4"], "syntax": "block", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv6", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack", "ipv6"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "DualStackAddressType represents both IPv4 and IPv6 together.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address.dual_stack

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- [local_ip.ip_address.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/)
- local_ip.ip_address.ip_address.dual_stack

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv6/): complete subsection reference.

## Next pages

- [local_ip.ip_address.ip_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv4/)
- [local_ip.ip_address.ip_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv6/)
- [local_ip.ip_address.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
