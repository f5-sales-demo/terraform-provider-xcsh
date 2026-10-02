---
page_title: "routes.route_destination.regex_rewrite"
subcategory: ""
description: "RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string."
xcsh_docs: {"aliases": ["routes route destination regex rewrite"], "body_bytes": 3944, "body_sha256": "sha256:c3b0c23000f9d0d4d8a4ae2456784a989990392ae636c16fac494e70997c3a12", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/regex_rewrite/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1002132200332031-2112010113122321-3210320213210322-3203102103022111-0020102231301312-3020023031212111-1313020132121103-1031331331113022", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "regex_rewrite"], "schema_version": 1, "sections": [{"aliases": ["pattern"], "anchor": "schema-routes--route_destination--regex_rewrite--pattern", "description": "The regular expression used to find portions of a string that should be replaced.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "regex_rewrite", "pattern"], "syntax": "attribute", "type": "string"}, {"aliases": ["substitution"], "anchor": "schema-routes--route_destination--regex_rewrite--substitution", "description": "The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "regex_rewrite", "substitution"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/regex_rewrite/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.regex_rewrite

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.regex_rewrite

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

Upstream description:

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

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
regex_rewrite {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--route_destination--regex_rewrite--pattern"></a>

### pattern property

Type: `"string"`. Optional.

The regular expression used to find portions of a string that should be replaced.

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

<a id="schema-routes--route_destination--regex_rewrite--substitution"></a>

### substitution property

Type: `"string"`. Optional.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

Upstream description:

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
