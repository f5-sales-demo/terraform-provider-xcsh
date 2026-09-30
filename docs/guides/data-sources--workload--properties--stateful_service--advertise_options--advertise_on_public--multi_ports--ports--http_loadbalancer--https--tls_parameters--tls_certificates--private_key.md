---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4273, "body_sha256": "sha256:b9f85d9dc0f1b0c1ad8b14334ad535300d6fcf4a68af6c1a078d4b4d92a97189", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates.md)
- [xcsh_workload](../data-sources/workload.md)
