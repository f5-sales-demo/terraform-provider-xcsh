---
page_title: "tunnel_interface.static_ip"
subcategory: ""
description: "tunnel_interface.static_ip for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1864, "body_sha256": "sha256:6bd61519ba6cafad7cd957c1ae742c2e76925003690065c51bdf811e1370e717", "canonical_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip", "child_ids": ["xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:node_static_ip"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip", "parent_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "path": "docs/guides/resources--network_interface--properties--tunnel_interface--static_ip.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tunnel_interface", "static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/tunnel_interface/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tunnel_interface.static_ip for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.static_ip

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [tunnel_interface](resources--network_interface--properties--tunnel_interface.md)
- tunnel_interface.static_ip

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

- [cluster_static_ip](resources--network_interface--properties--tunnel_interface--static_ip--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](resources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md): complete subsection reference.

## Next pages

- [tunnel_interface.static_ip.cluster_static_ip](resources--network_interface--properties--tunnel_interface--static_ip--cluster_static_ip.md)
- [tunnel_interface.static_ip.node_static_ip](resources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md)
- [tunnel_interface](resources--network_interface--properties--tunnel_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
