---
page_title: "http1_config.header_transformation.default_header_transformation"
subcategory: ""
description: "http1_config.header_transformation.default_header_transformation for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1164, "body_sha256": "sha256:31427dc41d3bbce77f1d269bb42ca47b708a501bfc3b90415db5a4b9069c6bc1", "canonical_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "parent_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "path": "docs/guides/resources--cluster--properties--http1_config--header_transformation--default_header_transformation.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http1_config", "header_transformation", "default_header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/header_transformation/default_header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http1_config.header_transformation.default_header_transformation for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation.default_header_transformation

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [http1_config](resources--cluster--properties--http1_config.md)
- [http1_config.header_transformation](resources--cluster--properties--http1_config--header_transformation.md)
- http1_config.header_transformation.default_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http1_config.header_transformation](resources--cluster--properties--http1_config--header_transformation.md)
- [xcsh_cluster](../resources/cluster.md)
