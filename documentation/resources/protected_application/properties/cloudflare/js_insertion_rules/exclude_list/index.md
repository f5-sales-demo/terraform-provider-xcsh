---
page_title: "cloudflare.js_insertion_rules.exclude_list"
subcategory: ""
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["cloudflare js insertion rules exclude list"], "body_bytes": 2751, "body_sha256": "sha256:d43492ff9796369fe2201b22280aa7713dc03de69fe6e8fed10d72b8ba6fb580", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules", "path": "documentation/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["cloudflare js insertion rules exclude list any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare js insertion rules exclude list domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare js insertion rules exclude list metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "type": "requires"}], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare js insertion rules exclude list path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path", "type": "conflicts"}], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/)
- cloudflare.js_insertion_rules.exclude_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/path/): complete subsection reference.
