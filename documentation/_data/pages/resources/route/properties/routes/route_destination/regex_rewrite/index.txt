---
page_title: "routes.route_destination.regex_rewrite"
subcategory: ""
description: "RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string."
xcsh_docs: {"aliases": ["routes route destination regex rewrite"], "body_bytes": 3436, "body_sha256": "sha256:c7dd7d83a6847a69eb44819d5f3d7057616e1f850f317e4bc5e8f1723fb59292", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/regex_rewrite/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1002132200332031-2112010113122321-3210320213210322-3203102103022111-0020102231301312-3020023031212111-1313020132121103-1031331331113022", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "regex_rewrite"], "schema_version": 1, "sections": [{"aliases": ["routes route destination regex rewrite pattern"], "anchor": "schema-routes--route_destination--regex_rewrite--pattern", "description": "The regular expression used to find portions of a string that should be replaced.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "regex_rewrite", "pattern"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination regex rewrite substitution"], "anchor": "schema-routes--route_destination--regex_rewrite--substitution", "description": "The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "regex_rewrite", "substitution"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/regex_rewrite/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["routeCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
