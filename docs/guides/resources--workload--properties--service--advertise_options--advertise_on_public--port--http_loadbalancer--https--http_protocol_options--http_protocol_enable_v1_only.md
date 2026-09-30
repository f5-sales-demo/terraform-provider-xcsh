---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2763, "body_sha256": "sha256:ca46eaa6e03706108194cb8e6b4e1c02394140eb90b32581d40279841a2cd759", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_only.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](resources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

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

- [header_transformation](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options.md)
- [xcsh_workload](../resources/workload.md)
