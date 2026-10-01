---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4987, "body_sha256": "sha256:102afc761ca2c5aa2ef9f8082f03ba3ba1291484d2107f193736f2261e4599ab", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config:default_security", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config:low_security", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "tls_cert_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](resources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--custom_security.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--default_security.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--low_security.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params--tls_config--medium_security.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params.md)
- [xcsh_workload](../resources/workload.md)
