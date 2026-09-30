---
page_title: "peers.external.default_gateway"
subcategory: ""
description: "peers.external.default_gateway for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 872, "body_sha256": "sha256:511f8ac71c2f6581e8bece1400bb55afa9efb926106bd9a7560490bc7f6c2d61", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:external:default_gateway", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:default_gateway", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "docs/guides/data-sources--bgp--properties--peers--external--default_gateway.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "default_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/default_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.default_gateway for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# peers.external.default_gateway

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.external](data-sources--bgp--properties--peers--external.md)
- peers.external.default_gateway

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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
