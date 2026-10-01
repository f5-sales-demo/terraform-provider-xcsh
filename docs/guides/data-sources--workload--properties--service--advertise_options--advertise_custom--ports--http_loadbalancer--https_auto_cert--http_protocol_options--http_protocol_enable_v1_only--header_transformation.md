---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5038, "body_sha256": "sha256:dfd5fa695981093391ae0299b8af9d102e5e5a81c34756882117903cc420ab4f", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](data-sources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

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

- [default_header_transformation](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md)
- [xcsh_workload](../data-sources/workload.md)
