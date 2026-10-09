---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.blindfold"
subcategory: "Load Balancing"
description: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."
xcsh_docs: {"aliases": ["origin pool use tls use mtls tls certificates blindfold"], "body_bytes": 4367, "body_sha256": "sha256:7eae40076e4304126251db429b55078ec67cd61c15b2a71a45ab2cca3849bfe8", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/blindfold/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2213232302132010-1003123212311122-3303103232103020-1213303132020302-0330220331302202-2013313313102111-3212030303000132-0121220232120333", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold"], "schema_version": 1, "sections": [{"aliases": ["origin pool use tls use mtls tls certificates blindfold algorithm"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--algorithm", "description": "algorithm", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold certificate file"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--certificate_file", "description": "certificate file", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "certificate_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold certificate pem"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--certificate_pem", "description": "certificate pem", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "certificate_pem"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold chain identity"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--chain_identity", "description": "chain identity", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "chain_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold context digest"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--context_digest", "description": "context digest", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "context_digest"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold encrypted location"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--encrypted_location", "description": "encrypted location", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "encrypted_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold expires at"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--expires_at", "description": "expires at", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "expires_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold fingerprint"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--fingerprint", "description": "fingerprint", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "fingerprint"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold id"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--id", "description": "id", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold material version"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--material_version", "description": "material version", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "material_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold passphrase env"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--passphrase_env", "description": "passphrase env", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "passphrase_env"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold passphrase wo"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--passphrase_wo", "description": "passphrase wo", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "passphrase_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold pkcs12 file"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--pkcs12_file", "description": "pkcs12 file", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "pkcs12_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold pkcs12 wo"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--pkcs12_wo", "description": "pkcs12 wo", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "pkcs12_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold policy"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--policy", "description": "policy", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold prepared identity"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--prepared_identity", "description": "prepared identity", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "prepared_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold private key file"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--private_key_file", "description": "private key file", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "private_key_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold private key wo"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--private_key_wo", "description": "private key wo", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "private_key_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pool use tls use mtls tls certificates blindfold spki identity"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--spki_identity", "description": "spki identity", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "blindfold", "spki_identity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/blindfold/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls.tls_certificates.blindfold

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [origin_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/)
- [origin_pool.use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/)
- origin_pool.use_tls.use_mtls.tls_certificates.blindfold

<a id="section"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

## Direct properties

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--certificate_file"></a>

### certificate_file property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--certificate_pem"></a>

### certificate_pem property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--chain_identity"></a>

### chain_identity property

Type: `"string"`. Computed.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--context_digest"></a>

### context_digest property

Type: `"string"`. Computed.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--encrypted_location"></a>

### encrypted_location property

Type: `"string"`. Computed, Sensitive.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--expires_at"></a>

### expires_at property

Type: `"string"`. Computed.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--fingerprint"></a>

### fingerprint property

Type: `"string"`. Computed.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--id"></a>

### id property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--material_version"></a>

### material_version property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--passphrase_env"></a>

### passphrase_env property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--passphrase_wo"></a>

### passphrase_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--pkcs12_file"></a>

### pkcs12_file property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--pkcs12_wo"></a>

### pkcs12_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--policy"></a>

### policy property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--prepared_identity"></a>

### prepared_identity property

Type: `"string"`. Computed.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--private_key_file"></a>

### private_key_file property

Type: `"string"`. Optional.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--private_key_wo"></a>

### private_key_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--blindfold--spki_identity"></a>

### spki_identity property

Type: `"string"`. Computed.
