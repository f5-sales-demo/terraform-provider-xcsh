---
page_title: "default_loadbalancer"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default loadbalancer"], "body_bytes": 1588, "body_sha256": "sha256:cf38fe279d530dbc87bfb7aaa42144c8077d05f10dd1cd8ee297a308ddf44491", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:default_loadbalancer", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/default_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2332201222012331-2020200003133331-3003021002300031-1121110112222301-0202020120202032-2121220121022313-1310230211023310-1000212330230231", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_loadbalancer"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/default_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_loadbalancer

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- default_loadbalancer

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_loadbalancer, non\_default\_loadbalancer; Default: default\_loadbalancer\]
Configuration parameter for default loadbalancer.

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

- [default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/default_loadbalancer/#section)
- [non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/non_default_loadbalancer/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
