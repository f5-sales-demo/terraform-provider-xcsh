---
page_title: "webhook.http_config"
subcategory: ""
description: "webhook.http_config for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4782, "body_sha256": "sha256:52c555b072745048e34be3cb0609f191db1b8d5d695c32a05317a3927488c7b8", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_tls", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/index.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["webhook", "http_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/): complete subsection reference.

- [client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/): complete subsection reference.

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

- [no_authorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_authorization/): complete subsection reference.

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_tls/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/): complete subsection reference.

## Next pages

- [webhook.http_config.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/)
- [webhook.http_config.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/)
- [webhook.http_config.client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/)
- [webhook.http_config.no_authorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_authorization/)
- [webhook.http_config.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_tls/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
