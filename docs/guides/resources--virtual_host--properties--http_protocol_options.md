---
page_title: "http_protocol_options"
subcategory: ""
description: "http_protocol_options for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 2330, "body_sha256": "sha256:219e0e88ca46dedbe591266393a202c76768899265fa34741e3b7801dc54b4cb", "canonical_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "child_ids": ["xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "docs/guides/resources--virtual_host--properties--http_protocol_options.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_protocol_options for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- http_protocol_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_protocol_enable_v1_only](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_only.md)
- [http_protocol_options.http_protocol_enable_v1_v2](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v1_v2.md)
- [http_protocol_options.http_protocol_enable_v2_only](resources--virtual_host--properties--http_protocol_options--http_protocol_enable_v2_only.md)
- [Property reference](resources--virtual_host--reference.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
