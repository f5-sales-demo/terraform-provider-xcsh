---
page_title: "upstream_conn_pool_reuse_type"
subcategory: ""
description: "upstream_conn_pool_reuse_type for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1824, "body_sha256": "sha256:1e2e689fb5377f4d2a8d346f9129279e4d2ba92a781555c5da03f1a52a639c70", "child_ids": ["xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/upstream_conn_pool_reuse_type/index.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upstream_conn_pool_reuse_type for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
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

- [disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/): complete subsection reference.

- [enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/): complete subsection reference.

## Next pages

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
