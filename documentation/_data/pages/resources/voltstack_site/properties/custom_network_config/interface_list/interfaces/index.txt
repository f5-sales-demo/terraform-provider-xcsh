---
page_title: "custom_network_config.interface_list.interfaces"
subcategory: ""
description: "custom_network_config.interface_list.interfaces for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 6852, "body_sha256": "sha256:d8479383bbfaf2d0d467982e7172f642098ae0716069933a6e3bcc1c5f1fdad4", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:labels", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list", "path": "documentation/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/)
- custom_network_config.interface_list.interfaces

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure network interfaces for this App Stack site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dc_cluster_group_connectivity_interface_disabled",
    "dc_cluster_group_connectivity_interface_enabled"),
  validators.ConflictingListObjectAttributes("dedicated_interface",
    "dedicated_management_interface"),
  validators.ConflictingListObjectAttributes("dedicated_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("dedicated_interface",
    "tunnel_interface"),
  validators.ConflictingListObjectAttributes("dedicated_management_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("dedicated_management_interface",
    "tunnel_interface"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "tunnel_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/): complete subsection reference.

- [dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/): complete subsection reference.

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/): complete subsection reference.

- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/labels/): complete subsection reference.

- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/)
- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/)
- [custom_network_config.interface_list.interfaces.dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/)
- [custom_network_config.interface_list.interfaces.dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/labels/)
- [custom_network_config.interface_list.interfaces.tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
