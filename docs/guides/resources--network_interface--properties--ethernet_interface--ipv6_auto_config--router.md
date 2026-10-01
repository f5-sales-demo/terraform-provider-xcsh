---
page_title: "ethernet_interface.ipv6_auto_config.router"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config.router for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 3306, "body_sha256": "sha256:e77253ec9077a856a9208ef42c8f850be8a7626d1f1f4dd5804681d87cf57b7a", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config.router for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- ethernet_interface.ipv6_auto_config.router

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dns_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md): complete subsection reference.

<a id="schema-ethernet_interface--ipv6_auto_config--router--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- [xcsh_network_interface](../resources/network_interface.md)
