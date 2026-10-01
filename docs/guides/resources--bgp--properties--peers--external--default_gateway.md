---
page_title: "peers.external.default_gateway"
subcategory: ""
description: "peers.external.default_gateway for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1011, "body_sha256": "sha256:a934717080936b0da60ff094ddc376a4a014a66ac0e7baecc348cf3a8323c77a", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "docs/guides/resources--bgp--properties--peers--external--default_gateway.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "default_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/default_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.default_gateway for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.default_gateway

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- peers.external.default_gateway

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_gateway = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers.external](resources--bgp--properties--peers--external.md)
- [xcsh_bgp](../resources/bgp.md)
