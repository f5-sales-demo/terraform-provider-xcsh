---
page_title: "http_proxy.more_option.request_cookies_to_add"
subcategory: ""
description: "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied."
xcsh_docs: {"aliases": ["http proxy more option request cookies to add"], "body_bytes": 5627, "body_sha256": "sha256:80a519250b92602a8f6e8664e29aec26041a3a905ba7a34045890ad1641343a2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:request_cookies_to_add:secret_value"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:request_cookies_to_add", "parent_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option", "path": "documentation/data-sources/proxy/properties/http_proxy/more_option/request_cookies_to_add/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2333120002103000-0021233110122311-0220031112322320-0123003233011103-0001231311223111-0232001220300220-1213321001012222-2003330220332213", "registry_path": "docs/guides/data-sources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "request_cookies_to_add"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-http_proxy--more_option--request_cookies_to_add--name", "description": "Name of the cookie in Cookie header.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:request_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_proxy", "more_option", "request_cookies_to_add", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["overwrite"], "anchor": "schema-http_proxy--more_option--request_cookies_to_add--overwrite", "description": "Should the value be overwritten? If true, the value is overwritten to existing values. Default value is do not overwrite.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:request_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_proxy", "more_option", "request_cookies_to_add", "overwrite"], "syntax": "attribute", "type": "bool"}, {"aliases": ["secret value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:request_cookies_to_add:secret_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_proxy", "more_option", "request_cookies_to_add", "secret_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["value"], "anchor": "schema-http_proxy--more_option--request_cookies_to_add--value", "description": "Exclusive with Value of the Cookie header.", "document_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:request_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_proxy", "more_option", "request_cookies_to_add", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/http_proxy/more_option/request_cookies_to_add/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.request_cookies_to_add

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/)
- http_proxy.more_option.request_cookies_to_add

<a id="section"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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

<a id="schema-http_proxy--more_option--request_cookies_to_add--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-http_proxy--more_option--request_cookies_to_add--overwrite"></a>

### overwrite property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/request_cookies_to_add/secret_value/): complete subsection reference.

<a id="schema-http_proxy--more_option--request_cookies_to_add--value"></a>

### value property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

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

- [http_proxy.more_option.request_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/request_cookies_to_add/secret_value/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
