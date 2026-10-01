---
page_title: "local_control_plane.bgp_config.peers.external.default_gateway_v6"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external.default_gateway_v6 for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1562, "body_sha256": "sha256:f9a59038375f6a44cae611b528b7ddacad0276196ecd56687b7584a20ce8637a", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:default_gateway_v6", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:default_gateway_v6", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--default_gateway_v6.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "default_gateway_v6"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/default_gateway_v6/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external.default_gateway_v6 for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.default_gateway_v6

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](resources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- local_control_plane.bgp_config.peers.external.default_gateway_v6

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway v6.

Upstream description:

This can be used for messages where no values are needed.

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
default_gateway_v6 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
