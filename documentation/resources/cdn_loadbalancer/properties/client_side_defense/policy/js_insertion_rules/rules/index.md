---
page_title: "client_side_defense.policy.js_insertion_rules.rules"
subcategory: "Load Balancing"
description: "Required list of pages to insert Client-Side Defense client JavaScript."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules rules"], "body_bytes": 4365, "body_sha256": "sha256:e171679cde5b0849e6c855a0c866517ea467e9bafbaf5c12d6e6e437bb401618", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "documentation/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--exact_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--exact_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--regex_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--regex_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--suffix_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--domain--suffix_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--metadata--name", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "type": "requires"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--path", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--path", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--prefix", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--prefix", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--regex", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--regex", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Required list of pages to insert Client-Side Defense client JavaScript.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- client_side_defense.policy.js_insertion_rules.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/path/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/any_domain/)
- [client_side_defense.policy.js_insertion_rules.rules.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/domain/)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/metadata/)
- [client_side_defense.policy.js_insertion_rules.rules.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/path/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
