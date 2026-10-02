---
page_title: "cloudflare.js_insertion_rules.exclude_list.path"
subcategory: ""
description: "Path match of the URI can be either be, Prefix match or exact match or regular expression match."
xcsh_docs: {"aliases": ["cloudflare js insertion rules exclude list path"], "body_bytes": 5931, "body_sha256": "sha256:e02ef5450e06ea30708cb4c368efd9a62aa8e12daae2e771a02910c73c577a46", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list", "path": "documentation/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/path/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0122200033020333-1233203121310121-3300011130303222-1021030311232233-0232031012010030-3001330100331330-1002033013202110-3303222112120231", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [{"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "path"], "schema_version": 1, "sections": [{"aliases": ["path"], "anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--path", "description": "Exclusive with Exact path value to match.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "path", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["prefix"], "anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "path", "prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex"], "anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--regex", "description": "Exclusive with Regular expression of path match (e.g. The value .* will match on all paths)", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "path", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/path/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.js_insertion_rules.exclude_list.path

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/)
- [cloudflare.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/)
- cloudflare.js_insertion_rules.exclude_list.path

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

<a id="schema-cloudflare--js_insertion_rules--exclude_list--path--path"></a>

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

<a id="schema-cloudflare--js_insertion_rules--exclude_list--path--prefix"></a>

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

<a id="schema-cloudflare--js_insertion_rules--exclude_list--path--regex"></a>

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

- [cloudflare.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
