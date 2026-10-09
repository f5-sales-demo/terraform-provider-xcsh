---
page_title: "blindfold"
subcategory: "Security"
description: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."
xcsh_docs: {"aliases": ["blindfold"], "body_bytes": 2696, "body_sha256": "sha256:ffd4caa5ca6b59283d4484983879650d224e3585a33147840f5e40259e68bfeb", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:blindfold", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "documentation/resources/certificate/properties/blindfold/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1133330130112321-1330211011221231-3112320213031213-1000313300133122-1013102303022130-2300232200213211-0121031230120100-0201032002231233", "registry_path": "docs/guides/resources--certificate--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blindfold"], "schema_version": 1, "sections": [{"aliases": ["blindfold algorithm"], "anchor": "schema-blindfold--algorithm", "description": "algorithm", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold certificate file"], "anchor": "schema-blindfold--certificate_file", "description": "certificate file", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "certificate_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold certificate pem"], "anchor": "schema-blindfold--certificate_pem", "description": "certificate pem", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "certificate_pem"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold chain identity"], "anchor": "schema-blindfold--chain_identity", "description": "chain identity", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "chain_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold context digest"], "anchor": "schema-blindfold--context_digest", "description": "context digest", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "context_digest"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold encrypted location"], "anchor": "schema-blindfold--encrypted_location", "description": "encrypted location", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "encrypted_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold expires at"], "anchor": "schema-blindfold--expires_at", "description": "expires at", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "expires_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold fingerprint"], "anchor": "schema-blindfold--fingerprint", "description": "fingerprint", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "fingerprint"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold id"], "anchor": "schema-blindfold--id", "description": "id", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold material version"], "anchor": "schema-blindfold--material_version", "description": "material version", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "material_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold passphrase env"], "anchor": "schema-blindfold--passphrase_env", "description": "passphrase env", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "passphrase_env"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold passphrase wo"], "anchor": "schema-blindfold--passphrase_wo", "description": "passphrase wo", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "passphrase_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold pkcs12 file"], "anchor": "schema-blindfold--pkcs12_file", "description": "pkcs12 file", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "pkcs12_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold pkcs12 wo"], "anchor": "schema-blindfold--pkcs12_wo", "description": "pkcs12 wo", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "pkcs12_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold policy"], "anchor": "schema-blindfold--policy", "description": "policy", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold prepared identity"], "anchor": "schema-blindfold--prepared_identity", "description": "prepared identity", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "prepared_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold private key file"], "anchor": "schema-blindfold--private_key_file", "description": "private key file", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "private_key_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold private key wo"], "anchor": "schema-blindfold--private_key_wo", "description": "private key wo", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "private_key_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["blindfold spki identity"], "anchor": "schema-blindfold--spki_identity", "description": "spki identity", "document_id": "xcsh-docs:resources:certificate:properties:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blindfold", "spki_identity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/blindfold/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["certificateCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blindfold

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/)
- blindfold

<a id="section"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

## Direct properties

<a id="schema-blindfold--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

<a id="schema-blindfold--certificate_file"></a>

### certificate_file property

Type: `"string"`. Optional.

<a id="schema-blindfold--certificate_pem"></a>

### certificate_pem property

Type: `"string"`. Optional.

<a id="schema-blindfold--chain_identity"></a>

### chain_identity property

Type: `"string"`. Computed.

<a id="schema-blindfold--context_digest"></a>

### context_digest property

Type: `"string"`. Computed.

<a id="schema-blindfold--encrypted_location"></a>

### encrypted_location property

Type: `"string"`. Computed, Sensitive.

<a id="schema-blindfold--expires_at"></a>

### expires_at property

Type: `"string"`. Computed.

<a id="schema-blindfold--fingerprint"></a>

### fingerprint property

Type: `"string"`. Computed.

<a id="schema-blindfold--id"></a>

### id property

Type: `"string"`. Optional.

<a id="schema-blindfold--material_version"></a>

### material_version property

Type: `"string"`. Optional.

<a id="schema-blindfold--passphrase_env"></a>

### passphrase_env property

Type: `"string"`. Optional.

<a id="schema-blindfold--passphrase_wo"></a>

### passphrase_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-blindfold--pkcs12_file"></a>

### pkcs12_file property

Type: `"string"`. Optional.

<a id="schema-blindfold--pkcs12_wo"></a>

### pkcs12_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-blindfold--policy"></a>

### policy property

Type: `"string"`. Optional.

<a id="schema-blindfold--prepared_identity"></a>

### prepared_identity property

Type: `"string"`. Computed.

<a id="schema-blindfold--private_key_file"></a>

### private_key_file property

Type: `"string"`. Optional.

<a id="schema-blindfold--private_key_wo"></a>

### private_key_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-blindfold--spki_identity"></a>

### spki_identity property

Type: `"string"`. Computed.
