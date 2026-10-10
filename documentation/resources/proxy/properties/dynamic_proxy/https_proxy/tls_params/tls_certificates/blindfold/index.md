---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_certificates.blindfold"
subcategory: ""
description: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."
xcsh_docs: {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold"], "body_bytes": 4506, "body_sha256": "sha256:d7b1e6dfb94756fccec3d0f6e5bf8bfc791f82329715d7a9ec99ad6b1f817959", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/blindfold/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0301223302301213-3210201223000002-2232002330033231-2331311233213120-2203121203100201-0003021013000222-0100000122312000-2333020122002201", "registry_path": "docs/guides/resources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold algorithm"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--algorithm", "description": "algorithm", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold certificate file"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--certificate_file", "description": "certificate file", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "certificate_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold certificate pem"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--certificate_pem", "description": "certificate pem", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "certificate_pem"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold chain identity"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--chain_identity", "description": "chain identity", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "chain_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold context digest"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--context_digest", "description": "context digest", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "context_digest"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold encrypted location"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--encrypted_location", "description": "encrypted location", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "encrypted_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold expires at"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--expires_at", "description": "expires at", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "expires_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold fingerprint"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--fingerprint", "description": "fingerprint", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "fingerprint"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold id"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--id", "description": "id", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold material version"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--material_version", "description": "material version", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "material_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold passphrase env"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--passphrase_env", "description": "passphrase env", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "passphrase_env"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold passphrase wo"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--passphrase_wo", "description": "passphrase wo", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "passphrase_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold pkcs12 file"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--pkcs12_file", "description": "pkcs12 file", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "pkcs12_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold pkcs12 wo"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--pkcs12_wo", "description": "pkcs12 wo", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "pkcs12_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold policy"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--policy", "description": "policy", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold prepared identity"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--prepared_identity", "description": "prepared identity", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "prepared_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold private key file"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--private_key_file", "description": "private key file", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "private_key_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold private key wo"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--private_key_wo", "description": "private key wo", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "private_key_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params tls certificates blindfold spki identity"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--spki_identity", "description": "spki identity", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "blindfold", "spki_identity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/blindfold/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.tls_certificates.blindfold

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.blindfold

<a id="section"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

## Direct properties

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--certificate_file"></a>

### certificate_file property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--certificate_pem"></a>

### certificate_pem property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--chain_identity"></a>

### chain_identity property

Type: `"string"`. Computed.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--context_digest"></a>

### context_digest property

Type: `"string"`. Computed.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--encrypted_location"></a>

### encrypted_location property

Type: `"string"`. Computed, Sensitive.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--expires_at"></a>

### expires_at property

Type: `"string"`. Computed.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--fingerprint"></a>

### fingerprint property

Type: `"string"`. Computed.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--id"></a>

### id property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--material_version"></a>

### material_version property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--passphrase_env"></a>

### passphrase_env property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--passphrase_wo"></a>

### passphrase_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--pkcs12_file"></a>

### pkcs12_file property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--pkcs12_wo"></a>

### pkcs12_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--policy"></a>

### policy property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--prepared_identity"></a>

### prepared_identity property

Type: `"string"`. Computed.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--private_key_file"></a>

### private_key_file property

Type: `"string"`. Optional.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--private_key_wo"></a>

### private_key_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--blindfold--spki_identity"></a>

### spki_identity property

Type: `"string"`. Computed.
