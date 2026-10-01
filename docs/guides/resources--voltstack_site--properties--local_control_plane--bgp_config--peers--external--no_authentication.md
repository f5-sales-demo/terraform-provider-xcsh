---
page_title: "local_control_plane.bgp_config.peers.external.no_authentication"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external.no_authentication for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1558, "body_sha256": "sha256:0d46283c465cee00a4538a70ca910a7870837c1928ecadb31e69e2fdc2c3ea62", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:no_authentication", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:no_authentication", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--no_authentication.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "no_authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/no_authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external.no_authentication for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.no_authentication

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](resources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- local_control_plane.bgp_config.peers.external.no_authentication

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authentication.

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
no_authentication = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
