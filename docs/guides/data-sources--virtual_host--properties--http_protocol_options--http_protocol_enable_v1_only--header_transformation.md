---
page_title: "http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: ""
description: "http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 2798, "body_sha256": "sha256:902e1df4231954c96b5fc144826c6e4cfbd3cb524850aa1b21cbae28a600b62e", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "path": "docs/guides/data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [http_protocol_options](data-sources--virtual_host--properties--http_protocol_options.md)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation

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

- [default_header_transformation](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
