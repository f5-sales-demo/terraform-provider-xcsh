---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5130, "body_sha256": "sha256:8483fb341f0be37f74edac781a8a1d609d7ad1f6ad3586ef12453f3bc62c9fc3", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

## Direct properties

- [default_header_transformation](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--h--b8082ea568bda581ea5732e3db3585fb.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--h--08a8944a083c2fc81ad767594d8787d0.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--h--e85fc9129683f744fb2bdc2618977a80.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--h--b8082ea568bda581ea5732e3db3585fb.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--h--08a8944a083c2fc81ad767594d8787d0.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--h--e85fc9129683f744fb2bdc2618977a80.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- [xcsh_workload](../data-sources/workload.md)
