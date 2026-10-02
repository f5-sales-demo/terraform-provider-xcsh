---
page_title: "limits.action_block"
subcategory: "Security"
description: "Action where a user is blocked from making further requests after exceeding rate limit threshold."
xcsh_docs: {"aliases": ["limits action block"], "body_bytes": 2506, "body_sha256": "sha256:167fa1ba350458ce33d9c1e44ff3ceef1052c98aac9f8461a6b8c7021f50fb38", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "parent_id": "xcsh-docs:resources:rate_limiter:properties:limits", "path": "documentation/resources/rate_limiter/properties/limits/action_block/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033", "registry_path": "docs/guides/resources--rate_limiter--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["limits", "action_block"], "schema_version": 1, "sections": [{"aliases": ["hours"], "anchor": "section", "description": "Input Duration Hours.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "hours"], "syntax": "block", "type": "object"}, {"aliases": ["minutes"], "anchor": "section", "description": "Input Duration Minutes.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "minutes"], "syntax": "block", "type": "object"}, {"aliases": ["seconds"], "anchor": "section", "description": "Input Duration Seconds.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "seconds"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/action_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.action_block

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/)
- limits.action_block

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/minutes/): complete subsection reference.

- [seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/seconds/): complete subsection reference.

## Next pages

- [limits.action_block.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/hours/)
- [limits.action_block.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/minutes/)
- [limits.action_block.seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/seconds/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
