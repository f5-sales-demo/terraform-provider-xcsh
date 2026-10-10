---
page_title: "re_acl.fast_acl_rules.action.protocol_policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["re acl fast acl rules action protocol policer action"], "body_bytes": 1332, "body_sha256": "sha256:eef919dd5e7d50f10161265561e9ee2f59e386eaf7ebd43978746a14af533cbb", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:action", "path": "documentation/data-sources/fast_acl/properties/re_acl/fast_acl_rules/action/protocol_policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3022303333123333-2130222103330000-1021102233030130-2311330100031233-0021100322313032-2103231320000001-3213101301221100-0110320211321322", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "action", "protocol_policer_action"], "schema_version": 1, "sections": [{"aliases": ["re acl fast acl rules action protocol policer action ref"], "anchor": "section", "description": "Reference to protocol policer object.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "action", "protocol_policer_action", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/fast_acl_rules/action/protocol_policer_action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_aclCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.action.protocol_policer_action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/)
- [re_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/action/)
- re_acl.fast_acl_rules.action.protocol_policer_action

<a id="section"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

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

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/action/protocol_policer_action/ref/): complete subsection reference.
