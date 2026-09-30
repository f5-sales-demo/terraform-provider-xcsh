---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 3150, "body_sha256": "sha256:49f97b65a0e4e257b816258dc3bbd9b7ba753fb1952210ff568fa746e3a67361", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:host", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "documentation/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/index.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/host/): complete subsection reference.

- [router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/host/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
