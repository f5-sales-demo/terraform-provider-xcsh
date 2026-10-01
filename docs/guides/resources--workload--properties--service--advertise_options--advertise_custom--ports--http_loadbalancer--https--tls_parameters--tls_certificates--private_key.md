---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3901, "body_sha256": "sha256:6517e20c0d69ce53ebc3d1a55fbcf3957f5baf10bfc4bae3066b4a26765e307b", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](resources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- [xcsh_workload](../resources/workload.md)
