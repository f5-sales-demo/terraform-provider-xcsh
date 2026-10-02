---
page_title: "drp_http_connect"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["drp http connect"], "body_bytes": 1166, "body_sha256": "sha256:8d6ca7c3540df5f4d78f1b48921d9086dd05f7d7b01f9ba9e497a0e1cf0cc697", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:drp_http_connect", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "documentation/resources/forward_proxy_policy/properties/drp_http_connect/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0121000321310323-0220232133302100-1210330110221132-1112130320100200-1022002100133130-3210200311003310-3011210001322311-1121030003102231", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["drp_http_connect"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/drp_http_connect/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# drp_http_connect

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- drp_http_connect

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for drp http connect.

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
drp_http_connect = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
