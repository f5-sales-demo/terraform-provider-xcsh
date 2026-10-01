---
page_title: "proxy_advertisement.advertise_custom.advertise_where"
subcategory: ""
description: "proxy_advertisement.advertise_custom.advertise_where for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 10811, "body_sha256": "sha256:d09e69dd63d4749364f33af118fd35a54aa18c41adf26783ef56f293e1b4e814", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_on_public", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:site", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:use_default_port", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "parent_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom", "path": "docs/guides/resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom.advertise_where for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--properties--proxy_advertisement--advertise_custom.md)
- proxy_advertisement.advertise_custom.advertise_where

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

- [advertise_dualstack_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public.md): complete subsection reference.

- [advertise_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public.md): complete subsection reference.

- [advertise_v6_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public.md): complete subsection reference.

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--port"></a>

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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges"></a>

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

- [site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md): complete subsection reference.

- [use_default_port](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--use_default_port.md): complete subsection reference.

- [virtual_network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md): complete subsection reference.

- [virtual_site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md): complete subsection reference.

- [virtual_site_with_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md): complete subsection reference.

- [vk8s_service](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public.md)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public.md)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public.md)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--use_default_port.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service.md)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--properties--proxy_advertisement--advertise_custom.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
