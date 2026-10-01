---
page_title: "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5477, "body_sha256": "sha256:0154c3999104c6637e4ee2a250ab3896b7acb305bfad75cf98ac2b4f5dfeb41d", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "documentation/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/)
- [eks_k8s.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/)
- [eks_k8s.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/)
- [eks_k8s.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

### configured_address property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/first_address/): complete subsection reference.

- [last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/last_address/): complete subsection reference.

## Next pages

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/first_address/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/last_address/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
