---
page_title: "palo_alto_fw_service.auto_setup.admin_password"
subcategory: ""
description: "palo_alto_fw_service.auto_setup.admin_password for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2020, "body_sha256": "sha256:fbd063c98603f79432e6b1b42e17f7c466744839902152b470be844be82cc6f0", "canonical_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:blindfold_secret_info", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "docs/guides/resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.auto_setup.admin_password for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# palo_alto_fw_service.auto_setup.admin_password

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--properties--palo_alto_fw_service--auto_setup.md)
- palo_alto_fw_service.auto_setup.admin_password

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
admin_password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md): complete subsection reference.

## Next pages

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--properties--palo_alto_fw_service--auto_setup.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
