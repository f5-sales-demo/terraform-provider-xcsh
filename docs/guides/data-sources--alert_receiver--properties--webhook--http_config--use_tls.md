---
page_title: "webhook.http_config.use_tls"
subcategory: ""
description: "webhook.http_config.use_tls for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4535, "body_sha256": "sha256:6cb6c3a4de0c697dfbef8805fd179a1244d98fea18eb5162bb02991e81dfd5ce", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "path": "docs/guides/data-sources--alert_receiver--properties--webhook--http_config--use_tls.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "use_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/use_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.use_tls for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# webhook.http_config.use_tls

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- webhook.http_config.use_tls

<a id="section"></a>

Type: `"single"`. Computed.

Configures the token request's TLS settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-server_validation_choice": "[\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\"]"
}
```

## Direct properties

- [disable_sni](data-sources--alert_receiver--properties--webhook--http_config--use_tls--disable_sni.md): complete subsection reference.

<a id="schema-webhook--http_config--use_tls--max_version"></a>

### max_version property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-webhook--http_config--use_tls--min_version"></a>

### min_version property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-webhook--http_config--use_tls--sni"></a>

### sni property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_server_verification](data-sources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification.md): complete subsection reference.

- [volterra_trusted_ca](data-sources--alert_receiver--properties--webhook--http_config--use_tls--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [webhook.http_config.use_tls.disable_sni](data-sources--alert_receiver--properties--webhook--http_config--use_tls--disable_sni.md)
- [webhook.http_config.use_tls.use_server_verification](data-sources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification.md)
- [webhook.http_config.use_tls.volterra_trusted_ca](data-sources--alert_receiver--properties--webhook--http_config--use_tls--volterra_trusted_ca.md)
- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
