---
page_title: "access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info"
subcategory: ""
description: "access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 3781, "body_sha256": "sha256:d55ee9ad15ea51e6df31c6556733793786318d1e3f858dcb9026d0c936ed4631", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:clear_secret_info", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "path": "docs/guides/data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "private_key", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](data-sources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.tls_config.common_params](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md)
- access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

## Direct properties

<a id="schema-access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
