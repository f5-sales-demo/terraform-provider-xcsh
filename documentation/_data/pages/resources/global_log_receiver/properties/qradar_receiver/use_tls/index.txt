---
page_title: "qradar_receiver.use_tls"
subcategory: ""
description: "TLS Parameters for client connection to the endpoint."
xcsh_docs: {"aliases": ["qradar receiver use tls"], "body_bytes": 4418, "body_sha256": "sha256:582dcd85dfc5480d075f2c98f913b943a6a5e12cf9d1f46ec00c46442d897d94", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_certificate", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_hostname", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_certificate", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_hostname", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_disabled", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:no_ca"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "path": "documentation/resources/global_log_receiver/properties/qradar_receiver/use_tls/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [{"anchor": "schema-qradar_receiver--use_tls--trusted_ca_url", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:no_ca", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["qradar_receiver", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["qradar receiver use tls disable verify certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "disable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver use tls disable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_hostname", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "disable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver use tls enable verify certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "enable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver use tls enable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_hostname", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "enable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver use tls mtls disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "mtls_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver use tls mtls enable"], "anchor": "section", "description": "MTLS Client config allows configuration of mTLS client OPTIONS.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "mtls_enable"], "syntax": "block", "type": "object"}, {"aliases": ["qradar receiver use tls no ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:no_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "no_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver use tls trusted ca url"], "anchor": "schema-qradar_receiver--use_tls--trusted_ca_url", "description": "Exclusive with The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/qradar_receiver/use_tls/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "TLS Parameters for client connection to the endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.use_tls

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/)
- qradar_receiver.use_tls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
    "trusted_ca_url")}
```

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

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/disable_verify_certificate/): complete subsection reference.

- [disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/disable_verify_hostname/): complete subsection reference.

- [enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/enable_verify_certificate/): complete subsection reference.

- [enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/enable_verify_hostname/): complete subsection reference.

- [mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_disabled/): complete subsection reference.

- [mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/): complete subsection reference.

- [no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/no_ca/): complete subsection reference.

<a id="schema-qradar_receiver--use_tls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
