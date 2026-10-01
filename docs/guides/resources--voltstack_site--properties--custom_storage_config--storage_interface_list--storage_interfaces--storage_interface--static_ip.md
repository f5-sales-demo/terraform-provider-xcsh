---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3035, "body_sha256": "sha256:5be339c7771812153e64b790ebf54dba0782759d90cc59fc18d5728ad5f13f54", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:cluster_static_ip", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:node_static_ip"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--node_static_ip.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--cluster_static_ip.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--node_static_ip.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
