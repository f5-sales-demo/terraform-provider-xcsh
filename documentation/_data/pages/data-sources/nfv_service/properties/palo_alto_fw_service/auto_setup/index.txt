---
page_title: "palo_alto_fw_service.auto_setup"
subcategory: ""
description: "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured."
xcsh_docs: {"aliases": ["palo alto fw service auto setup"], "body_bytes": 3110, "body_sha256": "sha256:f54dbf5055f9e4cc499f89c69671b186d7cac94e0f6f29d1616bbe8945690846", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup"], "schema_version": 1, "sections": [{"aliases": ["admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:admin_password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_password"], "syntax": "attribute", "type": "object"}, {"aliases": ["admin username"], "anchor": "schema-palo_alto_fw_service--auto_setup--admin_username", "description": "Firewall Admin Username.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "admin_username"], "syntax": "attribute", "type": "string"}, {"aliases": ["manual ssh keys"], "anchor": "section", "description": "SSH Key includes both public and private key.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- palo_alto_fw_service.auto_setup

<a id="section"></a>

Type: `"single"`. Computed.

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Upstream description:

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

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/): complete subsection reference.

<a id="schema-palo_alto_fw_service--auto_setup--admin_username"></a>

### admin_username property

Type: `"string"`. Computed.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

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

- [manual_ssh_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.auto_setup.admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/admin_password/)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
