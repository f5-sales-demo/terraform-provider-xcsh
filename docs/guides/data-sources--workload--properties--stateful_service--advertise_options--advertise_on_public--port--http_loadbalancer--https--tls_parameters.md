---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3855, "body_sha256": "sha256:3a97adf636e7a4b1f0a644056054128240ca05a9124da01eef313dcf0e84517d", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:no_mtls", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_config", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

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

- [no_mtls](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--no_mtls.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_config.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [xcsh_workload](../data-sources/workload.md)
