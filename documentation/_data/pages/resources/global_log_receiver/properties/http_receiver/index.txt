---
page_title: "http_receiver"
subcategory: ""
description: "Configuration for HTTP endpoint."
xcsh_docs: {"aliases": ["http receiver"], "body_bytes": 4931, "body_sha256": "sha256:0e51760bbb66413c85a60e8de95f737593e834630f79e6b5e46006a5cb0920dd", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/http_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_none,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_none,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-http_receiver--uri", "enforcement": "provider-schema", "group": "http_receiver:RequiredObjectAttributes:uri", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver"], "schema_version": 1, "sections": [{"aliases": ["http receiver auth basic"], "anchor": "section", "description": "Authentication parameters to access HTPP Log Receiver Endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "auth_basic"], "syntax": "block", "type": "object"}, {"aliases": ["http receiver auth none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "auth_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["http receiver auth token"], "anchor": "section", "description": "Authentication Token for access.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "auth_token"], "syntax": "block", "type": "object"}, {"aliases": ["http receiver batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_receiver--batch--max_bytes", "enforcement": "provider-schema", "group": "http_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "type": "conflicts"}, {"anchor": "schema-http_receiver--batch--max_events", "enforcement": "provider-schema", "group": "http_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "type": "conflicts"}, {"anchor": "schema-http_receiver--batch--timeout_seconds", "enforcement": "provider-schema", "group": "http_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch:max_bytes_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch:max_events_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch:timeout_seconds_default", "type": "conflicts"}], "schema_path": ["http_receiver", "batch"], "syntax": "block", "type": "object"}, {"aliases": ["http receiver compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_none", "type": "conflicts"}], "schema_path": ["http_receiver", "compression"], "syntax": "block", "type": "object"}, {"aliases": ["http receiver no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:no_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["http receiver uri"], "anchor": "schema-http_receiver--uri", "description": "HTTP URI is the URI of the HTTP endpoint to send logs to,.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["http receiver use tls"], "anchor": "section", "description": "TLS Parameters for client connection to the endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_receiver--use_tls--trusted_ca_url", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:no_ca", "type": "conflicts"}], "schema_path": ["http_receiver", "use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration for HTTP endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- http_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http receiver.

Upstream description:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri"),
  validators.ConflictingObjectAttributes("auth_basic",
    "auth_none"),
  validators.ConflictingObjectAttributes("auth_basic",
    "auth_token"),
  validators.ConflictingObjectAttributes("auth_none",
    "auth_token"),
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
  "x-ves-oneof-field-auth_choice": "[\"auth_basic\",\"auth_none\",\"auth_token\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
http_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auth_basic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/): complete subsection reference.

- [auth_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_none/): complete subsection reference.

- [auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/): complete subsection reference.

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/no_tls/): complete subsection reference.

<a id="schema-http_receiver--uri"></a>

### uri property

Type: `"string"`. Optional.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/): complete subsection reference.

## Next pages

- [http_receiver.auth_basic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/)
- [http_receiver.auth_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_none/)
- [http_receiver.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/)
- [http_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/)
- [http_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/)
- [http_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/no_tls/)
- [http_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
