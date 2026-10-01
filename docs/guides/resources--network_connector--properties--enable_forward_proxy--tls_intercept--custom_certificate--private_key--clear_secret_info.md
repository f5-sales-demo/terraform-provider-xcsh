---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 4100, "body_sha256": "sha256:b92b3fbada1038dec95ee075bb66f3e116ca9d04a21228cf13eb2bd65aba84ef", "canonical_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key", "path": "docs/guides/resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [enable_forward_proxy](resources--network_connector--properties--enable_forward_proxy.md)
- [enable_forward_proxy.tls_intercept](resources--network_connector--properties--enable_forward_proxy--tls_intercept.md)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key.md)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key.md)
- [xcsh_network_connector](../resources/network_connector.md)
