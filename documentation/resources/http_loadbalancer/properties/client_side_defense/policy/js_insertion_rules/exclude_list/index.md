---
page_title: "client_side_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules exclude list"], "body_bytes": 4248, "body_sha256": "sha256:b5686040c8d30ebd321edff3c36388bc49719f8b3c395ef26825bd615e219294", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "documentation/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "type": "requires"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- client_side_defense.policy.js_insertion_rules.exclude_list

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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/path/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/any_domain/)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/domain/)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/metadata/)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/path/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
