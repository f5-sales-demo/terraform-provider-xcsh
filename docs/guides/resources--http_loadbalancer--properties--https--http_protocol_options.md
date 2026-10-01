---
page_title: "https.http_protocol_options"
subcategory: "Load Balancing"
description: "https.http_protocol_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2519, "body_sha256": "sha256:498cb8dcf7c574df18044b9ff553aaa4ef4dae99dd8b59dd08b7f2a5ea99ae22", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https", "path": "docs/guides/resources--http_loadbalancer--properties--https--http_protocol_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.http_protocol_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.http_protocol_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- https.http_protocol_options

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

- [http_protocol_enable_v1_only](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_only.md)
- [https.http_protocol_options.http_protocol_enable_v1_v2](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v1_v2.md)
- [https.http_protocol_options.http_protocol_enable_v2_only](resources--http_loadbalancer--properties--https--http_protocol_options--http_protocol_enable_v2_only.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
