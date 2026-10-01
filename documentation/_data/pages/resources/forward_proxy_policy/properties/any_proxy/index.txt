---
page_title: "any_proxy"
subcategory: "Security"
description: "any_proxy for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1890, "body_sha256": "sha256:37586b48a089108cd6bb526c475cf851abdf44d9b2fb13968a04b22928c23f57", "child_ids": [], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:any_proxy", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "documentation/resources/forward_proxy_policy/properties/any_proxy/index.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["any_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/any_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_proxy for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
