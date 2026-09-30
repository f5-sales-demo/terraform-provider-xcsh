---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3639, "body_sha256": "sha256:2270f77b792a790b79830a98a5ad1af50bd0f7d8db06a0d4860fea6310539018", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

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

- [http_protocol_enable_v1_only](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_v2.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v2_only.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert.md)
- [xcsh_workload](../data-sources/workload.md)
