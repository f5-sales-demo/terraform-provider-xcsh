---
page_title: "upstream_conn_pool_reuse_type"
subcategory: "Load Balancing"
description: "upstream_conn_pool_reuse_type for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1864, "body_sha256": "sha256:99ff8e97d3cd0c55c2d4f7a15082d14f82904689d7db44ecba7e53bc4967683a", "canonical_id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "child_ids": ["xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:resources:origin_pool:reference", "path": "docs/guides/resources--origin_pool--properties--upstream_conn_pool_reuse_type.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upstream_conn_pool_reuse_type for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- upstream_conn_pool_reuse_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
```

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

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_conn_pool_reuse](resources--origin_pool--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md): complete subsection reference.

- [enable_conn_pool_reuse](resources--origin_pool--properties--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md): complete subsection reference.

## Next pages

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--origin_pool--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--origin_pool--properties--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md)
- [Property reference](resources--origin_pool--reference.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
