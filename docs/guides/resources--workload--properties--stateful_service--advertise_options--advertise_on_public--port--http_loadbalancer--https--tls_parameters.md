---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4148, "body_sha256": "sha256:1b1fa07f4acfaa779de2c9d99208d2a16aa4af70bd59c0f4d072e7c9d046b40d", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:no_mtls", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_certificates", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:tls_config", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_mtls](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--no_mtls.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--tls_config.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [xcsh_workload](../resources/workload.md)
