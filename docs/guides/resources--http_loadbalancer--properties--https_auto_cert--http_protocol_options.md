---
page_title: "https_auto_cert.http_protocol_options"
subcategory: "Load Balancing"
description: "https_auto_cert.http_protocol_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2570, "body_sha256": "sha256:6d0a7e62c0f1becc9c5a19221344afac17b683741f002456e2a4121514cef29f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:http_protocol_options", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:http_protocol_options", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert", "path": "docs/guides/resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert", "http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.http_protocol_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_auto_cert.http_protocol_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- https_auto_cert.http_protocol_options

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

- [http_protocol_enable_v1_only](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options--http_protocol_enable_v1_v2.md)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options--http_protocol_enable_v2_only.md)
- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
