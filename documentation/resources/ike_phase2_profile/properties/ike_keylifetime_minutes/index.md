---
page_title: "ike_keylifetime_minutes"
subcategory: ""
description: "Set IKE Key Lifetime in minutes."
xcsh_docs: {"aliases": ["ike keylifetime minutes"], "body_bytes": 2345, "body_sha256": "sha256:d0e02027a6fad84c41071bb4d22815efef2e451dc48ab0788d7506a10fca2de4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase2_profile:properties:ike_keylifetime_minutes", "parent_id": "xcsh-docs:resources:ike_phase2_profile:reference", "path": "documentation/resources/ike_phase2_profile/properties/ike_keylifetime_minutes/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3013020132011201-3031230131033311-1222301300203032-0032302231133013-2323200223031321-3010001230222303-0120122032301323-1023103330212020", "registry_path": "docs/guides/resources--ike_phase2_profile--reference--group-001.md", "relationships": [{"anchor": "schema-ike_keylifetime_minutes--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike_phase2_profile:properties:ike_keylifetime_minutes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_minutes"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-ike_keylifetime_minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike_phase2_profile:properties:ike_keylifetime_minutes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/properties/ike_keylifetime_minutes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set IKE Key Lifetime in minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_minutes

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- ike_keylifetime_minutes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

Terraform syntax:

```terraform
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ike_keylifetime_minutes--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
