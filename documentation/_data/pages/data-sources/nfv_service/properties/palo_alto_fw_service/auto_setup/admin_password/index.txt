---
page_title: "palo_alto_fw_service.auto_setup.admin_password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["palo alto fw service auto setup admin password"], "body_bytes": 2336, "body_sha256": "sha256:ee03d2d9121fc43e5c8c51667ce6b99d4a613b35f1b7835dce881a34d98e8360", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:blindfold_secret_info", "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2033133133233333-0310333313110010-2321223111220000-3300033010122102-2223130020001101-0310011313113301-0330330002023330-2030303233222223", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service auto setup admin password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["palo alto fw service auto setup admin password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup.admin_password

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/clear_secret_info/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/blindfold_secret_info/)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/clear_secret_info/)
- [palo_alto_fw_service.auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
