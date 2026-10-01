---
page_title: "local_control_plane.bgp_config.peers.external.interface_list"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external.interface_list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2017, "body_sha256": "sha256:c8cdb7584c153280832c7053c84cd1bd309e16f80db1e9a72ad7efd62780a0c2", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface_list", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface_list:interfaces"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--interface_list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external.interface_list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.interface_list

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](resources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- local_control_plane.bgp_config.peers.external.interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--interface_list--interfaces.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.interface_list.interfaces](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--interface_list--interfaces.md)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
