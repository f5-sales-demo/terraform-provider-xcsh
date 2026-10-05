---
page_title: "remote_ip.ip.dual_stack"
subcategory: ""
description: "DualStackAddressType represents both IPv4 and IPv6 together."
xcsh_docs: {"aliases": ["remote ip ip dual stack"], "body_bytes": 1854, "body_sha256": "sha256:79bac51dfac1a47e1f97aa2d0b667b311c5690039bfd58add346413e0a3dadaa", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv6"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "path": "documentation/resources/tunnel/properties/remote_ip/ip/dual_stack/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0213230203000003-2112030131320032-2223233132332012-0102130123133003-3213101001120131-1112202321122013-3122131200031223-2130210323322303", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "ip", "dual_stack"], "schema_version": 1, "sections": [{"aliases": ["remote ip ip dual stack ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip", "dual_stack", "ipv4"], "syntax": "block", "type": "object"}, {"aliases": ["remote ip ip dual stack ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv6", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip", "dual_stack", "ipv6"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/ip/dual_stack/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "DualStackAddressType represents both IPv4 and IPv6 together.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip.dual_stack

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/)
- remote_ip.ip.dual_stack

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

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv6/): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/)
- [remote_ip.ip.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv6/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
