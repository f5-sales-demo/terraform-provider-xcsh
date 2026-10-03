---
page_title: "any_proxy"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["any proxy"], "body_bytes": 1890, "body_sha256": "sha256:37586b48a089108cd6bb526c475cf851abdf44d9b2fb13968a04b22928c23f57", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:any_proxy", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "documentation/resources/forward_proxy_policy/properties/any_proxy/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2332333300301220-2020011003301331-0311130310031221-0302003113033233-1113111033022331-1120333131201021-0101002021020003-0020122103301200", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["any_proxy"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/any_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_proxy

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- any_proxy

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

- [any_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/any_proxy/#section)
- [drp_http_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/drp_http_connect/#section)
- [network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/#section)
- [proxy_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/proxy_label_selector/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_proxy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
