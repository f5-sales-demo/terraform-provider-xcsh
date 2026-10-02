---
page_title: "more_option.request_cookies_to_add"
subcategory: "Load Balancing"
description: "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied."
xcsh_docs: {"aliases": ["more option request cookies to add"], "body_bytes": 6097, "body_sha256": "sha256:c3721106ddfc10bd81c2d5c6a8b905dd3a9da608af9c4ec903c068d2b475c5a3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add:secret_value"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option", "path": "documentation/resources/http_loadbalancer/properties/more_option/request_cookies_to_add/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-021.md", "relationships": [{"anchor": "schema-more_option--request_cookies_to_add--value", "enforcement": "provider-schema", "group": "more_option.request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "more_option.request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-more_option--request_cookies_to_add--name", "enforcement": "provider-schema", "group": "more_option.request_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option", "request_cookies_to_add"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-more_option--request_cookies_to_add--name", "description": "Name of the cookie in Cookie header.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "request_cookies_to_add", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["overwrite"], "anchor": "schema-more_option--request_cookies_to_add--overwrite", "description": "Should the value be overwritten? If true, the value is overwritten to existing values. Default value is do not overwrite.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "request_cookies_to_add", "overwrite"], "syntax": "attribute", "type": "bool"}, {"aliases": ["secret value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add:secret_value", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "more_option.request_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "more_option.request_cookies_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "schema_path": ["more_option", "request_cookies_to_add", "secret_value"], "syntax": "block", "type": "object"}, {"aliases": ["value"], "anchor": "schema-more_option--request_cookies_to_add--value", "description": "Exclusive with Value of the Cookie header.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:request_cookies_to_add", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "request_cookies_to_add", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/request_cookies_to_add/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.request_cookies_to_add

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/)
- more_option.request_cookies_to_add

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-more_option--request_cookies_to_add--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="schema-more_option--request_cookies_to_add--overwrite"></a>

### overwrite property

Type: `"bool"`. Optional.

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

- [secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/request_cookies_to_add/secret_value/): complete subsection reference.

<a id="schema-more_option--request_cookies_to_add--value"></a>

### value property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [more_option.request_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/request_cookies_to_add/secret_value/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
