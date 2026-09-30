---
page_title: "access_info.tls_config"
subcategory: ""
description: "access_info.tls_config for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 6082, "body_sha256": "sha256:3ab06ce4e02f8c447cb0a763f8b7d40f90a1ab182d077cbfe9e8f190d3f2efe5", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:default_session_key_caching", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_sni", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:use_host_header_as_sni"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "docs/guides/resources--secret_management_access--properties--access_info--tls_config.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# access_info.tls_config

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- access_info.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cert_params",
    "common_params"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cert_params](resources--secret_management_access--properties--access_info--tls_config--cert_params.md): complete subsection reference.

- [common_params](resources--secret_management_access--properties--access_info--tls_config--common_params.md): complete subsection reference.

- [default_session_key_caching](resources--secret_management_access--properties--access_info--tls_config--default_session_key_caching.md): complete subsection reference.

- [disable_session_key_caching](resources--secret_management_access--properties--access_info--tls_config--disable_session_key_caching.md): complete subsection reference.

- [disable_sni](resources--secret_management_access--properties--access_info--tls_config--disable_sni.md): complete subsection reference.

<a id="schema-access_info--tls_config--max_session_keys"></a>

### max_session_keys property

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

<a id="schema-access_info--tls_config--sni"></a>

### sni property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [use_host_header_as_sni](resources--secret_management_access--properties--access_info--tls_config--use_host_header_as_sni.md): complete subsection reference.

## Next pages

- [access_info.tls_config.cert_params](resources--secret_management_access--properties--access_info--tls_config--cert_params.md)
- [access_info.tls_config.common_params](resources--secret_management_access--properties--access_info--tls_config--common_params.md)
- [access_info.tls_config.default_session_key_caching](resources--secret_management_access--properties--access_info--tls_config--default_session_key_caching.md)
- [access_info.tls_config.disable_session_key_caching](resources--secret_management_access--properties--access_info--tls_config--disable_session_key_caching.md)
- [access_info.tls_config.disable_sni](resources--secret_management_access--properties--access_info--tls_config--disable_sni.md)
- [access_info.tls_config.use_host_header_as_sni](resources--secret_management_access--properties--access_info--tls_config--use_host_header_as_sni.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
