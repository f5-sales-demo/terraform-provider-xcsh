---
page_title: "http_receiver"
subcategory: ""
description: "Configuration for HTTP endpoint."
xcsh_docs: {"aliases": ["http receiver"], "body_bytes": 4295, "body_sha256": "sha256:52208d69856de9148f927a7de6c49b2199319f74205f91d84e16112620e6e26a", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_basic", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_none", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:no_tls", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/http_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver"], "schema_version": 1, "sections": [{"aliases": ["auth basic", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Authentication parameters to access HTPP Log Receiver Endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_basic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "auth_basic"], "syntax": "attribute", "type": "object"}, {"aliases": ["auth none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "auth_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["auth token", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Authentication Token for access.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "auth_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "batch"], "syntax": "attribute", "type": "object"}, {"aliases": ["compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:compression", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "compression"], "syntax": "attribute", "type": "object"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:no_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["uri"], "anchor": "schema-http_receiver--uri", "description": "HTTP URI is the URI of the HTTP endpoint to send logs to,.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "uri"], "syntax": "attribute", "type": "string"}, {"aliases": ["use tls"], "anchor": "section", "description": "TLS Parameters for client connection to the endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "use_tls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for HTTP endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- http_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http receiver.

Upstream description:

Configuration for HTTP endpoint.

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

## Direct properties

- [auth_basic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_basic/): complete subsection reference.

- [auth_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_none/): complete subsection reference.

- [auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/): complete subsection reference.

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/compression/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/no_tls/): complete subsection reference.

<a id="schema-http_receiver--uri"></a>

### uri property

Type: `"string"`. Computed.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

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

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/): complete subsection reference.

## Next pages

- [http_receiver.auth_basic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_basic/)
- [http_receiver.auth_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_none/)
- [http_receiver.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/)
- [http_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/batch/)
- [http_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/compression/)
- [http_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/no_tls/)
- [http_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
