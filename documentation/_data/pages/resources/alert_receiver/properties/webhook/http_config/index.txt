---
page_title: "webhook.http_config"
subcategory: ""
description: "Configuration for HTTP endpoint."
xcsh_docs: {"aliases": ["webhook http config"], "body_bytes": 3493, "body_sha256": "sha256:79a5b58ae7d8d83f56ce321ec53555bd0746dedfad4700ca023846a610e70432", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1030200301112321-0230212112121020-3303323333001300-0101003301133101-2123123032220110-0012111023302011-2330222013130102-3103032011103022", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:auth_token,basic_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:auth_token,client_cert_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:auth_token,no_authorization", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:auth_token,basic_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:basic_auth,client_cert_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:basic_auth,no_authorization", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:auth_token,client_cert_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:basic_auth,client_cert_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:client_cert_obj,no_authorization", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:auth_token,no_authorization", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:basic_auth,no_authorization", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:client_cert_obj,no_authorization", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "sections": [{"aliases": ["webhook http config auth token"], "anchor": "section", "description": "Authentication Token for access.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "auth_token"], "syntax": "block", "type": "object"}, {"aliases": ["webhook http config basic auth"], "anchor": "section", "description": "Authorization parameters to access HTPP alert Receiver Endpoint.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-webhook--http_config--basic_auth--user_name", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth:RequiredObjectAttributes:user_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "type": "requires"}], "schema_path": ["webhook", "http_config", "basic_auth"], "syntax": "block", "type": "object"}, {"aliases": ["webhook http config client cert obj"], "anchor": "section", "description": "Configuration for client certificate.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj"], "syntax": "block", "type": "object"}, {"aliases": ["webhook http config enable http2"], "anchor": "schema-webhook--http_config--enable_http2", "description": "Configure to use HTTP2 protocol.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "enable_http2"], "syntax": "attribute", "type": "bool"}, {"aliases": ["webhook http config follow redirects"], "anchor": "schema-webhook--http_config--follow_redirects", "description": "Configure whether HTTP requests follow HTTP 3xx redirects.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "follow_redirects"], "syntax": "attribute", "type": "bool"}, {"aliases": ["webhook http config no authorization"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "no_authorization"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook http config use tls"], "anchor": "section", "description": "Configures the token request's TLS settings.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-webhook--http_config--use_tls--sni", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca", "type": "conflicts"}], "schema_path": ["webhook", "http_config", "use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration for HTTP endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- webhook.http_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP Configuration. Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_token",
    "basic_auth"),
  validators.ConflictingObjectAttributes("auth_token",
    "client_cert_obj"),
  validators.ConflictingObjectAttributes("auth_token",
    "no_authorization"),
  validators.ConflictingObjectAttributes("basic_auth",
    "client_cert_obj"),
  validators.ConflictingObjectAttributes("basic_auth",
    "no_authorization"),
  validators.ConflictingObjectAttributes("client_cert_obj",
    "no_authorization"),
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
  "x-ves-oneof-field-auth_choice": "[\"auth_token\",\"basic_auth\",\"client_cert_obj\",\"no_authorization\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
http_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/): complete subsection reference.

- [client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/): complete subsection reference.

<a id="schema-webhook--http_config--enable_http2"></a>

### enable_http2 property

Type: `"bool"`. Optional.

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

Type: `"bool"`. Optional.

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

- [no_authorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_authorization/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_tls/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/): complete subsection reference.
