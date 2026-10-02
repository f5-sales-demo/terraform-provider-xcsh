---
page_title: "action"
subcategory: ""
description: "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic."
xcsh_docs: {"aliases": ["action"], "body_bytes": 3178, "body_sha256": "sha256:77b8d9e68941da07efab006af8ff7565abd023077365a3da7f33d45089b9a384", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:action", "parent_id": "xcsh-docs:resources:fast_acl_rule:reference", "path": "documentation/resources/fast_acl_rule/properties/action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232", "registry_path": "docs/guides/resources--fast_acl_rule--reference--group-001.md", "relationships": [{"anchor": "schema-action--simple_action", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "type": "conflicts"}, {"anchor": "schema-action--simple_action", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["action"], "schema_version": 1, "sections": [{"aliases": ["policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["action", "policer_action"], "syntax": "block", "type": "object"}, {"aliases": ["protocol policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["action", "protocol_policer_action"], "syntax": "block", "type": "object"}, {"aliases": ["simple action"], "anchor": "schema-action--simple_action", "description": "FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the traffic Forward the traffic.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["action", "simple_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# action

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("policer_action",
    "protocol_policer_action"),
  validators.ConflictingObjectAttributes("policer_action",
    "simple_action"),
  validators.ConflictingObjectAttributes("protocol_policer_action",
    "simple_action")}
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
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/): complete subsection reference.

- [protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/): complete subsection reference.

<a id="schema-action--simple_action"></a>

### simple_action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [action.policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/)
- [action.protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
