---
page_title: "peers.external.no_authentication"
subcategory: ""
description: "peers.external.no_authentication for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 977, "body_sha256": "sha256:83da9d03e55cc61956dddd80b83cc32db180650037579a293fa4aa160425fa29", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:external:no_authentication", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:no_authentication", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "docs/guides/data-sources--bgp--properties--peers--external--no_authentication.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "no_authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/no_authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.no_authentication for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.no_authentication

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- peers.external.no_authentication

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers.external](data-sources--bgp--properties--peers--external.md)
- [xcsh_bgp](../data-sources/bgp.md)
