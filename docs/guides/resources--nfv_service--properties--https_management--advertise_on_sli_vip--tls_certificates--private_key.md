---
page_title: "https_management.advertise_on_sli_vip.tls_certificates.private_key"
subcategory: ""
description: "https_management.advertise_on_sli_vip.tls_certificates.private_key for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2487, "body_sha256": "sha256:4e07cbcb99f189283be2dd2bcea5d59498b05b6f50d90f0d83f33d1eaa489e9c", "canonical_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates", "path": "docs/guides/resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_sli_vip", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_sli_vip.tls_certificates.private_key for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_sli_vip.tls_certificates.private_key

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_sli_vip](resources--nfv_service--properties--https_management--advertise_on_sli_vip.md)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md)
- https_management.advertise_on_sli_vip.tls_certificates.private_key

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

- [blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info.md)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info.md)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
