---
page_title: "splunk_receiver.use_tls"
subcategory: ""
description: "TLS Parameters for client connection to the endpoint."
xcsh_docs: {"aliases": ["splunk receiver use tls"], "body_bytes": 3681, "body_sha256": "sha256:2c10f71aefd781d35094a5b50dc168aec50c1fd6afd6aa7c8c7acce307241da8", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:disable_verify_certificate", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:disable_verify_hostname", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:enable_verify_certificate", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:enable_verify_hostname", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:no_ca"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver", "path": "documentation/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["splunk_receiver", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["splunk receiver use tls disable verify certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:disable_verify_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "disable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls disable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:disable_verify_hostname", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "disable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls enable verify certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:enable_verify_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "enable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls enable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:enable_verify_hostname", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "enable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls mtls disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "mtls_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls mtls enable"], "anchor": "section", "description": "MTLS Client config allows configuration of mTLS client OPTIONS.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:mtls_enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "mtls_enable"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls no ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls:no_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "no_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver use tls trusted ca url"], "anchor": "schema-splunk_receiver--use_tls--trusted_ca_url", "description": "Exclusive with The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "use_tls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "TLS Parameters for client connection to the endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# splunk_receiver.use_tls

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/)
- splunk_receiver.use_tls

<a id="section"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

## Direct properties

- [disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/disable_verify_certificate/): complete subsection reference.

- [disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/disable_verify_hostname/): complete subsection reference.

- [enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/enable_verify_certificate/): complete subsection reference.

- [enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/enable_verify_hostname/): complete subsection reference.

- [mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_disabled/): complete subsection reference.

- [mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/): complete subsection reference.

- [no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/no_ca/): complete subsection reference.

<a id="schema-splunk_receiver--use_tls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```
