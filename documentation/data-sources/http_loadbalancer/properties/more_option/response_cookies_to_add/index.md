---
page_title: "more_option.response_cookies_to_add"
subcategory: "Load Balancing"
description: "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied."
xcsh_docs: {"aliases": ["more option response cookies to add"], "body_bytes": 15605, "body_sha256": "sha256:0170c49f71fa7b0f310016f68c30c6bb3f9ef8e45c7b156cf827974a5a39b2be", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:add_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:add_partitioned", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:add_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_expiry", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_max_age", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_partitioned", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_path", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_samesite", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_value", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_lax", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_none", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_strict", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:secret_value"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "path": "documentation/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option", "response_cookies_to_add"], "schema_version": 1, "sections": [{"aliases": ["add domain"], "anchor": "schema-more_option--response_cookies_to_add--add_domain", "description": "Exclusive with Add domain attribute.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "add_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["add expiry"], "anchor": "schema-more_option--response_cookies_to_add--add_expiry", "description": "Exclusive with Add expiry attribute.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "add_expiry"], "syntax": "attribute", "type": "string"}, {"aliases": ["add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:add_httponly", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["add partitioned"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:add_partitioned", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "add_partitioned"], "syntax": "attribute", "type": "object"}, {"aliases": ["add path"], "anchor": "schema-more_option--response_cookies_to_add--add_path", "description": "Exclusive with Add path attribute.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "add_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:add_secure", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore expiry"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_expiry", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_expiry"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_httponly", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore max age"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_max_age", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_max_age"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore partitioned"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_partitioned", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_partitioned"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore path"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_samesite", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_secure", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore value"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "ignore_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["max age value"], "anchor": "schema-more_option--response_cookies_to_add--max_age_value", "description": "Exclusive with Add max age attribute.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "max_age_value"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-more_option--response_cookies_to_add--name", "description": "Name of the cookie in Cookie header.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["overwrite"], "anchor": "schema-more_option--response_cookies_to_add--overwrite", "description": "Should the value be overwritten? If true, the value is overwritten to existing values. Default value is do not overwrite.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "overwrite"], "syntax": "attribute", "type": "bool"}, {"aliases": ["samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_lax", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:samesite_strict", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "samesite_strict"], "syntax": "attribute", "type": "object"}, {"aliases": ["secret value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add:secret_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "secret_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["value"], "anchor": "schema-more_option--response_cookies_to_add--value", "description": "Exclusive with Value of the Cookie header.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_add", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.response_cookies_to_add

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/)
- more_option.response_cookies_to_add

<a id="section"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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

<a id="schema-more_option--response_cookies_to_add--add_domain"></a>

### add_domain property

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
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
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-more_option--response_cookies_to_add--add_expiry"></a>

### add_expiry property

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/add_httponly/): complete subsection reference.

- [add_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/add_partitioned/): complete subsection reference.

<a id="schema-more_option--response_cookies_to_add--add_path"></a>

### add_path property

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/add_secure/): complete subsection reference.

- [ignore_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_domain/): complete subsection reference.

- [ignore_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_expiry/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_httponly/): complete subsection reference.

- [ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_max_age/): complete subsection reference.

- [ignore_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_partitioned/): complete subsection reference.

- [ignore_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_path/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_secure/): complete subsection reference.

- [ignore_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_value/): complete subsection reference.

<a id="schema-more_option--response_cookies_to_add--max_age_value"></a>

### max_age_value property

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="schema-more_option--response_cookies_to_add--name"></a>

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

<a id="schema-more_option--response_cookies_to_add--overwrite"></a>

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

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_strict/): complete subsection reference.

- [secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/): complete subsection reference.

<a id="schema-more_option--response_cookies_to_add--value"></a>

### value property

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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

- [more_option.response_cookies_to_add.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/add_httponly/)
- [more_option.response_cookies_to_add.add_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/add_partitioned/)
- [more_option.response_cookies_to_add.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/add_secure/)
- [more_option.response_cookies_to_add.ignore_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_domain/)
- [more_option.response_cookies_to_add.ignore_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_expiry/)
- [more_option.response_cookies_to_add.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_httponly/)
- [more_option.response_cookies_to_add.ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_max_age/)
- [more_option.response_cookies_to_add.ignore_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_partitioned/)
- [more_option.response_cookies_to_add.ignore_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_path/)
- [more_option.response_cookies_to_add.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_samesite/)
- [more_option.response_cookies_to_add.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_secure/)
- [more_option.response_cookies_to_add.ignore_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_value/)
- [more_option.response_cookies_to_add.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_lax/)
- [more_option.response_cookies_to_add.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_none/)
- [more_option.response_cookies_to_add.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/samesite_strict/)
- [more_option.response_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/secret_value/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
