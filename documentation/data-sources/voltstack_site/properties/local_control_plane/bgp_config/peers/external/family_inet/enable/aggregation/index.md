---
page_title: "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation"
subcategory: ""
description: "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements."
xcsh_docs: {"aliases": ["local control plane bgp config peers external family inet enable aggregation"], "body_bytes": 4272, "body_sha256": "sha256:b4cdfb6cafb8e3bf99bc5d7f635e22ae5a6d7ef14bf3557c691b5e44355e6c2a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable", "path": "documentation/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation"], "schema_version": 1, "sections": [{"aliases": ["local control plane bgp config peers external family inet enable aggregation ip prefix"], "anchor": "schema-local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--ip_prefix", "description": "Specify IPv4 subnet for aggregation.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["local control plane bgp config peers external family inet enable aggregation options"], "anchor": "section", "description": "Configuration parameter for options", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation", "options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- [local_control_plane.bgp_config.peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--ip_prefix"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/options/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/options/)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
