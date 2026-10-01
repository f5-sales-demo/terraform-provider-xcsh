---
page_title: "peers.passive_mode_enabled"
subcategory: ""
description: "peers.passive_mode_enabled for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 896, "body_sha256": "sha256:69c31652a90b6fa97335c779e637d7c84cfed1c5e9712fb3a517d6963698995e", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:passive_mode_enabled", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:passive_mode_enabled", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "docs/guides/resources--bgp--properties--peers--passive_mode_enabled.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "passive_mode_enabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/passive_mode_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.passive_mode_enabled for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.passive_mode_enabled

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- peers.passive_mode_enabled

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
passive_mode_enabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers](resources--bgp--properties--peers.md)
- [xcsh_bgp](../resources/bgp.md)
