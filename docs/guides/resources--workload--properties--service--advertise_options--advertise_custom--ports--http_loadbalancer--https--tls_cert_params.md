---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3958, "body_sha256": "sha256:71ee335f675a72a1333152549761efb44ab3cc27fb4d347bc273f388be4967d2", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:certificates", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:no_mtls", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](resources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
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
tls_cert_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [certificates](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--certificates.md): complete subsection reference.

- [no_mtls](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--no_mtls.md): complete subsection reference.

- [tls_config](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config.md): complete subsection reference.

- [use_mtls](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--certificates.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--no_mtls.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--tls_config.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- [xcsh_workload](../resources/workload.md)
