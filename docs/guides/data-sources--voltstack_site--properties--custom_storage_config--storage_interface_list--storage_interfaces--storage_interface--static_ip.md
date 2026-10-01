---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2772, "body_sha256": "sha256:ea47fde842955da30c264de65994c637a8ea02bb665d5df40570732c0ded5b42", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:cluster_static_ip", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:node_static_ip"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

## Direct properties

- [cluster_static_ip](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--node_static_ip.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--cluster_static_ip.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--node_static_ip.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
