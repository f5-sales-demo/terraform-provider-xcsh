---
page_title: "https_management.advertise_on_slo_internet_vip.tls_certificates.private_key"
subcategory: ""
description: "https_management.advertise_on_slo_internet_vip.tls_certificates.private_key for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2613, "body_sha256": "sha256:e277079f46a7c8bd8116e72ae4b5b61febb511385c3b3e2ddaadbe7b133b167d", "canonical_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "path": "docs/guides/resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_internet_vip.tls_certificates.private_key for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_internet_vip.tls_certificates.private_key

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip.md)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key

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

- [blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info.md)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info.md)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
