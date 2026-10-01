---
page_title: "tunnel_interface.static_ip"
subcategory: ""
description: "tunnel_interface.static_ip for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1592, "body_sha256": "sha256:0705c1d1b8d7002707f038453dc73df353205a31f090cbf266ec9892dfdce91f", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:node_static_ip"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip", "parent_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface", "path": "docs/guides/data-sources--network_interface--properties--tunnel_interface--static_ip.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tunnel_interface", "static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/tunnel_interface/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tunnel_interface.static_ip for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.static_ip

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [tunnel_interface](data-sources--network_interface--properties--tunnel_interface.md)
- tunnel_interface.static_ip

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

- [cluster_static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md): complete subsection reference.

## Next pages

- [tunnel_interface.static_ip.cluster_static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip--cluster_static_ip.md)
- [tunnel_interface.static_ip.node_static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md)
- [tunnel_interface](data-sources--network_interface--properties--tunnel_interface.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
