---
page_title: "palo_alto_fw_service.auto_setup.manual_ssh_keys"
subcategory: ""
description: "SSH Key includes both public and private key."
xcsh_docs: {"aliases": ["palo alto fw service auto setup manual ssh keys"], "body_bytes": 2675, "body_sha256": "sha256:356b7737ecca7a0b85392158c0e12e1d93c9d52f16352a9b1a39c585b022691c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service auto setup manual ssh keys private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys", "private_key"], "syntax": "block", "type": "object"}, {"aliases": ["palo alto fw service auto setup manual ssh keys public key"], "anchor": "schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key", "description": "Authorized Public SSH key which will be programmed on the node.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys", "public_key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "SSH Key includes both public and private key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup.manual_ssh_keys

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/)
- palo_alto_fw_service.auto_setup.manual_ssh_keys

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SSH Key includes both public and private key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
manual_ssh_keys {
  # Configure direct properties listed below.
}
```

## Direct properties

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/private_key/): complete subsection reference.

<a id="schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key"></a>

### public_key property

Type: `"string"`. Optional.

Authorized Public SSH key which will be programmed on the node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----\\n.*\\n-----END (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
