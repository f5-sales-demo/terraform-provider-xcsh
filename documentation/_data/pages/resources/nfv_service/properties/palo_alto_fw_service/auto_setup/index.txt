---
page_title: "palo_alto_fw_service.auto_setup"
subcategory: ""
description: "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured."
xcsh_docs: {"aliases": ["palo alto fw service auto setup"], "body_bytes": 3493, "body_sha256": "sha256:f36114e0980c73051324505132ceda81f6af5fcd915c8965eadd66a86d5ff944", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123", "registry_path": "docs/guides/resources--nfv_service--reference--group-003.md", "relationships": [{"anchor": "schema-palo_alto_fw_service--auto_setup--admin_username", "enforcement": "provider-schema", "group": "palo_alto_fw_service.auto_setup:RequiredObjectAttributes:admin_username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup"], "schema_version": 1, "sections": [{"aliases": ["admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.auto_setup.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.auto_setup.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password:clear_secret_info", "type": "conflicts"}], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password"], "syntax": "block", "type": "object"}, {"aliases": ["admin username"], "anchor": "schema-palo_alto_fw_service--auto_setup--admin_username", "description": "Firewall Admin Username.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_username"], "syntax": "attribute", "type": "string"}, {"aliases": ["manual ssh keys"], "anchor": "section", "description": "SSH Key includes both public and private key.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key", "enforcement": "provider-schema", "group": "palo_alto_fw_service.auto_setup.manual_ssh_keys:RequiredObjectAttributes:public_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- palo_alto_fw_service.auto_setup

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Upstream description:

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("admin_username")}
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
  "x-ves-oneof-field-ssh_keys_choice": "[\"manual_ssh_keys\"]"
}
```

Terraform syntax:

```terraform
auto_setup {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/): complete subsection reference.

<a id="schema-palo_alto_fw_service--auto_setup--admin_username"></a>

### admin_username property

Type: `"string"`. Optional.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [manual_ssh_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.auto_setup.admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
