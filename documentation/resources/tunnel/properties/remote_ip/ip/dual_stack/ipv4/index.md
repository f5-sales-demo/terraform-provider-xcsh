---
page_title: "remote_ip.ip.dual_stack.ipv4"
subcategory: ""
description: "remote_ip.ip.dual_stack.ipv4 for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 2326, "body_sha256": "sha256:b80134798415db8d8c13548cb908b4854efd1798feed572b588f9e5d7f653a88", "child_ids": [], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "path": "documentation/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["remote_ip", "ip", "dual_stack", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/ip/dual_stack/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip.dual_stack.ipv4 for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# remote_ip.ip.dual_stack.ipv4

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/)
- [remote_ip.ip.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/)
- remote_ip.ip.dual_stack.ipv4

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

<a id="schema-remote_ip--ip--dual_stack--ipv4--addr"></a>

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

- [remote_ip.ip.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/ip/dual_stack/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
