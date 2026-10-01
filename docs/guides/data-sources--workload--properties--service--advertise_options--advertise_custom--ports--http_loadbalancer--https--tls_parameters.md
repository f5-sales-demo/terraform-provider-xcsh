---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3652, "body_sha256": "sha256:3f721e7ef9e708878782c87bf91094b970418355353dab2e4a162adac2af8df4", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:no_mtls", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_config", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](data-sources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

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

- [no_mtls](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--no_mtls.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_config.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--use_mtls.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- [xcsh_workload](../data-sources/workload.md)
