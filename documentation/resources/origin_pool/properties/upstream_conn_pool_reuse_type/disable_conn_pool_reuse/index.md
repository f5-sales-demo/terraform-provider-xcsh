---
page_title: "upstream_conn_pool_reuse_type.disable_conn_pool_reuse"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["upstream conn pool reuse type disable conn pool reuse"], "body_bytes": 1396, "body_sha256": "sha256:3219a2a482af6f4655b40e7e264a91167dffa58d7a8c27f8f69111afa043758d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "parent_id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "path": "documentation/resources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323", "registry_path": "docs/guides/resources--origin_pool--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type", "disable_conn_pool_reuse"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type.disable_conn_pool_reuse

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
