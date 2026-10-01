---
page_title: "local_ip.ip_address.ip_address.dual_stack.ipv4"
subcategory: ""
description: "local_ip.ip_address.ip_address.dual_stack.ipv4 for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 2315, "body_sha256": "sha256:b807ba510627b43003dccf6a8e0a4c3892cdbe18d91bb6e919631a88e4c54a21", "canonical_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "child_ids": [], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "path": "docs/guides/resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.ip_address.dual_stack.ipv4 for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address.dual_stack.ipv4

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [local_ip](resources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](resources--tunnel--properties--local_ip--ip_address.md)
- [local_ip.ip_address.ip_address](resources--tunnel--properties--local_ip--ip_address--ip_address.md)
- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md)
- local_ip.ip_address.ip_address.dual_stack.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_ip--ip_address--ip_address--dual_stack--ipv4--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md)
- [xcsh_tunnel](../resources/tunnel.md)
