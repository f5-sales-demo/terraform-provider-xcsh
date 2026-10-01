---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4875, "body_sha256": "sha256:1cc5c55ce000385c357bdabbe72cd9401a42eb809239519627367e7d4f0071e6", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:tls_config", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:tls_config:custom_security", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:tls_config:default_security", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:tls_config:low_security", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:tls_config", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

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

- [custom_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--custom_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--default_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--low_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--tls_config--medium_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert.md)
- [xcsh_workload](../resources/workload.md)
