---
page_title: "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router"
subcategory: ""
description: "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4004, "body_sha256": "sha256:1c284a2bc910d785c8e10bff6c4e0dab53025fc4b7a42b1622e38e58de8bcb0b", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config", "path": "docs/guides/resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [eks_k8s](resources--securemesh_site_v2--properties--eks_k8s.md)
- [eks_k8s.not_managed](resources--securemesh_site_v2--properties--eks_k8s--not_managed.md)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list.md)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list.md)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config.md)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router

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

- [dns_config](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md): complete subsection reference.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix"></a>

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

- [stateful](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md): complete subsection reference.

## Next pages

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
