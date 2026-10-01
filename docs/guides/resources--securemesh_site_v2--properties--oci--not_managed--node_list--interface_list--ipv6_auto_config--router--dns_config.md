---
page_title: "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2897, "body_sha256": "sha256:83f2d4ae571e732518dbe5617c16868292c1b092a2bc75c99f77c6ef52da63cf", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "docs/guides/resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [oci](resources--securemesh_site_v2--properties--oci.md)
- [oci.not_managed](resources--securemesh_site_v2--properties--oci--not_managed.md)
- [oci.not_managed.node_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list.md)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

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

- [configured_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md): complete subsection reference.

## Next pages

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
