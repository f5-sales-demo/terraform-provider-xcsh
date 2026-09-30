---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5660, "body_sha256": "sha256:a5f2ea54c4f5d4e6ab7e0c5a13fa99528b8ef7ba91e0357c3c254b4c3875da19", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

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

- [default_header_transformation](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header--b2279382a1a969a074ca9757a7094f29.md): complete subsection reference.

- [proper_case_header_transformation](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header--8f2cde02c4ad6ed0a793772c6d1e8dc0.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header--b2279382a1a969a074ca9757a7094f29.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header--8f2cde02c4ad6ed0a793772c6d1e8dc0.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- [xcsh_workload](../resources/workload.md)
