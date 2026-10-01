---
page_title: "http_receiver"
subcategory: ""
description: "http_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4023, "body_sha256": "sha256:80730618a89331d2cb0538450d2b689d5ba19e35645f2e9b9c51435c635338ee", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--http_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- http_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http receiver.

Upstream description:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri"),
  validators.ConflictingObjectAttributes("auth_basic",
    "auth_none"),
  validators.ConflictingObjectAttributes("auth_basic",
    "auth_token"),
  validators.ConflictingObjectAttributes("auth_none",
    "auth_token"),
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
  "x-ves-oneof-field-auth_choice": "[\"auth_basic\",\"auth_none\",\"auth_token\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
http_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auth_basic](resources--global_log_receiver--properties--http_receiver--auth_basic.md): complete subsection reference.

- [auth_none](resources--global_log_receiver--properties--http_receiver--auth_none.md): complete subsection reference.

- [auth_token](resources--global_log_receiver--properties--http_receiver--auth_token.md): complete subsection reference.

- [batch](resources--global_log_receiver--properties--http_receiver--batch.md): complete subsection reference.

- [compression](resources--global_log_receiver--properties--http_receiver--compression.md): complete subsection reference.

- [no_tls](resources--global_log_receiver--properties--http_receiver--no_tls.md): complete subsection reference.

<a id="schema-http_receiver--uri"></a>

### uri property

Type: `"string"`. Optional.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

- [use_tls](resources--global_log_receiver--properties--http_receiver--use_tls.md): complete subsection reference.

## Next pages

- [http_receiver.auth_basic](resources--global_log_receiver--properties--http_receiver--auth_basic.md)
- [http_receiver.auth_none](resources--global_log_receiver--properties--http_receiver--auth_none.md)
- [http_receiver.auth_token](resources--global_log_receiver--properties--http_receiver--auth_token.md)
- [http_receiver.batch](resources--global_log_receiver--properties--http_receiver--batch.md)
- [http_receiver.compression](resources--global_log_receiver--properties--http_receiver--compression.md)
- [http_receiver.no_tls](resources--global_log_receiver--properties--http_receiver--no_tls.md)
- [http_receiver.use_tls](resources--global_log_receiver--properties--http_receiver--use_tls.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
