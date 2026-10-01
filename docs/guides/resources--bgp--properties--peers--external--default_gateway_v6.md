---
page_title: "peers.external.default_gateway_v6"
subcategory: ""
description: "peers.external.default_gateway_v6 for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:3976268a5c6bb9228f52d8fccb927eaa15c3f589174b6c342870a3703a564e6d", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "docs/guides/resources--bgp--properties--peers--external--default_gateway_v6.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "default_gateway_v6"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/default_gateway_v6/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.default_gateway_v6 for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.default_gateway_v6

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- peers.external.default_gateway_v6

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway v6.

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
default_gateway_v6 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers.external](resources--bgp--properties--peers--external.md)
- [xcsh_bgp](../resources/bgp.md)
