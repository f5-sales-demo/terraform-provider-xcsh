---
page_title: "client_side_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules exclude list"], "body_bytes": 2986, "body_sha256": "sha256:110995f4f14e71d0e83b745f82d4bd986f8acddcc34a738a714d51adfa478f99", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "documentation/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-013.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insertion rules exclude list any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules exclude list domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["client side defense policy js insertion rules exclude list metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "type": "requires"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["client side defense policy js insertion rules exclude list path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--path", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-client_side_defense--policy--js_insertion_rules--exclude_list--path--regex", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
