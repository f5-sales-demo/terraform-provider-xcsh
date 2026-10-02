---
page_title: "default_pool.upstream_conn_pool_reuse_type"
subcategory: "Load Balancing"
description: "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only."
xcsh_docs: {"aliases": ["default pool upstream conn pool reuse type"], "body_bytes": 2264, "body_sha256": "sha256:368f8ca509ade5e90fae69d4272bf7c7551ab141f118e82ce531221916102c04", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "upstream_conn_pool_reuse_type"], "schema_version": 1, "sections": [{"aliases": ["disable conn pool reuse"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "upstream_conn_pool_reuse_type", "disable_conn_pool_reuse"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable conn pool reuse"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "upstream_conn_pool_reuse_type", "enable_conn_pool_reuse"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- default_pool.upstream_conn_pool_reuse_type

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

- [disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/): complete subsection reference.

- [enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/): complete subsection reference.

## Next pages

- [default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/)
- [default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
