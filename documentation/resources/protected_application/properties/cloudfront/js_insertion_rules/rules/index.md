---
page_title: "cloudfront.js_insertion_rules.rules"
subcategory: ""
description: "Required list of pages to insert Bot Defense client JavaScript."
xcsh_docs: {"aliases": ["cloudfront js insertion rules rules"], "body_bytes": 7749, "body_sha256": "sha256:99b9eb5347ac7b5f7cc86715386cd6be5ac81b55878b7ee3e01b3a4f3a2f6e5c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:any_domain", "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:metadata"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules", "path": "documentation/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--rules--exact_path", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,glob", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--exact_path", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--glob", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,glob", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--glob", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:glob,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--prefix", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--prefix", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:glob,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "js_insertion_rules", "rules"], "schema_version": 1, "sections": [{"aliases": ["cloudfront js insertion rules rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront js insertion rules rules domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--rules--domain--exact_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--domain--exact_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--domain--regex_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--domain--regex_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}], "schema_path": ["cloudfront", "js_insertion_rules", "rules", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["cloudfront js insertion rules rules exact path"], "anchor": "schema-cloudfront--js_insertion_rules--rules--exact_path", "description": "Exclusive with Exact path value to match.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "rules", "exact_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules rules glob"], "anchor": "schema-cloudfront--js_insertion_rules--rules--glob", "description": "Exclusive with Accepts wildcards * to match multiple characters or ? To match a single character.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "rules", "glob"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--rules--metadata--name", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:metadata", "type": "requires"}], "schema_path": ["cloudfront", "js_insertion_rules", "rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["cloudfront js insertion rules rules prefix"], "anchor": "schema-cloudfront--js_insertion_rules--rules--prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "rules", "prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Required list of pages to insert Bot Defense client JavaScript.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/)
- cloudfront.js_insertion_rules.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("exact_path",
    "glob"),
  validators.ConflictingListObjectAttributes("exact_path",
    "prefix"),
  validators.ConflictingListObjectAttributes("glob",
    "prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/domain/): complete subsection reference.

<a id="schema-cloudfront--js_insertion_rules--rules--exact_path"></a>

### exact_path property

Type: `"string"`. Optional.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-cloudfront--js_insertion_rules--rules--glob"></a>

### glob property

Type: `"string"`. Optional.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
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
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/metadata/): complete subsection reference.

<a id="schema-cloudfront--js_insertion_rules--rules--prefix"></a>

### prefix property

Type: `"string"`. Optional.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [cloudfront.js_insertion_rules.rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/any_domain/)
- [cloudfront.js_insertion_rules.rules.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/domain/)
- [cloudfront.js_insertion_rules.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/metadata/)
- [cloudfront.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
