---
page_title: "aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2897, "body_sha256": "sha256:73689cbe8cdc778fdeb7ac055f58b25ee2c0916d2c7b76ee1f305279efa87bbe", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "docs/guides/resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [aws](resources--securemesh_site_v2--properties--aws.md)
- [aws.not_managed](resources--securemesh_site_v2--properties--aws--not_managed.md)
- [aws.not_managed.node_list](resources--securemesh_site_v2--properties--aws--not_managed--node_list.md)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list.md)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [configured_list](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md): complete subsection reference.

## Next pages

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
