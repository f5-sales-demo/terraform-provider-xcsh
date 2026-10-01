---
page_title: "proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip"
subcategory: ""
description: "proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3696, "body_sha256": "sha256:9c33d8f2b32c24eb914093aaeebb7cc611dd98a3ba209ed3301ff4e52577a7c8", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip:virtual_site"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "docs/guides/data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="section"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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

## Direct properties

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip"></a>

### ip property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network"></a>

### network property

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
