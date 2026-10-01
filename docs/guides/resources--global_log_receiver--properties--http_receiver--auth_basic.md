---
page_title: "http_receiver.auth_basic"
subcategory: ""
description: "http_receiver.auth_basic for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2189, "body_sha256": "sha256:edc63b139c389ebf4d0dbd0f3955a2a2050ee77ae788fb34d1582950fff03a99", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic:password"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "path": "docs/guides/resources--global_log_receiver--properties--http_receiver--auth_basic.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "auth_basic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/auth_basic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.auth_basic for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.auth_basic

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [http_receiver](resources--global_log_receiver--properties--http_receiver.md)
- http_receiver.auth_basic

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters to access HTPP Log Receiver Endpoint.

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

Terraform syntax:

```terraform
auth_basic {
  # Configure direct properties listed below.
}
```

## Direct properties

- [password](resources--global_log_receiver--properties--http_receiver--auth_basic--password.md): complete subsection reference.

<a id="schema-http_receiver--auth_basic--user_name"></a>

### user_name property

Type: `"string"`. Optional.

User Name. HTTP Basic Auth User Name.

Upstream description:

HTTP Basic Auth User Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [http_receiver.auth_basic.password](resources--global_log_receiver--properties--http_receiver--auth_basic--password.md)
- [http_receiver](resources--global_log_receiver--properties--http_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
