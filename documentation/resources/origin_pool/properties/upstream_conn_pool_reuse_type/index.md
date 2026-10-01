---
page_title: "upstream_conn_pool_reuse_type"
subcategory: "Load Balancing"
description: "upstream_conn_pool_reuse_type for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2272, "body_sha256": "sha256:946b6c07721d47d19c80dc35af0d3c5773fb618095185129f5e52b43bda9c144", "child_ids": ["xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:resources:origin_pool:reference", "path": "documentation/resources/origin_pool/properties/upstream_conn_pool_reuse_type/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upstream_conn_pool_reuse_type for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
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

- [disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/): complete subsection reference.

- [enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/): complete subsection reference.

## Next pages

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
