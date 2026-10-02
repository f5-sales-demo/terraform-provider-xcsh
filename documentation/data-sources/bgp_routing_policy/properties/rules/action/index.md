---
page_title: "rules.action"
subcategory: ""
description: "Action to be enforced if the BGP route matches the rule."
xcsh_docs: {"aliases": ["rules action"], "body_bytes": 4048, "body_sha256": "sha256:29731177e789273f492fb77f8268e7db2bf7a8a9d4c8f36984db6a33a8e74360", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:allow", "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:community", "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:deny"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules", "path": "documentation/data-sources/bgp_routing_policy/properties/rules/action/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321", "registry_path": "docs/guides/data-sources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "sections": [{"aliases": ["allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:allow", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["as path"], "anchor": "schema-rules--action--as_path", "description": "Exclusive with AS-Path Prepending is generally used to influence incoming traffic.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "as_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["community"], "anchor": "section", "description": "List of BGP communities.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:community", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "action", "community"], "syntax": "attribute", "type": "object"}, {"aliases": ["deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "deny"], "syntax": "attribute", "type": "object"}, {"aliases": ["local preference"], "anchor": "schema-rules--action--local_preference", "description": "Exclusive with BGP Local Preference is generally used to influence outgoing traffic.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "local_preference"], "syntax": "attribute", "type": "number"}, {"aliases": ["metric"], "anchor": "schema-rules--action--metric", "description": "Exclusive with The Multi-Exit Discriminator metric to indicate the preferred path to AS.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "metric"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action to be enforced if the BGP route matches the rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/)
- rules.action

<a id="section"></a>

Type: `"single"`. Computed.

Action to be enforced if the BGP route matches the rule.

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

## Direct properties

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/allow/): complete subsection reference.

<a id="schema-rules--action--as_path"></a>

### as_path property

Type: `"string"`. Computed.

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

- [community](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/community/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/deny/): complete subsection reference.

<a id="schema-rules--action--local_preference"></a>

### local_preference property

Type: `"number"`. Computed.

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

Type: `"number"`. Computed.

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

- [rules.action.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/allow/)
- [rules.action.community](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/community/)
- [rules.action.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/deny/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
