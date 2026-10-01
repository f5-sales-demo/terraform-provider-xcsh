---
page_title: "proxy_advertisement.advertise_custom.advertise_where.virtual_network"
subcategory: ""
description: "proxy_advertisement.advertise_custom.advertise_where.virtual_network for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 5335, "body_sha256": "sha256:7ec88b482ccbb3c604db6a7d51e123914bbe56878d8c4d588063b3fe1a10b3ad", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network:default_v6_vip", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network:default_vip", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network:virtual_network"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom.advertise_where.virtual_network for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.virtual_network

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

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

- [default_v6_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md): complete subsection reference.

- [default_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md): complete subsection reference.

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip"></a>

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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_vip"></a>

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

- [virtual_network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
