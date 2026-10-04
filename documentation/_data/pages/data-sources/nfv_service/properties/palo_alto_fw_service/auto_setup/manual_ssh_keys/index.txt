---
page_title: "palo_alto_fw_service.auto_setup.manual_ssh_keys"
subcategory: ""
description: "SSH Key includes both public and private key."
xcsh_docs: {"aliases": ["palo alto fw service auto setup manual ssh keys"], "body_bytes": 3066, "body_sha256": "sha256:5dd1abe70f6b216084485f4e638bf844c275ba3097e7d9afb834e861871dfd8d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service auto setup manual ssh keys private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["palo alto fw service auto setup manual ssh keys public key"], "anchor": "schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key", "description": "Authorized Public SSH key which will be programmed on the node.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys", "public_key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SSH Key includes both public and private key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/private_key/)
- [palo_alto_fw_service.auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
