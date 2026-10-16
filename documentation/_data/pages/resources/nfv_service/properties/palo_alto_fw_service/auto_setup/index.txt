---
page_title: "palo_alto_fw_service.auto_setup"
subcategory: ""
description: "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured."
xcsh_docs: {"aliases": ["palo alto fw service auto setup"], "body_bytes": 2375, "body_sha256": "sha256:f38eb9f67e72618922d6739d1a5f592aab5cf82dd088a35922ae42ed5ae9f8bb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service auto setup admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password"], "syntax": "block", "type": "object"}, {"aliases": ["palo alto fw service auto setup admin username"], "anchor": "schema-palo_alto_fw_service--auto_setup--admin_username", "description": "Firewall Admin Username.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_username"], "syntax": "attribute", "type": "string"}, {"aliases": ["palo alto fw service auto setup manual ssh keys"], "anchor": "section", "description": "SSH Key includes both public and private key.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
