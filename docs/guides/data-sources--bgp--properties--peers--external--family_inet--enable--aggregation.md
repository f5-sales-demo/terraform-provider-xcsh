---
page_title: "peers.external.family_inet.enable.aggregation"
subcategory: ""
description: "peers.external.family_inet.enable.aggregation for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 2845, "body_sha256": "sha256:04c1dd34bb38bbe678be3db025a62699d3e7c2086514e0c9305e88237cbeeaed", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation:options"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable", "path": "docs/guides/data-sources--bgp--properties--peers--external--family_inet--enable--aggregation.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable", "aggregation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.family_inet.enable.aggregation for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable.aggregation

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- [peers.external.family_inet](data-sources--bgp--properties--peers--external--family_inet.md)
- [peers.external.family_inet.enable](data-sources--bgp--properties--peers--external--family_inet--enable.md)
- peers.external.family_inet.enable.aggregation

<a id="section"></a>

Type: `"list"`. Computed.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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

<a id="schema-peers--external--family_inet--enable--aggregation--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

- [options](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options.md): complete subsection reference.

## Next pages

- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options.md)
- [peers.external.family_inet.enable](data-sources--bgp--properties--peers--external--family_inet--enable.md)
- [xcsh_bgp](../data-sources/bgp.md)
