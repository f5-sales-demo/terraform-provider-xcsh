---
page_title: "cloudfront.js_insertion_rules.exclude_list"
subcategory: ""
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["cloudfront js insertion rules exclude list"], "body_bytes": 3851, "body_sha256": "sha256:830e12364795f1b2a8f5efcbcdbf240dd6c1f238e903dd924270c58d3fe5df65", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules", "path": "documentation/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:metadata", "type": "requires"}], "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:path", "type": "conflicts"}], "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/)
- cloudfront.js_insertion_rules.exclude_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/path/): complete subsection reference.

## Next pages

- [cloudfront.js_insertion_rules.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/any_domain/)
- [cloudfront.js_insertion_rules.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/domain/)
- [cloudfront.js_insertion_rules.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/metadata/)
- [cloudfront.js_insertion_rules.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/path/)
- [cloudfront.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
