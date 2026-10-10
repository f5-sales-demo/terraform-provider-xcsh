---
page_title: "webhook.http_config.use_tls"
subcategory: ""
description: "Configures the token request's TLS settings."
xcsh_docs: {"aliases": ["webhook http config use tls"], "body_bytes": 5550, "body_sha256": "sha256:c6cdf983c0c194b5621cbcfc61859fc302987201b92a52bbc71e91349ca45704", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/use_tls/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2233200131202023-2133021101000112-3033100330210313-1022303021210003-0333000013320003-0130013021000100-1022220003213321-3223012110031213", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-webhook--http_config--use_tls--sni", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["webhook http config use tls disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config use tls max version"], "anchor": "schema-webhook--http_config--use_tls--max_version", "description": "TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["TLS_AUTO", "TLSv1_0", "TLSv1_1", "TLSv1_2", "TLSv1_3"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "max_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["webhook http config use tls min version"], "anchor": "schema-webhook--http_config--use_tls--min_version", "description": "TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["TLS_AUTO", "TLSv1_0", "TLSv1_1", "TLSv1_2", "TLSv1_3"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "min_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["webhook http config use tls sni"], "anchor": "schema-webhook--http_config--use_tls--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["webhook http config use tls use server verification"], "anchor": "section", "description": "Upstream TLS Validation Context.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification"], "syntax": "block", "type": "object"}, {"aliases": ["webhook http config use tls volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configures the token request's TLS settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- webhook.http_config.use_tls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configures the token request's TLS settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-server_validation_choice": "[\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/disable_sni/): complete subsection reference.

<a id="schema-webhook--http_config--use_tls--max_version"></a>

### max_version property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-webhook--http_config--use_tls--min_version"></a>

### min_version property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-webhook--http_config--use_tls--sni"></a>

### sni property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/volterra_trusted_ca/): complete subsection reference.
