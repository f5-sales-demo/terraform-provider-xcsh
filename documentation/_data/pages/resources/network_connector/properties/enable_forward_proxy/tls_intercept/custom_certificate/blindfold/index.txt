---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.blindfold"
subcategory: "Networking"
description: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept custom certificate blindfold"], "body_bytes": 4428, "body_sha256": "sha256:71af86c350031a95aefe002d7f5d3783da007c3ea2a4f1a782ccd2bac305d0c8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/blindfold/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1323113001110323-2022313322202012-3113330302132330-1022112330001203-0211200220232130-1303033232222322-3031120302223303-1201321311233221", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold"], "schema_version": 1, "sections": [{"aliases": ["enable forward proxy tls intercept custom certificate blindfold algorithm"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--algorithm", "description": "algorithm", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold certificate file"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--certificate_file", "description": "certificate file", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "certificate_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold certificate pem"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--certificate_pem", "description": "certificate pem", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "certificate_pem"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold chain identity"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--chain_identity", "description": "chain identity", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "chain_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold context digest"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--context_digest", "description": "context digest", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "context_digest"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold encrypted location"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--encrypted_location", "description": "encrypted location", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "encrypted_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold expires at"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--expires_at", "description": "expires at", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "expires_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold fingerprint"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--fingerprint", "description": "fingerprint", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "fingerprint"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold id"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--id", "description": "id", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold material version"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--material_version", "description": "material version", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "material_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold passphrase env"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--passphrase_env", "description": "passphrase env", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "passphrase_env"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold passphrase wo"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--passphrase_wo", "description": "passphrase wo", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "passphrase_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold pkcs12 file"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--pkcs12_file", "description": "pkcs12 file", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "pkcs12_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold pkcs12 wo"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--pkcs12_wo", "description": "pkcs12 wo", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "pkcs12_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold policy"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--policy", "description": "policy", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold prepared identity"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--prepared_identity", "description": "prepared identity", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "prepared_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold private key file"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--private_key_file", "description": "private key file", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "private_key_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold private key wo"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--private_key_wo", "description": "private key wo", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "private_key_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate blindfold spki identity"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--spki_identity", "description": "spki identity", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "blindfold", "spki_identity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/blindfold/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_connectorCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.blindfold

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/)
- enable_forward_proxy.tls_intercept.custom_certificate.blindfold

<a id="section"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

## Direct properties

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--certificate_file"></a>

### certificate_file property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--certificate_pem"></a>

### certificate_pem property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--chain_identity"></a>

### chain_identity property

Type: `"string"`. Computed.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--context_digest"></a>

### context_digest property

Type: `"string"`. Computed.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--encrypted_location"></a>

### encrypted_location property

Type: `"string"`. Computed, Sensitive.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--expires_at"></a>

### expires_at property

Type: `"string"`. Computed.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--fingerprint"></a>

### fingerprint property

Type: `"string"`. Computed.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--id"></a>

### id property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--material_version"></a>

### material_version property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--passphrase_env"></a>

### passphrase_env property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--passphrase_wo"></a>

### passphrase_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--pkcs12_file"></a>

### pkcs12_file property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--pkcs12_wo"></a>

### pkcs12_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--policy"></a>

### policy property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--prepared_identity"></a>

### prepared_identity property

Type: `"string"`. Computed.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--private_key_file"></a>

### private_key_file property

Type: `"string"`. Optional.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--private_key_wo"></a>

### private_key_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--blindfold--spki_identity"></a>

### spki_identity property

Type: `"string"`. Computed.
