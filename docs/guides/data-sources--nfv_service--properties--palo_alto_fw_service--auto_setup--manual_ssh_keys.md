---
page_title: "palo_alto_fw_service.auto_setup.manual_ssh_keys"
subcategory: ""
description: "palo_alto_fw_service.auto_setup.manual_ssh_keys for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2664, "body_sha256": "sha256:19b3d41e78683a6b6eaa9d16cc91d483f63a6fb4e3582f3202dec8add1db8dba", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys:private_key"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup:manual_ssh_keys", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:auto_setup", "path": "docs/guides/data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "auto_setup", "manual_ssh_keys"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/auto_setup/manual_ssh_keys/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.auto_setup.manual_ssh_keys for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.auto_setup.manual_ssh_keys

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [palo_alto_fw_service](data-sources--nfv_service--properties--palo_alto_fw_service.md)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup.md)
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

- [private_key](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key.md): complete subsection reference.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key.md)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--properties--palo_alto_fw_service--auto_setup.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
