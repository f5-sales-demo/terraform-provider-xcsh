---
page_title: "any_proxy"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["any proxy"], "body_bytes": 1862, "body_sha256": "sha256:707f85a9915357e82811afad13686f949cc6996ad3366085c266dc46ef5e7801", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:any_proxy", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:reference", "path": "documentation/data-sources/forward_proxy_policy/properties/any_proxy/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0301302122312100-0303121321221121-2233002013020312-1211203203103312-1033012321032120-1121133311231131-1221332120300202-1103321303300103", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["any_proxy"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/any_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_proxy

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
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

- [any_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/any_proxy/#section)
- [drp_http_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/drp_http_connect/#section)
- [network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/network_connector/#section)
- [proxy_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/proxy_label_selector/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
