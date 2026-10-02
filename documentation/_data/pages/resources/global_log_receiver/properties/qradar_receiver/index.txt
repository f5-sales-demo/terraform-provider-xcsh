---
page_title: "qradar_receiver"
subcategory: ""
description: "Configuration for IBM QRadar endpoint."
xcsh_docs: {"aliases": ["qradar receiver"], "body_bytes": 3730, "body_sha256": "sha256:ceb3f6aa7902abdbaa4f08e9dae8bae10f73b709a19a6e168a35ab1aaac0e51e", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/qradar_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-qradar_receiver--uri", "enforcement": "provider-schema", "group": "qradar_receiver:RequiredObjectAttributes:uri", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["qradar_receiver"], "schema_version": 1, "sections": [{"aliases": ["batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-qradar_receiver--batch--max_bytes", "enforcement": "provider-schema", "group": "qradar_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch", "type": "conflicts"}, {"anchor": "schema-qradar_receiver--batch--max_events", "enforcement": "provider-schema", "group": "qradar_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch", "type": "conflicts"}, {"anchor": "schema-qradar_receiver--batch--timeout_seconds", "enforcement": "provider-schema", "group": "qradar_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch:max_bytes_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch:max_events_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:batch:timeout_seconds_default", "type": "conflicts"}], "schema_path": ["qradar_receiver", "batch"], "syntax": "block", "type": "object"}, {"aliases": ["compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_none", "type": "conflicts"}], "schema_path": ["qradar_receiver", "compression"], "syntax": "block", "type": "object"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:no_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["uri"], "anchor": "schema-qradar_receiver--uri", "description": "Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["use tls"], "anchor": "section", "description": "TLS Parameters for client connection to the endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-qradar_receiver--use_tls--trusted_ca_url", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:disable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:enable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:no_ca", "type": "conflicts"}], "schema_path": ["qradar_receiver", "use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/qradar_receiver/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for IBM QRadar endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- qradar_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for qradar receiver.

Upstream description:

Configuration for IBM QRadar endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
qradar_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/compression/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/no_tls/): complete subsection reference.

<a id="schema-qradar_receiver--uri"></a>

### uri property

Type: `"string"`. Optional.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/): complete subsection reference.

## Next pages

- [qradar_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/)
- [qradar_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/compression/)
- [qradar_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/no_tls/)
- [qradar_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
