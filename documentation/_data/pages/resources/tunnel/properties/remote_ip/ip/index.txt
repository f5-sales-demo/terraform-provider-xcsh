---
page_title: "remote_ip.ip"
subcategory: ""
description: "remote_ip.ip for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 2183, "body_sha256": "sha256:222bb82089141bd662d8f273b4bcf707a07768761a98f1f09b0d2badb57316e9", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv4", "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv6"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip", "path": "documentation/resources/tunnel/properties/remote_ip/ip/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["remote_ip", "ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# remote_ip.ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- remote_ip.ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
```

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

Terraform syntax:

```terraform
ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv6/): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/)
- [remote_ip.ip.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv4/)
- [remote_ip.ip.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/ipv6/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
