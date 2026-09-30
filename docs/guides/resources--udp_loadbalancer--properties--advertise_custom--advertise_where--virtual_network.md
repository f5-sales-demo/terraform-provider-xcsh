---
page_title: "advertise_custom.advertise_where.virtual_network"
subcategory: ""
description: "advertise_custom.advertise_where.virtual_network for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4756, "body_sha256": "sha256:515598fdf56672687dbe78792d9c737e33126bb89866beb8bb0e884e743a1f45", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_v6_vip", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_vip", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:virtual_network"], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "parent_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where", "path": "docs/guides/resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.virtual_network for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advertise_custom.advertise_where.virtual_network

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- [advertise_custom](resources--udp_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- advertise_custom.advertise_where.virtual_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_v6_vip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_v6_vip.md): complete subsection reference.

- [default_vip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_vip.md): complete subsection reference.

<a id="schema-advertise_custom--advertise_where--virtual_network--specific_v6_vip"></a>

### specific_v6_vip property

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-advertise_custom--advertise_where--virtual_network--specific_vip"></a>

### specific_vip property

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

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

- [virtual_network](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_v6_vip.md)
- [advertise_custom.advertise_where.virtual_network.default_vip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_vip.md)
- [advertise_custom.advertise_where.virtual_network.virtual_network](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
