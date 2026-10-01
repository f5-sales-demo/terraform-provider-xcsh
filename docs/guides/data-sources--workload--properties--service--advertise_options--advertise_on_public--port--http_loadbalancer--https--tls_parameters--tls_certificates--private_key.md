---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3687, "body_sha256": "sha256:3bd3c022a6d9c328b265c419a4df5a633ed78b7a8b7cd4e0edbaefa76e040831", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates--private_key.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- [xcsh_workload](../data-sources/workload.md)
