---
page_title: "ike_keylifetime_hours"
subcategory: ""
description: "Input Hours."
xcsh_docs: {"aliases": ["ike keylifetime hours"], "body_bytes": 3004, "body_sha256": "sha256:8e436e20f8210a5a7ca2227228b84ae9836bf6cb04fcd02be12fbe6238559af0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_hours", "parent_id": "xcsh-docs:resources:ike_phase1_profile:reference", "path": "documentation/resources/ike_phase1_profile/properties/ike_keylifetime_hours/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0202020210020100-0010323001013223-3102230231101231-0332020330300301-3131031101323222-2122211233201101-2032303103300212-2333331323213101", "registry_path": "docs/guides/resources--ike_phase1_profile--reference--group-001.md", "relationships": [{"anchor": "schema-ike_keylifetime_hours--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_hours", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_hours"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-ike_keylifetime_hours--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike_phase1_profile:properties:ike_keylifetime_hours", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_hours", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/properties/ike_keylifetime_hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Input Hours.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_hours

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- ike_keylifetime_hours

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
```

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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/ike_keylifetime_hours/#section)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/ike_keylifetime_minutes/#section)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/use_default_keylifetime/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ike_keylifetime_hours--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
