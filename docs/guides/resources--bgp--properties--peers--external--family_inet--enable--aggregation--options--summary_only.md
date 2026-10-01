---
page_title: "peers.external.family_inet.enable.aggregation.options.summary_only"
subcategory: ""
description: "peers.external.family_inet.enable.aggregation.options.summary_only for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1652, "body_sha256": "sha256:49d1344551c80d489efdecee3898d8e6c7b9a039b39c67d0ae9fd8bacd800997", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options:summary_only", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options:summary_only", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options", "path": "docs/guides/resources--bgp--properties--peers--external--family_inet--enable--aggregation--options--summary_only.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "options", "summary_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/summary_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.family_inet.enable.aggregation.options.summary_only for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable.aggregation.options.summary_only

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- [peers.external.family_inet](resources--bgp--properties--peers--external--family_inet.md)
- [peers.external.family_inet.enable](resources--bgp--properties--peers--external--family_inet--enable.md)
- [peers.external.family_inet.enable.aggregation](resources--bgp--properties--peers--external--family_inet--enable--aggregation.md)
- [peers.external.family_inet.enable.aggregation.options](resources--bgp--properties--peers--external--family_inet--enable--aggregation--options.md)
- peers.external.family_inet.enable.aggregation.options.summary_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for summary only.

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
summary_only {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers.external.family_inet.enable.aggregation.options](resources--bgp--properties--peers--external--family_inet--enable--aggregation--options.md)
- [xcsh_bgp](../resources/bgp.md)
