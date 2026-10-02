---
page_title: "local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation"
subcategory: ""
description: "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements."
xcsh_docs: {"aliases": ["local control plane bgp config peers external family inet enable aggregation"], "body_bytes": 4356, "body_sha256": "sha256:7425ebe44db636a8029e9611063c3639e2855116c27cac014a5f3a3e8aa98250", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2312310113001021-3013203310001033-3323012223020012-1311122002233221-0332210111120323-2221031113100012-3031013131311100-1330212320013023", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation"], "schema_version": 1, "sections": [{"aliases": ["ip prefix"], "anchor": "schema-local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--ip_prefix", "description": "Specify IPv4 subnet for aggregation.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["options"], "anchor": "section", "description": "Configuration parameter for options", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation:options", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation", "options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- [local_control_plane.bgp_config.peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
aggregation {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_control_plane--bgp_config--peers--external--family_inet--enable--aggregation--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Optional.

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

- [options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/options/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/options/)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
