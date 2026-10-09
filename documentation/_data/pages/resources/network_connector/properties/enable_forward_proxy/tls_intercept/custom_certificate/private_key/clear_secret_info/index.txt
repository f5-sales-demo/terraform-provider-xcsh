---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info"
subcategory: "Networking"
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["enable forward proxy tls intercept custom certificate private key clear secret info"], "body_bytes": 3963, "body_sha256": "sha256:37a368a1ced33ecceaf5eed42d0442e3330ccac56b08f7db82a10fa2929765b6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0032033013102232-3303100213113021-2233201201133120-2101321230020221-3000032021020022-1111303102333230-3311212113232320-1313231203200032", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [{"anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["enable forward proxy tls intercept custom certificate private key clear secret info provider ref"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable forward proxy tls intercept custom certificate private key clear secret info url"], "anchor": "schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
