---
page_title: "ethernet_interface.ipv6_auto_config.router.dns_config.configured_list"
subcategory: ""
description: "IPV6DnsList."
xcsh_docs: {"aliases": ["ethernet interface ipv6 auto config router dns config configured list"], "body_bytes": 3532, "body_sha256": "sha256:6a7415d740b19d0fe175b778f0fbffaad0be1298b1bd5fa0ac83c30b5992943b", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "path": "documentation/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0231301112223203-3321303100222022-0010023322110221-1023200010001230-3031001110033330-0033213121333032-3333200331032100-2012132203112220", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.configured_list:RequiredObjectAttributes:dns_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config", "configured_list"], "schema_version": 1, "sections": [{"aliases": ["ethernet interface ipv6 auto config router dns config configured list dns list"], "anchor": "schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list", "description": "List of IPv6 Addresses acting as DNS servers.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config", "configured_list", "dns_list"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPV6DnsList.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/)
- [ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/)
- [ethernet_interface.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/)
- ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
```

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
configured_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list"></a>

### dns_list property

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
