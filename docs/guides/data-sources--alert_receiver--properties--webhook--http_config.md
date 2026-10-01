---
page_title: "webhook.http_config"
subcategory: ""
description: "webhook.http_config for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3228, "body_sha256": "sha256:deae41912212e5c4603ea93ab2b5a7cca41fcb294d70672a158fe8356514d2e9", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "path": "docs/guides/data-sources--alert_receiver--properties--webhook--http_config.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
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

- [auth_token](data-sources--alert_receiver--properties--webhook--http_config--auth_token.md): complete subsection reference.

- [basic_auth](data-sources--alert_receiver--properties--webhook--http_config--basic_auth.md): complete subsection reference.

- [client_cert_obj](data-sources--alert_receiver--properties--webhook--http_config--client_cert_obj.md): complete subsection reference.

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

- [no_authorization](data-sources--alert_receiver--properties--webhook--http_config--no_authorization.md): complete subsection reference.

- [no_tls](data-sources--alert_receiver--properties--webhook--http_config--no_tls.md): complete subsection reference.

- [use_tls](data-sources--alert_receiver--properties--webhook--http_config--use_tls.md): complete subsection reference.

## Next pages

- [webhook.http_config.auth_token](data-sources--alert_receiver--properties--webhook--http_config--auth_token.md)
- [webhook.http_config.basic_auth](data-sources--alert_receiver--properties--webhook--http_config--basic_auth.md)
- [webhook.http_config.client_cert_obj](data-sources--alert_receiver--properties--webhook--http_config--client_cert_obj.md)
- [webhook.http_config.no_authorization](data-sources--alert_receiver--properties--webhook--http_config--no_authorization.md)
- [webhook.http_config.no_tls](data-sources--alert_receiver--properties--webhook--http_config--no_tls.md)
- [webhook.http_config.use_tls](data-sources--alert_receiver--properties--webhook--http_config--use_tls.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
