---
page_title: "http_protocol_options"
subcategory: ""
description: "http_protocol_options for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1808, "body_sha256": "sha256:b5f95449e38567d2757a6a9ab97e9c4bb35011e9e3e20d3923305131e6d5d208", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:http_protocol_options", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "docs/guides/data-sources--virtual_host--properties--http_protocol_options.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_protocol_options for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- http_protocol_options

<a id="section"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

## Direct properties

- [http_protocol_enable_v1_only](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md)
- [http_protocol_options.http_protocol_enable_v1_v2](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_v2.md)
- [http_protocol_options.http_protocol_enable_v2_only](data-sources--virtual_host--properties--http_protocol_options--http_protocol_enable_v2_only.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
