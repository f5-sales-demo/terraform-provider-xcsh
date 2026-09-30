---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4218, "body_sha256": "sha256:50e6c7ad311e36a932b72367fcca4fe9b1fa5f52aa321d35c1944b79dd0be7b0", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

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

- [http_protocol_enable_v1_only](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_only.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_v2.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v2_only.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [xcsh_workload](../resources/workload.md)
