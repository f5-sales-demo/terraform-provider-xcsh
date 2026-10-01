---
page_title: "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2995, "body_sha256": "sha256:149cb43a0908bea0cc7c2d84390f6d03033fb15a680f428a5fc6c1b6ce6b1a2e", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options:summary_only"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--options.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation", "options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation.md)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

<a id="section"></a>

Type: `"list"`. Computed.

Aggregation OPTIONS. Configuration parameter for options

Upstream description:

Configuration parameter for options

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [summary_only](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--options--summary_only.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--options--summary_only.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
