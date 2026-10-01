---
page_title: "http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation"
subcategory: ""
description: "http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1600, "body_sha256": "sha256:83e9485aa9a40c1d43be168d86b7c05137f4acb242ef2141220a16dbaa48c1c3", "canonical_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "parent_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "path": "docs/guides/resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "default_header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [http_protocol_options](resources--virtual_host--properties--http_protocol_options.md)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

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

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
