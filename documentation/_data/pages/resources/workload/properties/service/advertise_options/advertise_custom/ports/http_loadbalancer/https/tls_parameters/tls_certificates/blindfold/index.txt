---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold"
subcategory: "Container"
description: "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold"], "body_bytes": 6698, "body_sha256": "sha256:2b663ac46eb76ed8bbebdef1f0d9950965483fe50462e4603689ae72630fff72", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_parameters/tls_certificates/blindfold/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2201300111232030-2000213230220101-3320220231010011-0001111230300101-2101122311223021-1232100113232301-0232200130230012-3022332213112111", "registry_path": "docs/guides/resources--workload--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold algorithm"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--algorithm", "description": "algorithm", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold certificate file"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--certificate_file", "description": "certificate file", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "certificate_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold certificate pem"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--certificate_pem", "description": "certificate pem", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "certificate_pem"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold chain identity"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--chain_identity", "description": "chain identity", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "chain_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold context digest"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--context_digest", "description": "context digest", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "context_digest"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold encrypted location"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--encrypted_location", "description": "encrypted location", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "encrypted_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold expires at"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--expires_at", "description": "expires at", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "expires_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold fingerprint"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--fingerprint", "description": "fingerprint", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "fingerprint"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold id"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--id", "description": "id", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold material version"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--material_version", "description": "material version", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "material_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold passphrase env"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--passphrase_env", "description": "passphrase env", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "passphrase_env"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold passphrase wo"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--passphrase_wo", "description": "passphrase wo", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "passphrase_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold pkcs12 file"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--pkcs12_file", "description": "pkcs12 file", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "pkcs12_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold pkcs12 wo"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--pkcs12_wo", "description": "pkcs12 wo", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "pkcs12_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold policy"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--policy", "description": "policy", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold prepared identity"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--prepared_identity", "description": "prepared identity", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "prepared_identity"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold private key file"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--private_key_file", "description": "private key file", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "private_key_file"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold private key wo"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--private_key_wo", "description": "private key wo", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive", "write_only"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "private_key_wo"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https tls parameters tls certificates blindfold spki identity"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--spki_identity", "description": "spki identity", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_parameters:tls_certificates:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "blindfold", "spki_identity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_parameters/tls_certificates/blindfold/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_parameters/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_parameters/tls_certificates/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

<a id="section"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

## Direct properties

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--certificate_file"></a>

### certificate_file property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--certificate_pem"></a>

### certificate_pem property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--chain_identity"></a>

### chain_identity property

Type: `"string"`. Computed.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--context_digest"></a>

### context_digest property

Type: `"string"`. Computed.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--encrypted_location"></a>

### encrypted_location property

Type: `"string"`. Computed, Sensitive.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--expires_at"></a>

### expires_at property

Type: `"string"`. Computed.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--fingerprint"></a>

### fingerprint property

Type: `"string"`. Computed.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--id"></a>

### id property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--material_version"></a>

### material_version property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--passphrase_env"></a>

### passphrase_env property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--passphrase_wo"></a>

### passphrase_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--pkcs12_file"></a>

### pkcs12_file property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--pkcs12_wo"></a>

### pkcs12_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--policy"></a>

### policy property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--prepared_identity"></a>

### prepared_identity property

Type: `"string"`. Computed.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--private_key_file"></a>

### private_key_file property

Type: `"string"`. Optional.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--private_key_wo"></a>

### private_key_wo property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_parameters--tls_certificates--blindfold--spki_identity"></a>

### spki_identity property

Type: `"string"`. Computed.
