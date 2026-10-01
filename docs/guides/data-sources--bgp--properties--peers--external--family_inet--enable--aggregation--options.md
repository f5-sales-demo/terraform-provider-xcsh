---
page_title: "peers.external.family_inet.enable.aggregation.options"
subcategory: ""
description: "peers.external.family_inet.enable.aggregation.options for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 2106, "body_sha256": "sha256:e1c445c13d409bdd0d9b8eacfff36dd7f4025aed80e3bfcb05135dda0d7179cd", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation:options", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation:options:summary_only"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation:options", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "path": "docs/guides/data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.family_inet.enable.aggregation.options for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable.aggregation.options

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- [peers.external.family_inet](data-sources--bgp--properties--peers--external--family_inet.md)
- [peers.external.family_inet.enable](data-sources--bgp--properties--peers--external--family_inet--enable.md)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation.md)
- peers.external.family_inet.enable.aggregation.options

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

- [summary_only](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options--summary_only.md): complete subsection reference.

## Next pages

- [peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options--summary_only.md)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation.md)
- [xcsh_bgp](../data-sources/bgp.md)
