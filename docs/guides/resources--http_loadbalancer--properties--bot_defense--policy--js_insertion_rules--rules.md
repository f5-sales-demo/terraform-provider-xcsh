---
page_title: "bot_defense.policy.js_insertion_rules.rules"
subcategory: "Load Balancing"
description: "bot_defense.policy.js_insertion_rules.rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4459, "body_sha256": "sha256:f63c1b6fe4258e2b48f076a34ce3685ea2231f877056791245abe1bb51ac19b2", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules:path"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "js_insertion_rules", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.js_insertion_rules.rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md)
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

- [any_domain](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--any_domain.md): complete subsection reference.

- [domain](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--domain.md): complete subsection reference.

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

- [metadata](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--metadata.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--path.md): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--any_domain.md)
- [bot_defense.policy.js_insertion_rules.rules.domain](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--domain.md)
- [bot_defense.policy.js_insertion_rules.rules.metadata](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--metadata.md)
- [bot_defense.policy.js_insertion_rules.rules.path](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules--path.md)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
