---
page_title: "http1_config.header_transformation"
subcategory: ""
description: "http1_config.header_transformation for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2612, "body_sha256": "sha256:7e0ae919f2e68d0913f0a04b5b7e7f77ec1d93c84e8a06ad04f854f85d7c3dde", "canonical_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "child_ids": ["xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "parent_id": "xcsh-docs:resources:cluster:properties:http1_config", "path": "docs/guides/resources--cluster--properties--http1_config--header_transformation.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http1_config", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http1_config.header_transformation for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [http1_config](resources--cluster--properties--http1_config.md)
- http1_config.header_transformation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_header_transformation](resources--cluster--properties--http1_config--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](resources--cluster--properties--http1_config--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](resources--cluster--properties--http1_config--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [http1_config.header_transformation.default_header_transformation](resources--cluster--properties--http1_config--header_transformation--default_header_transformation.md)
- [http1_config.header_transformation.preserve_case_header_transformation](resources--cluster--properties--http1_config--header_transformation--preserve_case_header_transformation.md)
- [http1_config.header_transformation.proper_case_header_transformation](resources--cluster--properties--http1_config--header_transformation--proper_case_header_transformation.md)
- [http1_config](resources--cluster--properties--http1_config.md)
- [xcsh_cluster](../resources/cluster.md)
