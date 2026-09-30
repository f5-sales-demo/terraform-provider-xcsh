---
page_title: "advertise_custom.advertise_where"
subcategory: ""
description: "advertise_custom.advertise_where for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 10083, "body_sha256": "sha256:794834acce74c670c3c6bd35f11cabbd0c9c74aa6048b81bf33d0b7652b88d2b", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:site", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service"], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where", "parent_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom", "path": "docs/guides/resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- [advertise_custom](resources--udp_loadbalancer--properties--advertise_custom.md)
- advertise_custom.advertise_where

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_dualstack_on_public](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public.md): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public.md): complete subsection reference.

- [advertise_v6_on_public](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md): complete subsection reference.

<a id="schema-advertise_custom--advertise_where--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-advertise_custom--advertise_where--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md): complete subsection reference.

- [use_default_port](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md): complete subsection reference.

- [virtual_network](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md): complete subsection reference.

- [virtual_site](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md): complete subsection reference.

- [virtual_site_with_vip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md): complete subsection reference.

- [vk8s_service](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public.md)
- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public.md)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md)
- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md)
- [advertise_custom.advertise_where.use_default_port](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md)
- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md)
- [advertise_custom](resources--udp_loadbalancer--properties--advertise_custom.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
