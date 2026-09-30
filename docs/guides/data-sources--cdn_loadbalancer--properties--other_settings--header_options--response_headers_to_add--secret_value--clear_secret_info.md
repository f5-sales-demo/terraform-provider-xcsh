---
page_title: "other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info"
subcategory: "Load Balancing"
description: "other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3588, "body_sha256": "sha256:e198ddac0de5f235992674cc9d83030d4f3c4f2d076654dab05af81b7afe3be2", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:clear_secret_info", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--other_settings--header_options--response_headers_to_add--secret_value--clear_secret_info.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings", "header_options", "response_headers_to_add", "secret_value", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [other_settings](data-sources--cdn_loadbalancer--properties--other_settings.md)
- [other_settings.header_options](data-sources--cdn_loadbalancer--properties--other_settings--header_options.md)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--response_headers_to_add.md)
- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--properties--other_settings--header_options--response_headers_to_add--secret_value.md)
- other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

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

<a id="schema-other_settings--header_options--response_headers_to_add--secret_value--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-other_settings--header_options--response_headers_to_add--secret_value--clear_secret_info--url"></a>

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

- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--properties--other_settings--header_options--response_headers_to_add--secret_value.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
