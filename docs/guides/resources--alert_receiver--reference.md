---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 28430, "body_sha256": "sha256:a1a8d894cbee61c4422c5c5ff71894451b1566d79f40bfad832464ccb605ce28", "canonical_id": "xcsh-docs:resources:alert_receiver:reference", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:email", "xcsh-docs:resources:alert_receiver:properties:opsgenie", "xcsh-docs:resources:alert_receiver:properties:pagerduty", "xcsh-docs:resources:alert_receiver:properties:slack", "xcsh-docs:resources:alert_receiver:properties:sms", "xcsh-docs:resources:alert_receiver:properties:timeouts", "xcsh-docs:resources:alert_receiver:properties:webhook"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:reference", "parent_id": "xcsh-docs:resources:alert_receiver:fundamentals", "path": "docs/guides/resources--alert_receiver--reference.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [email](resources--alert_receiver--properties--email.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Alert Receiver. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Alert Receiver is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [opsgenie](resources--alert_receiver--properties--opsgenie.md): complete subsection reference.

- [pagerduty](resources--alert_receiver--properties--pagerduty.md): complete subsection reference.

- [slack](resources--alert_receiver--properties--slack.md): complete subsection reference.

- [sms](resources--alert_receiver--properties--sms.md): complete subsection reference.

- [timeouts](resources--alert_receiver--properties--timeouts.md): complete subsection reference.

- [webhook](resources--alert_receiver--properties--webhook.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--alert_receiver--reference.md#schema-annotations) |
| `description` | [description](resources--alert_receiver--reference.md#schema-description) |
| `disable` | [disable](resources--alert_receiver--reference.md#schema-disable) |
| `email` | [email](resources--alert_receiver--properties--email.md#section) |
| `email.email` | [email.email](resources--alert_receiver--properties--email.md#schema-email--email) |
| `id` | [id](resources--alert_receiver--reference.md#schema-id) |
| `labels` | [labels](resources--alert_receiver--reference.md#schema-labels) |
| `name` | [name](resources--alert_receiver--reference.md#schema-name) |
| `namespace` | [namespace](resources--alert_receiver--reference.md#schema-namespace) |
| `opsgenie` | [opsgenie](resources--alert_receiver--properties--opsgenie.md#section) |
| `opsgenie.api_key` | [opsgenie.api_key](resources--alert_receiver--properties--opsgenie--api_key.md#section) |
| `opsgenie.api_key.blindfold_secret_info` | [opsgenie.api_key.blindfold_secret_info](resources--alert_receiver--properties--opsgenie--api_key--blindfold_secret_info.md#section) |
| `opsgenie.api_key.blindfold_secret_info.decryption_provider` | [opsgenie.api_key.blindfold_secret_info.decryption_provider](resources--alert_receiver--properties--opsgenie--api_key--blindfold_secret_info.md#schema-opsgenie--api_key--blindfold_secret_info--decryption_provider) |
| `opsgenie.api_key.blindfold_secret_info.location` | [opsgenie.api_key.blindfold_secret_info.location](resources--alert_receiver--properties--opsgenie--api_key--blindfold_secret_info.md#schema-opsgenie--api_key--blindfold_secret_info--location) |
| `opsgenie.api_key.blindfold_secret_info.store_provider` | [opsgenie.api_key.blindfold_secret_info.store_provider](resources--alert_receiver--properties--opsgenie--api_key--blindfold_secret_info.md#schema-opsgenie--api_key--blindfold_secret_info--store_provider) |
| `opsgenie.api_key.clear_secret_info` | [opsgenie.api_key.clear_secret_info](resources--alert_receiver--properties--opsgenie--api_key--clear_secret_info.md#section) |
| `opsgenie.api_key.clear_secret_info.provider_ref` | [opsgenie.api_key.clear_secret_info.provider_ref](resources--alert_receiver--properties--opsgenie--api_key--clear_secret_info.md#schema-opsgenie--api_key--clear_secret_info--provider_ref) |
| `opsgenie.api_key.clear_secret_info.url` | [opsgenie.api_key.clear_secret_info.url](resources--alert_receiver--properties--opsgenie--api_key--clear_secret_info.md#schema-opsgenie--api_key--clear_secret_info--url) |
| `opsgenie.url` | [opsgenie.url](resources--alert_receiver--properties--opsgenie.md#schema-opsgenie--url) |
| `pagerduty` | [pagerduty](resources--alert_receiver--properties--pagerduty.md#section) |
| `pagerduty.routing_key` | [pagerduty.routing_key](resources--alert_receiver--properties--pagerduty--routing_key.md#section) |
| `pagerduty.routing_key.blindfold_secret_info` | [pagerduty.routing_key.blindfold_secret_info](resources--alert_receiver--properties--pagerduty--routing_key--blindfold_secret_info.md#section) |
| `pagerduty.routing_key.blindfold_secret_info.decryption_provider` | [pagerduty.routing_key.blindfold_secret_info.decryption_provider](resources--alert_receiver--properties--pagerduty--routing_key--blindfold_secret_info.md#schema-pagerduty--routing_key--blindfold_secret_info--decryption_provider) |
| `pagerduty.routing_key.blindfold_secret_info.location` | [pagerduty.routing_key.blindfold_secret_info.location](resources--alert_receiver--properties--pagerduty--routing_key--blindfold_secret_info.md#schema-pagerduty--routing_key--blindfold_secret_info--location) |
| `pagerduty.routing_key.blindfold_secret_info.store_provider` | [pagerduty.routing_key.blindfold_secret_info.store_provider](resources--alert_receiver--properties--pagerduty--routing_key--blindfold_secret_info.md#schema-pagerduty--routing_key--blindfold_secret_info--store_provider) |
| `pagerduty.routing_key.clear_secret_info` | [pagerduty.routing_key.clear_secret_info](resources--alert_receiver--properties--pagerduty--routing_key--clear_secret_info.md#section) |
| `pagerduty.routing_key.clear_secret_info.provider_ref` | [pagerduty.routing_key.clear_secret_info.provider_ref](resources--alert_receiver--properties--pagerduty--routing_key--clear_secret_info.md#schema-pagerduty--routing_key--clear_secret_info--provider_ref) |
| `pagerduty.routing_key.clear_secret_info.url` | [pagerduty.routing_key.clear_secret_info.url](resources--alert_receiver--properties--pagerduty--routing_key--clear_secret_info.md#schema-pagerduty--routing_key--clear_secret_info--url) |
| `pagerduty.url` | [pagerduty.url](resources--alert_receiver--properties--pagerduty.md#schema-pagerduty--url) |
| `slack` | [slack](resources--alert_receiver--properties--slack.md#section) |
| `slack.channel` | [slack.channel](resources--alert_receiver--properties--slack.md#schema-slack--channel) |
| `slack.url` | [slack.url](resources--alert_receiver--properties--slack--url.md#section) |
| `slack.url.blindfold_secret_info` | [slack.url.blindfold_secret_info](resources--alert_receiver--properties--slack--url--blindfold_secret_info.md#section) |
| `slack.url.blindfold_secret_info.decryption_provider` | [slack.url.blindfold_secret_info.decryption_provider](resources--alert_receiver--properties--slack--url--blindfold_secret_info.md#schema-slack--url--blindfold_secret_info--decryption_provider) |
| `slack.url.blindfold_secret_info.location` | [slack.url.blindfold_secret_info.location](resources--alert_receiver--properties--slack--url--blindfold_secret_info.md#schema-slack--url--blindfold_secret_info--location) |
| `slack.url.blindfold_secret_info.store_provider` | [slack.url.blindfold_secret_info.store_provider](resources--alert_receiver--properties--slack--url--blindfold_secret_info.md#schema-slack--url--blindfold_secret_info--store_provider) |
| `slack.url.clear_secret_info` | [slack.url.clear_secret_info](resources--alert_receiver--properties--slack--url--clear_secret_info.md#section) |
| `slack.url.clear_secret_info.provider_ref` | [slack.url.clear_secret_info.provider_ref](resources--alert_receiver--properties--slack--url--clear_secret_info.md#schema-slack--url--clear_secret_info--provider_ref) |
| `slack.url.clear_secret_info.url` | [slack.url.clear_secret_info.url](resources--alert_receiver--properties--slack--url--clear_secret_info.md#schema-slack--url--clear_secret_info--url) |
| `sms` | [sms](resources--alert_receiver--properties--sms.md#section) |
| `sms.contact_number` | [sms.contact_number](resources--alert_receiver--properties--sms.md#schema-sms--contact_number) |
| `timeouts` | [timeouts](resources--alert_receiver--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--alert_receiver--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--alert_receiver--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--alert_receiver--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--alert_receiver--properties--timeouts.md#schema-timeouts--update) |
| `webhook` | [webhook](resources--alert_receiver--properties--webhook.md#section) |
| `webhook.http_config` | [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md#section) |
| `webhook.http_config.auth_token` | [webhook.http_config.auth_token](resources--alert_receiver--properties--webhook--http_config--auth_token.md#section) |
| `webhook.http_config.auth_token.token` | [webhook.http_config.auth_token.token](resources--alert_receiver--properties--webhook--http_config--auth_token--token.md#section) |
| `webhook.http_config.auth_token.token.blindfold_secret_info` | [webhook.http_config.auth_token.token.blindfold_secret_info](resources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md#section) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider](resources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md#schema-webhook--http_config--auth_token--token--blindfold_secret_info--decryption_provider) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.location` | [webhook.http_config.auth_token.token.blindfold_secret_info.location](resources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md#schema-webhook--http_config--auth_token--token--blindfold_secret_info--location) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.store_provider](resources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md#schema-webhook--http_config--auth_token--token--blindfold_secret_info--store_provider) |
| `webhook.http_config.auth_token.token.clear_secret_info` | [webhook.http_config.auth_token.token.clear_secret_info](resources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md#section) |
| `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` | [webhook.http_config.auth_token.token.clear_secret_info.provider_ref](resources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md#schema-webhook--http_config--auth_token--token--clear_secret_info--provider_ref) |
| `webhook.http_config.auth_token.token.clear_secret_info.url` | [webhook.http_config.auth_token.token.clear_secret_info.url](resources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md#schema-webhook--http_config--auth_token--token--clear_secret_info--url) |
| `webhook.http_config.basic_auth` | [webhook.http_config.basic_auth](resources--alert_receiver--properties--webhook--http_config--basic_auth.md#section) |
| `webhook.http_config.basic_auth.password` | [webhook.http_config.basic_auth.password](resources--alert_receiver--properties--webhook--http_config--basic_auth--password.md#section) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info` | [webhook.http_config.basic_auth.password.blindfold_secret_info](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--blindfold_secret_info.md#section) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--blindfold_secret_info.md#schema-webhook--http_config--basic_auth--password--blindfold_secret_info--decryption_provider) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.location` | [webhook.http_config.basic_auth.password.blindfold_secret_info.location](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--blindfold_secret_info.md#schema-webhook--http_config--basic_auth--password--blindfold_secret_info--location) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--blindfold_secret_info.md#schema-webhook--http_config--basic_auth--password--blindfold_secret_info--store_provider) |
| `webhook.http_config.basic_auth.password.clear_secret_info` | [webhook.http_config.basic_auth.password.clear_secret_info](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--clear_secret_info.md#section) |
| `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` | [webhook.http_config.basic_auth.password.clear_secret_info.provider_ref](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--clear_secret_info.md#schema-webhook--http_config--basic_auth--password--clear_secret_info--provider_ref) |
| `webhook.http_config.basic_auth.password.clear_secret_info.url` | [webhook.http_config.basic_auth.password.clear_secret_info.url](resources--alert_receiver--properties--webhook--http_config--basic_auth--password--clear_secret_info.md#schema-webhook--http_config--basic_auth--password--clear_secret_info--url) |
| `webhook.http_config.basic_auth.user_name` | [webhook.http_config.basic_auth.user_name](resources--alert_receiver--properties--webhook--http_config--basic_auth.md#schema-webhook--http_config--basic_auth--user_name) |
| `webhook.http_config.client_cert_obj` | [webhook.http_config.client_cert_obj](resources--alert_receiver--properties--webhook--http_config--client_cert_obj.md#section) |
| `webhook.http_config.client_cert_obj.use_tls_obj` | [webhook.http_config.client_cert_obj.use_tls_obj](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md#section) |
| `webhook.http_config.client_cert_obj.use_tls_obj.kind` | [webhook.http_config.client_cert_obj.use_tls_obj.kind](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md#schema-webhook--http_config--client_cert_obj--use_tls_obj--kind) |
| `webhook.http_config.client_cert_obj.use_tls_obj.name` | [webhook.http_config.client_cert_obj.use_tls_obj.name](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md#schema-webhook--http_config--client_cert_obj--use_tls_obj--name) |
| `webhook.http_config.client_cert_obj.use_tls_obj.namespace` | [webhook.http_config.client_cert_obj.use_tls_obj.namespace](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md#schema-webhook--http_config--client_cert_obj--use_tls_obj--namespace) |
| `webhook.http_config.client_cert_obj.use_tls_obj.tenant` | [webhook.http_config.client_cert_obj.use_tls_obj.tenant](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md#schema-webhook--http_config--client_cert_obj--use_tls_obj--tenant) |
| `webhook.http_config.client_cert_obj.use_tls_obj.uid` | [webhook.http_config.client_cert_obj.use_tls_obj.uid](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md#schema-webhook--http_config--client_cert_obj--use_tls_obj--uid) |
| `webhook.http_config.enable_http2` | [webhook.http_config.enable_http2](resources--alert_receiver--properties--webhook--http_config.md#schema-webhook--http_config--enable_http2) |
| `webhook.http_config.follow_redirects` | [webhook.http_config.follow_redirects](resources--alert_receiver--properties--webhook--http_config.md#schema-webhook--http_config--follow_redirects) |
| `webhook.http_config.no_authorization` | [webhook.http_config.no_authorization](resources--alert_receiver--properties--webhook--http_config--no_authorization.md#section) |
| `webhook.http_config.no_tls` | [webhook.http_config.no_tls](resources--alert_receiver--properties--webhook--http_config--no_tls.md#section) |
| `webhook.http_config.use_tls` | [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md#section) |
| `webhook.http_config.use_tls.disable_sni` | [webhook.http_config.use_tls.disable_sni](resources--alert_receiver--properties--webhook--http_config--use_tls--disable_sni.md#section) |
| `webhook.http_config.use_tls.max_version` | [webhook.http_config.use_tls.max_version](resources--alert_receiver--properties--webhook--http_config--use_tls.md#schema-webhook--http_config--use_tls--max_version) |
| `webhook.http_config.use_tls.min_version` | [webhook.http_config.use_tls.min_version](resources--alert_receiver--properties--webhook--http_config--use_tls.md#schema-webhook--http_config--use_tls--min_version) |
| `webhook.http_config.use_tls.sni` | [webhook.http_config.use_tls.sni](resources--alert_receiver--properties--webhook--http_config--use_tls.md#schema-webhook--http_config--use_tls--sni) |
| `webhook.http_config.use_tls.use_server_verification` | [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification.md#section) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj.md#section) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md#section) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--kind) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--name) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--namespace) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--tenant) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--uid) |
| `webhook.http_config.use_tls.volterra_trusted_ca` | [webhook.http_config.use_tls.volterra_trusted_ca](resources--alert_receiver--properties--webhook--http_config--use_tls--volterra_trusted_ca.md#section) |
| `webhook.url` | [webhook.url](resources--alert_receiver--properties--webhook--url.md#section) |
| `webhook.url.blindfold_secret_info` | [webhook.url.blindfold_secret_info](resources--alert_receiver--properties--webhook--url--blindfold_secret_info.md#section) |
| `webhook.url.blindfold_secret_info.decryption_provider` | [webhook.url.blindfold_secret_info.decryption_provider](resources--alert_receiver--properties--webhook--url--blindfold_secret_info.md#schema-webhook--url--blindfold_secret_info--decryption_provider) |
| `webhook.url.blindfold_secret_info.location` | [webhook.url.blindfold_secret_info.location](resources--alert_receiver--properties--webhook--url--blindfold_secret_info.md#schema-webhook--url--blindfold_secret_info--location) |
| `webhook.url.blindfold_secret_info.store_provider` | [webhook.url.blindfold_secret_info.store_provider](resources--alert_receiver--properties--webhook--url--blindfold_secret_info.md#schema-webhook--url--blindfold_secret_info--store_provider) |
| `webhook.url.clear_secret_info` | [webhook.url.clear_secret_info](resources--alert_receiver--properties--webhook--url--clear_secret_info.md#section) |
| `webhook.url.clear_secret_info.provider_ref` | [webhook.url.clear_secret_info.provider_ref](resources--alert_receiver--properties--webhook--url--clear_secret_info.md#schema-webhook--url--clear_secret_info--provider_ref) |
| `webhook.url.clear_secret_info.url` | [webhook.url.clear_secret_info.url](resources--alert_receiver--properties--webhook--url--clear_secret_info.md#schema-webhook--url--clear_secret_info--url) |

## Next pages

- [email](resources--alert_receiver--properties--email.md)
- [opsgenie](resources--alert_receiver--properties--opsgenie.md)
- [pagerduty](resources--alert_receiver--properties--pagerduty.md)
- [slack](resources--alert_receiver--properties--slack.md)
- [sms](resources--alert_receiver--properties--sms.md)
- [timeouts](resources--alert_receiver--properties--timeouts.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
