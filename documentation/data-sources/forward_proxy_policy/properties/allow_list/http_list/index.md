---
page_title: "allow_list.http_list"
subcategory: "Security"
description: "URLs for HTTP connections."
xcsh_docs: {"aliases": ["allow list http list"], "body_bytes": 9303, "body_sha256": "sha256:6d00db61a7ac282d85b27956466f20fb9496392283f9f1f656b9b50387fceee6", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list:any_path"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list", "path": "documentation/data-sources/forward_proxy_policy/properties/allow_list/http_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_list", "http_list"], "schema_version": 1, "sections": [{"aliases": ["any path"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list:any_path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "any_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact value"], "anchor": "schema-allow_list--http_list--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["path exact value"], "anchor": "schema-allow_list--http_list--path_exact_value", "description": "Exclusive with Exact Path to match.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "path_exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["path prefix value"], "anchor": "schema-allow_list--http_list--path_prefix_value", "description": "Exclusive with Prefix of Path e.g \"/abc/xyz\" will match \"/abc/xyz/.*\"", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "path_prefix_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["path regex value"], "anchor": "schema-allow_list--http_list--path_regex_value", "description": "Exclusive with Regular Expression value for the Path to match.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "path_regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex value"], "anchor": "schema-allow_list--http_list--regex_value", "description": "Exclusive with Regular Expression value for the domain name.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["suffix value"], "anchor": "schema-allow_list--http_list--suffix_value", "description": "Exclusive with Suffix of domain names e.g \"xyz.com\" will match \"*.xyz.com\"", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "http_list", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/allow_list/http_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "URLs for HTTP connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_list.http_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/)
- allow_list.http_list

<a id="section"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/http_list/any_path/): complete subsection reference.

<a id="schema-allow_list--http_list--exact_value"></a>

### exact_value property

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="schema-allow_list--http_list--path_exact_value"></a>

### path_exact_value property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-allow_list--http_list--path_prefix_value"></a>

### path_prefix_value property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-allow_list--http_list--path_regex_value"></a>

### path_regex_value property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-allow_list--http_list--regex_value"></a>

### regex_value property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-allow_list--http_list--suffix_value"></a>

### suffix_value property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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

## Next pages

- [allow_list.http_list.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/http_list/any_path/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
