---
page_title: "tunnel_interface"
subcategory: ""
description: "tunnel_interface for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 6479, "body_sha256": "sha256:c62faa8a82a8cc1f3a7a39725e65ca0b16143315aabd615f8f9a272397dcd754", "child_ids": ["xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_inside_network", "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_network", "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip", "xcsh-docs:resources:network_interface:properties:tunnel_interface:tunnel"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "parent_id": "xcsh-docs:resources:network_interface:reference", "path": "documentation/resources/network_interface/properties/tunnel_interface/index.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["tunnel_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/tunnel_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tunnel_interface for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- tunnel_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

Terraform syntax:

```terraform
tunnel_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tunnel_interface--mtu"></a>

### mtu property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="schema-tunnel_interface--node"></a>

### node property

Type: `"string"`. Optional.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-tunnel_interface--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_network/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/): complete subsection reference.

- [tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/tunnel/): complete subsection reference.

## Next pages

- [tunnel_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_inside_network/)
- [tunnel_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_network/)
- [tunnel_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/)
- [tunnel_interface.tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
