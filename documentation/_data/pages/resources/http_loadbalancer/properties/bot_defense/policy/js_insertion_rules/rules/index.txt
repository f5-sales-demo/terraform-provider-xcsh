---
page_title: "bot_defense.policy.js_insertion_rules.rules"
subcategory: "Load Balancing"
description: "Required list of pages to insert Bot Defense client JavaScript."
xcsh_docs: {"aliases": ["bot defense policy js insertion rules rules"], "body_bytes": 5288, "body_sha256": "sha256:be2c224a416c839d02c7114b10b16a92dd92d969c5e9613ff5f8de3ed14111fa", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy js insertion rules rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy js insertion rules rules domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--domain--exact_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--domain--regex_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--domain--suffix_value", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["bot defense policy js insertion rules rules javascript location"], "anchor": "schema-bot_defense--policy--js_insertion_rules--rules--javascript_location", "description": "All inside networks. Insert JavaScript after <HEAD> tag Insert JavaScript after </title> tag. Insert JavaScript before first tag.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules", "javascript_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot defense policy js insertion rules rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--metadata--name", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:metadata", "type": "requires"}], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["bot defense policy js insertion rules rules path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--path--path", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--path--prefix", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--js_insertion_rules--rules--path--regex", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insertion_rules.rules.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules", "path"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Required list of pages to insert Bot Defense client JavaScript.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/)
- bot_defense.policy.js_insertion_rules.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/domain/): complete subsection reference.

<a id="schema-bot_defense--policy--js_insertion_rules--rules--javascript_location"></a>

### javascript_location property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/path/): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insertion_rules.rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/any_domain/)
- [bot_defense.policy.js_insertion_rules.rules.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/domain/)
- [bot_defense.policy.js_insertion_rules.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/metadata/)
- [bot_defense.policy.js_insertion_rules.rules.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/path/)
- [bot_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
