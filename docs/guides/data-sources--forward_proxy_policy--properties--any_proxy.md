---
page_title: "any_proxy"
subcategory: "Security"
description: "any_proxy for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1450, "body_sha256": "sha256:fca4cbf5b2570dd08df7df2efb4f64884e35bf84ad098956794af816eeff4f18", "canonical_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:any_proxy", "child_ids": [], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:any_proxy", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:reference", "path": "docs/guides/data-sources--forward_proxy_policy--properties--any_proxy.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["any_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/any_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_proxy for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_proxy

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- any_proxy

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_proxy, drp\_http\_connect, network\_connector, proxy\_label\_selector\] Enable this
option

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

- [any_proxy](data-sources--forward_proxy_policy--properties--any_proxy.md#section)
- [drp_http_connect](data-sources--forward_proxy_policy--properties--drp_http_connect.md#section)
- [network_connector](data-sources--forward_proxy_policy--properties--network_connector.md#section)
- [proxy_label_selector](data-sources--forward_proxy_policy--properties--proxy_label_selector.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--forward_proxy_policy--reference.md)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
