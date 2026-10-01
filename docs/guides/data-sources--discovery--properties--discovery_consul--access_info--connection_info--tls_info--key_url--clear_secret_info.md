---
page_title: "discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info"
subcategory: ""
description: "discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 3721, "body_sha256": "sha256:9d67fa94d105d4ceb546d2fa650bf85a75ea3d6784c59b99c1288af73bda4e7f", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url", "path": "docs/guides/data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info", "key_url", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_consul](data-sources--discovery--properties--discovery_consul.md)
- [discovery_consul.access_info](data-sources--discovery--properties--discovery_consul--access_info.md)
- [discovery_consul.access_info.connection_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info.md)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info.md)
- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url.md)
- discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

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

<a id="schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--url"></a>

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

- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--properties--discovery_consul--access_info--connection_info--tls_info--key_url.md)
- [xcsh_discovery](../data-sources/discovery.md)
