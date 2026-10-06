---
page_title: "webhook.http_config"
subcategory: ""
description: "Configuration for HTTP endpoint."
xcsh_docs: {"aliases": ["webhook http config"], "body_bytes": 2730, "body_sha256": "sha256:d3c8f51554a0f362ea731efab57c44b4a40f6906bd9c91bf783941263cfd57cf", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "sections": [{"aliases": ["webhook http config auth token"], "anchor": "section", "description": "Authentication Token for access.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "auth_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config basic auth"], "anchor": "section", "description": "Authorization parameters to access HTPP alert Receiver Endpoint.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:basic_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "basic_auth"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config client cert obj"], "anchor": "section", "description": "Configuration for client certificate.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config enable http2"], "anchor": "schema-webhook--http_config--enable_http2", "description": "Configure to use HTTP2 protocol.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "enable_http2"], "syntax": "attribute", "type": "bool"}, {"aliases": ["webhook http config follow redirects"], "anchor": "schema-webhook--http_config--follow_redirects", "description": "Configure whether HTTP requests follow HTTP 3xx redirects.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "follow_redirects"], "syntax": "attribute", "type": "bool"}, {"aliases": ["webhook http config no authorization"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_authorization", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "no_authorization"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config use tls"], "anchor": "section", "description": "Configures the token request's TLS settings.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "use_tls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configuration for HTTP endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- webhook.http_config

<a id="section"></a>

Type: `"single"`. Computed.

HTTP Configuration. Configuration for HTTP endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_choice": "[\"auth_token\",\"basic_auth\",\"client_cert_obj\",\"no_authorization\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

## Direct properties

- [auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/auth_token/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/basic_auth/): complete subsection reference.

- [client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/): complete subsection reference.

<a id="schema-webhook--http_config--enable_http2"></a>

### enable_http2 property

Type: `"bool"`. Computed.

Enable HTTP2. Configure to use HTTP2 protocol.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-webhook--http_config--follow_redirects"></a>

### follow_redirects property

Type: `"bool"`. Computed.

Configure whether HTTP requests follow HTTP 3xx redirects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_authorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/no_authorization/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/no_tls/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/): complete subsection reference.
