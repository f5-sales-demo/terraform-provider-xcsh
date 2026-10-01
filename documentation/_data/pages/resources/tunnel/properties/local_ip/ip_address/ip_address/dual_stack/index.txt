---
page_title: "local_ip.ip_address.ip_address.dual_stack"
subcategory: ""
description: "local_ip.ip_address.ip_address.dual_stack for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 2198, "body_sha256": "sha256:497b3c289a88230d620b1942488cabd0c87c99ab0fa9fc54fa916f27c16a9814", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv6"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.ip_address.dual_stack for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
