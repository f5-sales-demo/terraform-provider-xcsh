---
page_title: "palo_alto_fw_service.auto_setup"
subcategory: ""
description: "palo_alto_fw_service.auto_setup for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2941, "body_sha256": "sha256:9b19e70cab0266b825f2a74592d2f13613d43c8e31b5a070f3804c9a641ce7e0", "canonical_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "path": "docs/guides/resources--nfv_service--properties--palo_alto_fw_service--auto_setup.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.auto_setup for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# palo_alto_fw_service.auto_setup

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
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

- [admin_password](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password.md): complete subsection reference.

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

- [manual_ssh_keys](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys.md): complete subsection reference.

## Next pages

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password.md)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
