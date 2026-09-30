---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4448, "body_sha256": "sha256:e9a18200d3b10f7ee4cb96cc6cf66403d708595f11e2f3708b87e672d9c722e8", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:default_security", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:low_security", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_cert_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

## Direct properties

- [custom_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--custom_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--default_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--low_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config--medium_security.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params.md)
- [xcsh_workload](../data-sources/workload.md)
