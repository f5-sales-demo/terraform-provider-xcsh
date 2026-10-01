---
page_title: "webhook.http_config"
subcategory: ""
description: "webhook.http_config for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4073, "body_sha256": "sha256:5109f1cae1e730ff8019eff8b83bdbfa79ad954959c9885a97d638c9cd4af5bc", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/index.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
