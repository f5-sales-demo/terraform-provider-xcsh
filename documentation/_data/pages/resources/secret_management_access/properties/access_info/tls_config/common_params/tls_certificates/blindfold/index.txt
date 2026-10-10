---
page_title: "access_info.tls_config.common_params.tls_certificates.blindfold"
subcategory: ""
description: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."
xcsh_docs: {"aliases": ["access info tls config common params tls certificates blindfold"], "body_bytes": 4629, "body_sha256": "sha256:ea6ef42629aeb21d63558016596b865afcda937593e1dd47fc288c2f256a5b9b", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/blindfold/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2313210131113220-3222330013031203-1320320130210112-1031010220330333-1232203310032201-3313312233103033-0133013322210313-3233113232122322", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold"], "schema_version": 1, "sections": [{"aliases": ["access info tls config common params tls certificates blindfold algorithm"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--algorithm", "description": "algorithm", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold certificate file"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--certificate_file", "description": "certificate file", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "certificate_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold certificate pem"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--certificate_pem", "description": "certificate pem", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "certificate_pem"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold chain identity"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--chain_identity", "description": "chain identity", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "chain_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold context digest"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--context_digest", "description": "context digest", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "context_digest"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold encrypted location"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--encrypted_location", "description": "encrypted location", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "encrypted_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold expires at"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--expires_at", "description": "expires at", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "expires_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold fingerprint"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--fingerprint", "description": "fingerprint", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "fingerprint"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold id"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--id", "description": "id", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold material version"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--material_version", "description": "material version", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "material_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold passphrase env"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--passphrase_env", "description": "passphrase env", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "passphrase_env"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold passphrase wo"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--passphrase_wo", "description": "passphrase wo", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "passphrase_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold pkcs12 file"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--pkcs12_file", "description": "pkcs12 file", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "pkcs12_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold pkcs12 wo"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--pkcs12_wo", "description": "pkcs12 wo", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "pkcs12_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold policy"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--policy", "description": "policy", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold prepared identity"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--prepared_identity", "description": "prepared identity", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "prepared_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold private key file"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--private_key_file", "description": "private key file", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "private_key_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold private key wo"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--private_key_wo", "description": "private key wo", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "private_key_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config common params tls certificates blindfold spki identity"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--blindfold--spki_identity", "description": "spki identity", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "blindfold", "spki_identity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/blindfold/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.tls_certificates.blindfold

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/)
- [access_info.tls_config.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/)
- access_info.tls_config.common_params.tls_certificates.blindfold

<a id="section"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

## Direct properties

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--certificate_file"></a>

### certificate_file property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--certificate_pem"></a>

### certificate_pem property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--chain_identity"></a>

### chain_identity property

Type: `"string"`. Computed.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--context_digest"></a>

### context_digest property

Type: `"string"`. Computed.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--encrypted_location"></a>

### encrypted_location property

Type: `"string"`. Computed, Sensitive.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--expires_at"></a>

### expires_at property

Type: `"string"`. Computed.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--fingerprint"></a>

### fingerprint property

Type: `"string"`. Computed.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--id"></a>

### id property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--material_version"></a>

### material_version property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--passphrase_env"></a>

### passphrase_env property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--passphrase_wo"></a>

### passphrase_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--pkcs12_file"></a>

### pkcs12_file property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--pkcs12_wo"></a>

### pkcs12_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--policy"></a>

### policy property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--prepared_identity"></a>

### prepared_identity property

Type: `"string"`. Computed.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--private_key_file"></a>

### private_key_file property

Type: `"string"`. Optional.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--private_key_wo"></a>

### private_key_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-access_info--tls_config--common_params--tls_certificates--blindfold--spki_identity"></a>

### spki_identity property

Type: `"string"`. Computed.
