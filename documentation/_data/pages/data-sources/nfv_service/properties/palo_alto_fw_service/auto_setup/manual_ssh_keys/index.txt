---
page_title: "palo_alto_fw_service.auto_setup.manual_ssh_keys"
subcategory: ""
description: "SSH Key includes both public and private key."
xcsh_docs: {"aliases": ["palo alto fw service auto setup manual ssh keys"], "body_bytes": 2566, "body_sha256": "sha256:0a5a7924e73aa192656684a34a8e86fffcbfd5fb65c59f93814c8062b7097b09", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service auto setup manual ssh keys private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["palo alto fw service auto setup manual ssh keys public key"], "anchor": "schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key", "description": "Authorized Public SSH key which will be programmed on the node.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys", "public_key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "SSH Key includes both public and private key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup.manual_ssh_keys

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/)
- palo_alto_fw_service.auto_setup.manual_ssh_keys

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/private_key/): complete subsection reference.

<a id="schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key"></a>

### public_key property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
