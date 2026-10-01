---
page_title: "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router"
subcategory: ""
description: "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4046, "body_sha256": "sha256:b946693b66822cfa3a6f92a160e884179f0bb086b1e107492b348c459095472a", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config", "path": "docs/guides/resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [openstack](resources--securemesh_site_v2--properties--openstack.md)
- [openstack.not_managed](resources--securemesh_site_v2--properties--openstack--not_managed.md)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--properties--openstack--not_managed--node_list.md)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config.md)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

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

- [dns_config](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md): complete subsection reference.

<a id="schema-openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix"></a>

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

- [stateful](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md): complete subsection reference.

## Next pages

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
