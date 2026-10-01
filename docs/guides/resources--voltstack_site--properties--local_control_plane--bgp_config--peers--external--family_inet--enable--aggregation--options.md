---
page_title: "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3072, "body_sha256": "sha256:676815b2ad20b0f0e720be12529a00bcee50a4b1510ff3ff9727ad2405440d24", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options:summary_only"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--options.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation", "options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](resources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.external](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external.md)
- [local_control_plane.bgp_config.peers.external.family_inet](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation.md)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [summary_only](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--options--summary_only.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--options--summary_only.md)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
