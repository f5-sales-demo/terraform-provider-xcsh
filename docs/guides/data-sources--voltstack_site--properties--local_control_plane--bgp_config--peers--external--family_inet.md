---
page_title: "local_control_plane.bgp_config.peers.external.family_inet"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external.family_inet for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2175, "body_sha256": "sha256:8d6e16818c4da8400ef66967c86a74329fda90559eb8d052cc94049a45983889", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:disable_spec", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external.family_inet for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- local_control_plane.bgp_config.peers.external.family_inet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

## Direct properties

- [disable_spec](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--disable_spec.md): complete subsection reference.

- [enable](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.disable_spec](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--disable_spec.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable.md)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
