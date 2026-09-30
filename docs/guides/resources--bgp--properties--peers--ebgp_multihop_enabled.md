---
page_title: "peers.ebgp_multihop_enabled"
subcategory: ""
description: "peers.ebgp_multihop_enabled for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 800, "body_sha256": "sha256:99d56669e84a7ab70ef66ddffa409537394cd43f3c4c7f94ec8108166c6e9d9a", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_enabled", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_enabled", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "docs/guides/resources--bgp--properties--peers--ebgp_multihop_enabled.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "ebgp_multihop_enabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/ebgp_multihop_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.ebgp_multihop_enabled for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# peers.ebgp_multihop_enabled

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- peers.ebgp_multihop_enabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
ebgp_multihop_enabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers](resources--bgp--properties--peers.md)
- [xcsh_bgp](../resources/bgp.md)
