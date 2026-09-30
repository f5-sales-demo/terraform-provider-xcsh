---
page_title: "local_control_plane.bgp_config.peers"
subcategory: ""
description: "local_control_plane.bgp_config.peers for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5354, "body_sha256": "sha256:e12c639b6711ab9dbf0d65638a50689defa0e1fd98ed67221af22edcc0f6cd89", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_disabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:disable_spec", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_disabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_enabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:metadata", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_disabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_enabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_control_plane.bgp_config.peers

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- local_control_plane.bgp_config.peers

<a id="section"></a>

Type: `"list"`. Computed.

Peers. BGP parameters for peer.

Upstream description:

BGP parameters for peer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [bfd_disabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--bfd_disabled.md): complete subsection reference.

- [bfd_enabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--bfd_enabled.md): complete subsection reference.

- [disable_spec](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--disable_spec.md): complete subsection reference.

- [ebgp_multihop_disabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--ebgp_multihop_disabled.md): complete subsection reference.

- [ebgp_multihop_enabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--ebgp_multihop_enabled.md): complete subsection reference.

- [external](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md): complete subsection reference.

<a id="schema-local_control_plane--bgp_config--peers--label"></a>

### label property

Type: `"string"`. Computed.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--metadata.md): complete subsection reference.

- [passive_mode_disabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--passive_mode_disabled.md): complete subsection reference.

- [passive_mode_enabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--passive_mode_enabled.md): complete subsection reference.

- [routing_policies](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.bfd_disabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--bfd_disabled.md)
- [local_control_plane.bgp_config.peers.bfd_enabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--bfd_enabled.md)
- [local_control_plane.bgp_config.peers.disable_spec](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--disable_spec.md)
- [local_control_plane.bgp_config.peers.ebgp_multihop_disabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--ebgp_multihop_disabled.md)
- [local_control_plane.bgp_config.peers.ebgp_multihop_enabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--ebgp_multihop_enabled.md)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [local_control_plane.bgp_config.peers.metadata](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--metadata.md)
- [local_control_plane.bgp_config.peers.passive_mode_disabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--passive_mode_disabled.md)
- [local_control_plane.bgp_config.peers.passive_mode_enabled](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--passive_mode_enabled.md)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
