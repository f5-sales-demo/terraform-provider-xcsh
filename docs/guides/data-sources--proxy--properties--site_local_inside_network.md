---
page_title: "site_local_inside_network"
subcategory: ""
description: "site_local_inside_network for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1069, "body_sha256": "sha256:cd3dc13e42d1ce398b79bd648e2b633f64ae4e870e5c3e00bdc13e2595c67aad", "canonical_id": "xcsh-docs:data-sources:proxy:properties:site_local_inside_network", "child_ids": [], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:site_local_inside_network", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "docs/guides/data-sources--proxy--properties--site_local_inside_network.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_local_inside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/site_local_inside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_local_inside_network for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_local_inside_network

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- site_local_inside_network

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: site\_local\_inside\_network, site\_local\_network\] Enable this option

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

OneOf alternatives in this subsection:

- [site_local_inside_network](data-sources--proxy--properties--site_local_inside_network.md#section)
- [site_local_network](data-sources--proxy--properties--site_local_network.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--proxy--reference.md)
- [xcsh_proxy](../data-sources/proxy.md)
