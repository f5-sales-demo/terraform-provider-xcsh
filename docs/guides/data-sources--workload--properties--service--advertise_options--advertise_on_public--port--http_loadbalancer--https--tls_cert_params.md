---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3718, "body_sha256": "sha256:858ae5015cdaa3f1949da0f69ae265fce08abbfc2c76754273f4cc3c309c05be", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:certificates", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:no_mtls", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

## Direct properties

- [certificates](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--certificates.md): complete subsection reference.

- [no_mtls](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--no_mtls.md): complete subsection reference.

- [tls_config](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--use_mtls.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--certificates.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--no_mtls.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--use_mtls.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [xcsh_workload](../data-sources/workload.md)
