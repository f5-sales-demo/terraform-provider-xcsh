---
page_title: "cloudflare.js_insertion_rules.rules"
subcategory: ""
description: "Required list of pages to insert Bot Defense client JavaScript."
xcsh_docs: {"aliases": ["cloudflare js insertion rules rules"], "body_bytes": 6871, "body_sha256": "sha256:8ef99dd9151a525ed224dd6f0942d7cb2d9585de3c21e4ab19e716b87f4b8ce8", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules:any_domain", "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules:domain", "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules:metadata"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules", "path": "documentation/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321", "registry_path": "docs/guides/data-sources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "js_insertion_rules", "rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "rules", "domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact path"], "anchor": "schema-cloudflare--js_insertion_rules--rules--exact_path", "description": "Exclusive with Exact path value to match.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "rules", "exact_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["glob"], "anchor": "schema-cloudflare--js_insertion_rules--rules--glob", "description": "Exclusive with Accepts wildcards * to match multiple characters or ? To match a single character.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "rules", "glob"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix"], "anchor": "schema-cloudflare--js_insertion_rules--rules--prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "rules", "prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Required list of pages to insert Bot Defense client JavaScript.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- [cloudflare.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/)
- cloudflare.js_insertion_rules.rules

<a id="section"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/domain/): complete subsection reference.

<a id="schema-cloudflare--js_insertion_rules--rules--exact_path"></a>

### exact_path property

Type: `"string"`. Computed.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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

<a id="schema-cloudflare--js_insertion_rules--rules--glob"></a>

### glob property

Type: `"string"`. Computed.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/metadata/): complete subsection reference.

<a id="schema-cloudflare--js_insertion_rules--rules--prefix"></a>

### prefix property

Type: `"string"`. Computed.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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

## Next pages

- [cloudflare.js_insertion_rules.rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/any_domain/)
- [cloudflare.js_insertion_rules.rules.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/domain/)
- [cloudflare.js_insertion_rules.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/rules/metadata/)
- [cloudflare.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/js_insertion_rules/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
