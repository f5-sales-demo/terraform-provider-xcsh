---
page_title: "webhook.http_config"
subcategory: ""
description: "Configuration for HTTP endpoint."
xcsh_docs: {"aliases": ["webhook http config"], "body_bytes": 4073, "body_sha256": "sha256:5109f1cae1e730ff8019eff8b83bdbfa79ad954959c9885a97d638c9cd4af5bc", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2331011100033223-2223121001120230-0320022321013132-0021122120120022-0012123302320121-2010100323020230-3333230202010122-0322013002200012", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "sections": [{"aliases": ["auth token", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Authentication Token for access.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "auth_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["basic auth"], "anchor": "section", "description": "Authorization parameters to access HTPP alert Receiver Endpoint.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:basic_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "basic_auth"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "client cert obj", "existing certificates", "tls certificates"], "anchor": "section", "description": "Configuration for client certificate.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable http2"], "anchor": "schema-webhook--http_config--enable_http2", "description": "Configure to use HTTP2 protocol.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "enable_http2"], "syntax": "attribute", "type": "bool"}, {"aliases": ["follow redirects"], "anchor": "schema-webhook--http_config--follow_redirects", "description": "Configure whether HTTP requests follow HTTP 3xx redirects.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "follow_redirects"], "syntax": "attribute", "type": "bool"}, {"aliases": ["no authorization"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_authorization", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "no_authorization"], "syntax": "attribute", "type": "object"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls"], "anchor": "section", "description": "Configures the token request's TLS settings.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "use_tls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for HTTP endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Configure to use HTTP2 protocol.

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

## Next pages

- [webhook.http_config.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/auth_token/)
- [webhook.http_config.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/basic_auth/)
- [webhook.http_config.client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/)
- [webhook.http_config.no_authorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/no_authorization/)
- [webhook.http_config.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/no_tls/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
