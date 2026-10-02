---
page_title: "rules"
subcategory: ""
description: "A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as rules are applied top to bottom."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 2545, "body_sha256": "sha256:60c6b407dc88edee6c26217a2218ca6de8c2fb2f70fbf29850b2e05a9ee69d76", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "xcsh-docs:resources:bgp_routing_policy:properties:rules:match"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules", "parent_id": "xcsh-docs:resources:bgp_routing_policy:reference", "path": "documentation/resources/bgp_routing_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3100200231201222-3331030201031021-0022112323211122-0210221002212021-3303031333321000-0101032300222002-1212202322312103-0013010332102321", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "Action to be enforced if the BGP route matches the rule.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,as_path", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--as_path", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--local_preference", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:local_preference,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "schema-rules--action--metric", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:local_preference,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,as_path", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:as_path,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:community,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,local_preference", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.action:ConflictingObjectAttributes:deny,metric", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny", "type": "conflicts"}], "schema_path": ["rules", "action"], "syntax": "block", "type": "object"}, {"aliases": ["match"], "anchor": "section", "description": "Predicates which have to match information in route for action to be applied.", "document_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--match--as_path", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "type": "conflicts"}, {"anchor": "schema-rules--match--as_path", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,community", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:community,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:as_path,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.match:ConflictingObjectAttributes:community,ip_prefixes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "type": "conflicts"}], "schema_path": ["rules", "match"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as rules are applied top to bottom.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Upstream description:

A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
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

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/): complete subsection reference.

- [match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/): complete subsection reference.

## Next pages

- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
