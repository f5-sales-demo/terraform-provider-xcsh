---
page_title: "upstream_conn_pool_reuse_type"
subcategory: "Load Balancing"
description: "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only."
xcsh_docs: {"aliases": ["upstream conn pool reuse type"], "body_bytes": 1963, "body_sha256": "sha256:79ff4a46eaf07fcd0c48f0c4cbe3364045daee8572546705a9d0febbabb22a7b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:data-sources:origin_pool:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "documentation/data-sources/origin_pool/properties/upstream_conn_pool_reuse_type/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type"], "schema_version": 1, "sections": [{"aliases": ["upstream conn pool reuse type disable conn pool reuse"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upstream_conn_pool_reuse_type", "disable_conn_pool_reuse"], "syntax": "attribute", "type": "object"}, {"aliases": ["upstream conn pool reuse type enable conn pool reuse"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upstream_conn_pool_reuse_type", "enable_conn_pool_reuse"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- upstream_conn_pool_reuse_type

<a id="section"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

## Direct properties

- [disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/): complete subsection reference.

- [enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/): complete subsection reference.

## Next pages

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
