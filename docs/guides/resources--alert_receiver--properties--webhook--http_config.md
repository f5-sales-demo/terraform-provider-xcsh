---
page_title: "webhook.http_config"
subcategory: ""
description: "webhook.http_config for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3937, "body_sha256": "sha256:e0ff436470a6258caa98c037095e7d8e82f82d8dd963d2aa9a3c1bd3aa5c7683", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- webhook.http_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP Configuration. Configuration for HTTP endpoint.

Upstream description:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
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

- [auth_token](resources--alert_receiver--properties--webhook--http_config--auth_token.md): complete subsection reference.

- [basic_auth](resources--alert_receiver--properties--webhook--http_config--basic_auth.md): complete subsection reference.

- [client_cert_obj](resources--alert_receiver--properties--webhook--http_config--client_cert_obj.md): complete subsection reference.

<a id="schema-webhook--http_config--enable_http2"></a>

### enable_http2 property

Type: `"bool"`. Optional.

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

- [no_authorization](resources--alert_receiver--properties--webhook--http_config--no_authorization.md): complete subsection reference.

- [no_tls](resources--alert_receiver--properties--webhook--http_config--no_tls.md): complete subsection reference.

- [use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md): complete subsection reference.

## Next pages

- [webhook.http_config.auth_token](resources--alert_receiver--properties--webhook--http_config--auth_token.md)
- [webhook.http_config.basic_auth](resources--alert_receiver--properties--webhook--http_config--basic_auth.md)
- [webhook.http_config.client_cert_obj](resources--alert_receiver--properties--webhook--http_config--client_cert_obj.md)
- [webhook.http_config.no_authorization](resources--alert_receiver--properties--webhook--http_config--no_authorization.md)
- [webhook.http_config.no_tls](resources--alert_receiver--properties--webhook--http_config--no_tls.md)
- [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
