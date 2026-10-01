---
page_title: "http1_config.header_transformation"
subcategory: ""
description: "http1_config.header_transformation for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2063, "body_sha256": "sha256:acab5a97c7c30c45271a4d242344aa39b762753332c2a649b78b9ed1374f7a74", "canonical_id": "xcsh-docs:data-sources:cluster:properties:http1_config:header_transformation", "child_ids": ["xcsh-docs:data-sources:cluster:properties:http1_config:header_transformation:default_header_transformation", "xcsh-docs:data-sources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:http1_config:header_transformation", "parent_id": "xcsh-docs:data-sources:cluster:properties:http1_config", "path": "docs/guides/data-sources--cluster--properties--http1_config--header_transformation.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http1_config", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http1_config.header_transformation for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Property reference](data-sources--cluster--reference.md)
- [http1_config](data-sources--cluster--properties--http1_config.md)
- http1_config.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

## Direct properties

- [default_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [http1_config.header_transformation.default_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--default_header_transformation.md)
- [http1_config.header_transformation.preserve_case_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--preserve_case_header_transformation.md)
- [http1_config.header_transformation.proper_case_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--proper_case_header_transformation.md)
- [http1_config](data-sources--cluster--properties--http1_config.md)
- [xcsh_cluster](../data-sources/cluster.md)
