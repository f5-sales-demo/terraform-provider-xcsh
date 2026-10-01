---
page_title: "http_protocol_options.http_protocol_enable_v1_only"
subcategory: ""
description: "http_protocol_options.http_protocol_enable_v1_only for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1388, "body_sha256": "sha256:aaa665a3fd9de11376bcefc52cd4aa5b84763f0f11f7926daabeaca66fb7b887", "canonical_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "child_ids": ["xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "path": "docs/guides/resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_protocol_options.http_protocol_enable_v1_only for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [http_protocol_options](resources--virtual_host--properties--http_protocol_options.md)
- http_protocol_options.http_protocol_enable_v1_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md)
- [http_protocol_options](resources--virtual_host--properties--http_protocol_options.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
