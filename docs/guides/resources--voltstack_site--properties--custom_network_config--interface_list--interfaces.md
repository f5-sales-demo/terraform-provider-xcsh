---
page_title: "custom_network_config.interface_list.interfaces"
subcategory: ""
description: "custom_network_config.interface_list.interfaces for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5775, "body_sha256": "sha256:d9a53319fbe3e10d4d316ee27e1451164508789cb91d44137106417e5d62ee55", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:labels", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.interface_list.interfaces

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](resources--voltstack_site--properties--custom_network_config--interface_list.md)
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

- [dc_cluster_group_connectivity_interface_disabled](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_disabled.md): complete subsection reference.

- [dc_cluster_group_connectivity_interface_enabled](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_enabled.md): complete subsection reference.

- [dedicated_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md): complete subsection reference.

- [dedicated_management_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md): complete subsection reference.

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

- [ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md): complete subsection reference.

- [labels](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--labels.md): complete subsection reference.

- [tunnel_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_disabled.md)
- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_enabled.md)
- [custom_network_config.interface_list.interfaces.dedicated_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md)
- [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- [custom_network_config.interface_list.interfaces.labels](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--labels.md)
- [custom_network_config.interface_list.interfaces.tunnel_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface.md)
- [custom_network_config.interface_list](resources--voltstack_site--properties--custom_network_config--interface_list.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
