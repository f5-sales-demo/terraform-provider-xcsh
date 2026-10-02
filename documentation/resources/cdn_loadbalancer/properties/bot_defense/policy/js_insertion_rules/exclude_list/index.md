---
page_title: "bot_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["bot defense policy js insertion rules exclude list"], "body_bytes": 4055, "body_sha256": "sha256:76ab439a60ed778584ad3f024598966d9fd0b0eee8290dabe0194853d26582f3", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0232312133003302-3231123221033012-2220322232002010-0110200313231100-2101100000322313-3201013130133031-2333102320303010-2120301002011010", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:metadata", "type": "requires"}], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/)
- bot_defense.policy.js_insertion_rules.exclude_list

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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/path/): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/any_domain/)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/domain/)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/metadata/)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/path/)
- [bot_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
