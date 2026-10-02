---
page_title: "bot_defense.policy.js_insert_all_pages_except.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["bot defense policy js insert all pages except exclude list"], "body_bytes": 4199, "body_sha256": "sha256:f5efbd6a57bafb3bd0d049c23ec49b7698d565fa7ad9a3b4ebd658cb646a935b", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3321230333023200-0223133222102101-3223130312131313-1330202302222030-3133112013000101-0023222133230111-3321013222321031-3111022101032131", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:domain", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "type": "requires"}], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:path", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/path/): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/any_domain/)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/domain/)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/metadata/)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/path/)
- [bot_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
