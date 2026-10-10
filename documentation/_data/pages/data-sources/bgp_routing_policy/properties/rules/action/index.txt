---
page_title: "rules.action"
subcategory: ""
description: "Action to be enforced if the BGP route matches the rule."
xcsh_docs: {"aliases": ["rules action"], "body_bytes": 2890, "body_sha256": "sha256:aff442d003150dec0f4341d1ff6609a15c942380c57db5f046b5a00c54a88c98", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:allow", "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:community", "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:deny"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules", "path": "documentation/data-sources/bgp_routing_policy/properties/rules/action/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321", "registry_path": "docs/guides/data-sources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "sections": [{"aliases": ["rules action allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:allow", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules action as path"], "anchor": "schema-rules--action--as_path", "description": "Exclusive with AS-Path Prepending is generally used to influence incoming traffic.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "as_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules action community"], "anchor": "section", "description": "List of BGP communities.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:community", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "action", "community"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules action deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:deny", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "deny"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules action local preference"], "anchor": "schema-rules--action--local_preference", "description": "Exclusive with BGP Local Preference is generally used to influence outgoing traffic.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "local_preference"], "syntax": "attribute", "type": "number"}, {"aliases": ["rules action metric"], "anchor": "schema-rules--action--metric", "description": "Exclusive with The Multi-Exit Discriminator metric to indicate the preferred path to AS.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "metric"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Action to be enforced if the BGP route matches the rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
