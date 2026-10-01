---
page_title: "palo_alto_fw_service.auto_setup.admin_password"
subcategory: ""
description: "palo_alto_fw_service.auto_setup.admin_password for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1838, "body_sha256": "sha256:6d4664bfd81ecff04fd782f388ea6c1efb9f651c13f1252c013ac2a0133fb697", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:blindfold_secret_info", "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "docs/guides/data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.auto_setup.admin_password for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup.admin_password

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [palo_alto_fw_service](data-sources--nfv_service--properties--palo_alto_fw_service.md)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup.md)
- palo_alto_fw_service.auto_setup.admin_password

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

- [blindfold_secret_info](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md): complete subsection reference.

## Next pages

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
