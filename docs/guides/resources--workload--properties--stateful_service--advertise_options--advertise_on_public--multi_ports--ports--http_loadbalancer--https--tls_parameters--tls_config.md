---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5742, "body_sha256": "sha256:4c3087e36a5f91f0412a108fc181072bcf20d824d9c6b77fb2025ab193d6011b", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_config", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_config:custom_security", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_config:default_security", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_config:low_security", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_config", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters.md)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

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

- [custom_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--custom_security.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--default_security.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--low_security.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_config--medium_security.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters.md)
- [xcsh_workload](../resources/workload.md)
