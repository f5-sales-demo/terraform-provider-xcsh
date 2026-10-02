---
page_title: "default_lb_with_sni"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default lb with sni"], "body_bytes": 1686, "body_sha256": "sha256:c73145c2565ab34ccc28d07abac2f33d7d9f46c41062c2f4d1e31948ebd02589", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:default_lb_with_sni", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "documentation/data-sources/tcp_loadbalancer/properties/default_lb_with_sni/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1120310112113021-0011220322222102-0130122230032031-0102010320303202-0001311211211213-3321333330113130-2001230311133133-1322321301312102", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_lb_with_sni"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/default_lb_with_sni/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_lb_with_sni

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- default_lb_with_sni

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

- [default_lb_with_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/default_lb_with_sni/#section)
- [no_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/no_sni/#section)
- [sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/sni/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
