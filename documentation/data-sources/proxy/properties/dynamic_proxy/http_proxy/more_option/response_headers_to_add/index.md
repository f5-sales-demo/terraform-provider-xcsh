---
page_title: "dynamic_proxy.http_proxy.more_option.response_headers_to_add"
subcategory: ""
description: "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied."
xcsh_docs: {"aliases": ["dynamic proxy http proxy more option response headers to add"], "body_bytes": 5796, "body_sha256": "sha256:94f4b3a62777f2a6a928f2fb91d20af9e2b940f7c1aac34e9af87094012b830e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add:secret_value"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123", "registry_path": "docs/guides/data-sources--proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add"], "schema_version": 1, "sections": [{"aliases": ["append"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--append", "description": "Should the value be appended? If true, the value is appended to existing values. Default value is do not append.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add", "append"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--name", "description": "Name of the HTTP header.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["secret value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add:secret_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add", "secret_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["value"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--value", "description": "Exclusive with Value of the HTTP header.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option.response_headers_to_add

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/)
- [dynamic_proxy.http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add

<a id="section"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--append"></a>

### append property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/secret_value/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--value"></a>

### value property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

## Next pages

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/secret_value/)
- [dynamic_proxy.http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
