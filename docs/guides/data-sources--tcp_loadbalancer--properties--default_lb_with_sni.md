---
page_title: "default_lb_with_sni"
subcategory: "Load Balancing"
description: "default_lb_with_sni for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1325, "body_sha256": "sha256:a154000dda0ea5332400061379e1f0548b7fb47864a7e716163b7483c611cd9d", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:default_lb_with_sni", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:default_lb_with_sni", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--default_lb_with_sni.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_lb_with_sni"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/default_lb_with_sni/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_lb_with_sni for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_lb_with_sni

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
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

- [default_lb_with_sni](data-sources--tcp_loadbalancer--properties--default_lb_with_sni.md#section)
- [no_sni](data-sources--tcp_loadbalancer--properties--no_sni.md#section)
- [sni](data-sources--tcp_loadbalancer--properties--sni.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
