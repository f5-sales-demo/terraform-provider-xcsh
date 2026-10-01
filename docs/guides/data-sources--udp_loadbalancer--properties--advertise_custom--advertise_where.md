---
page_title: "advertise_custom.advertise_where"
subcategory: ""
description: "advertise_custom.advertise_where for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6957, "body_sha256": "sha256:b1f68ea540fbb863c66884c3ab80d70b877f86ddbe62acfbe788f8e99836b761", "canonical_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:site", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service"], "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom", "path": "docs/guides/data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
- [Property reference](data-sources--udp_loadbalancer--reference.md)
- [advertise_custom](data-sources--udp_loadbalancer--properties--advertise_custom.md)
- advertise_custom.advertise_where

<a id="section"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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

## Direct properties

- [advertise_dualstack_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public.md): complete subsection reference.

- [advertise_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public.md): complete subsection reference.

- [advertise_v6_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md): complete subsection reference.

<a id="schema-advertise_custom--advertise_where--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

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

- [site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md): complete subsection reference.

- [use_default_port](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md): complete subsection reference.

- [virtual_network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md): complete subsection reference.

- [virtual_site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md): complete subsection reference.

- [virtual_site_with_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md): complete subsection reference.

- [vk8s_service](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public.md)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public.md)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md)
- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md)
- [advertise_custom.advertise_where.use_default_port](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md)
- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md)
- [advertise_custom](data-sources--udp_loadbalancer--properties--advertise_custom.md)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
