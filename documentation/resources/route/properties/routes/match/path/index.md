---
page_title: "routes.match.path"
subcategory: ""
description: "Path match of the URI can be either be, Prefix match or exact match or regular expression match."
xcsh_docs: {"aliases": ["routes match path"], "body_bytes": 5357, "body_sha256": "sha256:d14911a5b582e507795af53946efb943ca3e5f47ea9dbecb700be3e20a2f1db2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match:path", "parent_id": "xcsh-docs:resources:route:properties:routes:match", "path": "documentation/resources/route/properties/routes/match/path/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1002312122123222-2033203330332312-1323123113131312-2032021222022000-2220031020111332-3000201300302323-3333213122233121-0112231033102033", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [{"anchor": "schema-routes--match--path--path", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--path", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--prefix", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--prefix", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--regex", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--regex", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "match", "path"], "schema_version": 1, "sections": [{"aliases": ["path"], "anchor": "schema-routes--match--path--path", "description": "Exclusive with Exact path value to match.", "document_id": "xcsh-docs:resources:route:properties:routes:match:path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "match", "path", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["prefix"], "anchor": "schema-routes--match--path--prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:resources:route:properties:routes:match:path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "match", "path", "prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex"], "anchor": "schema-routes--match--path--regex", "description": "Exclusive with Regular expression of path match (e.g. The value .* will match on all paths)", "document_id": "xcsh-docs:resources:route:properties:routes:match:path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "match", "path", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/path/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.match.path

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/)
- routes.match.path

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--match--path--path"></a>

### path property

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="schema-routes--match--path--prefix"></a>

### prefix property

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="schema-routes--match--path--regex"></a>

### regex property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

## Next pages

- [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
