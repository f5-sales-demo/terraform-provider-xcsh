---
page_title: "rules.action"
subcategory: ""
description: "Action to be enforced if the BGP route matches the rule."
xcsh_docs: {"aliases": ["rules action"], "body_bytes": 5271, "body_sha256": "sha256:9bdcb6c466e75ae4e31ac1a43d3d5e5c4869f072abfaeb685ebfbf081da01e14", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules", "path": "documentation/resources/bgp_routing_policy/properties/rules/action/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2121133023232222-3022321101013020-3111022210121220-1220231321203201-0003302103002213-2010322121312130-3202022002222021-0311330313120023", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,as_path", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:local_preference,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:local_preference,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,as_path", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "sections": [{"aliases": ["allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["as path"], "anchor": "schema-rules--action--as_path", "description": "Exclusive with AS-Path Prepending is generally used to influence incoming traffic.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "as_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["community"], "anchor": "section", "description": "List of BGP communities.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--action--community--community", "enforcement": "provider-schema", "group": "rules.action.community:RequiredObjectAttributes:community", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "requires"}], "schema_path": ["rules", "action", "community"], "syntax": "block", "type": "object"}, {"aliases": ["deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "deny"], "syntax": "attribute", "type": "object"}, {"aliases": ["local preference"], "anchor": "schema-rules--action--local_preference", "description": "Exclusive with BGP Local Preference is generally used to influence outgoing traffic.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "local_preference"], "syntax": "attribute", "type": "number"}, {"aliases": ["metric"], "anchor": "schema-rules--action--metric", "description": "Exclusive with The Multi-Exit Discriminator metric to indicate the preferred path to AS.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "metric"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action to be enforced if the BGP route matches the rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- rules.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action to be enforced if the BGP route matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "as_path"),
  validators.ConflictingObjectAttributes("allow",
    "community"),
  validators.ConflictingObjectAttributes("allow",
    "deny"),
  validators.ConflictingObjectAttributes("allow",
    "local_preference"),
  validators.ConflictingObjectAttributes("allow",
    "metric"),
  validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "deny"),
  validators.ConflictingObjectAttributes("as_path",
    "local_preference"),
  validators.ConflictingObjectAttributes("as_path",
    "metric"),
  validators.ConflictingObjectAttributes("community",
    "deny"),
  validators.ConflictingObjectAttributes("community",
    "local_preference"),
  validators.ConflictingObjectAttributes("community",
    "metric"),
  validators.ConflictingObjectAttributes("deny",
    "local_preference"),
  validators.ConflictingObjectAttributes("deny",
    "metric"),
  validators.ConflictingObjectAttributes("local_preference",
    "metric")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"allow\",\"as_path\",\"community\",\"deny\",\"local_preference\",\"metric\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/allow/): complete subsection reference.

<a id="schema-rules--action--as_path"></a>

### as_path property

Type: `"string"`. Optional.

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Upstream description:

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [community](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/community/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/deny/): complete subsection reference.

<a id="schema-rules--action--local_preference"></a>

### local_preference property

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

Upstream description:

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-rules--action--metric"></a>

### metric property

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

Upstream description:

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [rules.action.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/allow/)
- [rules.action.community](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/community/)
- [rules.action.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/deny/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
