---
page_title: "dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info"
subcategory: ""
description: "dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4143, "body_sha256": "sha256:5d3a5d8293e90a148ad734dd5026ce1ef81b05aa44a243cd97aac03a81f30880", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add:secret_value:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add:secret_value:clear_secret_info", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add:secret_value", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add", "secret_value", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/secret_value/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.http_proxy](resources--proxy--properties--dynamic_proxy--http_proxy.md)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--properties--dynamic_proxy--http_proxy--more_option.md)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add.md)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value.md)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

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

<a id="schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--url"></a>

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

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value.md)
- [xcsh_proxy](../resources/proxy.md)
