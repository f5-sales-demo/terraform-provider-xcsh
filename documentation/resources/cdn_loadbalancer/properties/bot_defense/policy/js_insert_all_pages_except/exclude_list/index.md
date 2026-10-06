---
page_title: "bot_defense.policy.js_insert_all_pages_except.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["bot defense policy js insert all pages except exclude list"], "body_bytes": 2944, "body_sha256": "sha256:c2c4cebd17818870c67568106c4a89f95c5f47fa494295d3f8d779cc754d81a2", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy js insert all pages except exclude list any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy js insert all pages except exclude list domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["bot defense policy js insert all pages except exclude list metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "type": "requires"}], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["bot defense policy js insert all pages except exclude list path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insert_all_pages_except.exclude_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/path/): complete subsection reference.
