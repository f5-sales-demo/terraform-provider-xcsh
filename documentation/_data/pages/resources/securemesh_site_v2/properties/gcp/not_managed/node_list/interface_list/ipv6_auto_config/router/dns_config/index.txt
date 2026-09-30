---
page_title: "gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3466, "body_sha256": "sha256:adcc5b67ef44ac34d88a9cb798fdea9476ac578f25781f5e2d9cd54525754319", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "documentation/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/)
- [gcp.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/)
- [gcp.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/)
- [gcp.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

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

- [configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/): complete subsection reference.

- [local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/): complete subsection reference.

## Next pages

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
