---
page_title: "proxy_advertisement.advertise_custom.advertise_where.virtual_network"
subcategory: ""
description: "proxy_advertisement.advertise_custom.advertise_where.virtual_network for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4551, "body_sha256": "sha256:788426ecb57dde8dd37a73b40111cedf35f41400e15d4626b29f5ef756355090", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network:default_v6_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network:default_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network:virtual_network"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "docs/guides/data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom.advertise_where.virtual_network for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.virtual_network

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="section"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

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

## Direct properties

- [default_v6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md): complete subsection reference.

- [default_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md): complete subsection reference.

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip"></a>

### specific_v6_vip property

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

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

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

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

- [virtual_network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
