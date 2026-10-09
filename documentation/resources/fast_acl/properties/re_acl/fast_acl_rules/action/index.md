---
page_title: "re_acl.fast_acl_rules.action"
subcategory: ""
description: "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic."
xcsh_docs: {"aliases": ["re acl fast acl rules action"], "body_bytes": 2954, "body_sha256": "sha256:12ac1ffa24f7e86410fc9d5696c3e8dbafdc6547560ffae96db32d3f3bf1e25e", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "path": "documentation/resources/fast_acl/properties/re_acl/fast_acl_rules/action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "schema-re_acl--fast_acl_rules--action--simple_action", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "type": "conflicts"}, {"anchor": "schema-re_acl--fast_acl_rules--action--simple_action", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "action"], "schema_version": 1, "sections": [{"aliases": ["re acl fast acl rules action policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "action", "policer_action"], "syntax": "block", "type": "object"}, {"aliases": ["re acl fast acl rules action protocol policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "action", "protocol_policer_action"], "syntax": "block", "type": "object"}, {"aliases": ["re acl fast acl rules action simple action"], "anchor": "schema-re_acl--fast_acl_rules--action--simple_action", "description": "FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the traffic Forward the traffic.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ALLOW", "DENY"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "action", "simple_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/action/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["fast_aclCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/)
- re_acl.fast_acl_rules.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/policer_action/): complete subsection reference.

- [protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/protocol_policer_action/): complete subsection reference.

<a id="schema-re_acl--fast_acl_rules--action--simple_action"></a>

### simple_action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
