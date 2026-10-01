---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list"
subcategory: ""
description: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4092, "body_sha256": "sha256:6d1c5653025afa9e84d67d8c822d55fc0302e71fa91ae0873e3b3ddb96b3f0f1", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "docs/guides/resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [openshift_virtualization](resources--securemesh_site_v2--properties--openshift_virtualization.md)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed.md)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list.md)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

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

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list--dns_list"></a>

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

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
