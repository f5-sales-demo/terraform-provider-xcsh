---
page_title: "ike_keylifetime_minutes"
subcategory: ""
description: "Set IKE Key Lifetime in minutes."
xcsh_docs: {"aliases": ["ike keylifetime minutes"], "body_bytes": 2261, "body_sha256": "sha256:ffca5652fec7665af8d8f838af8cd99e1db704f29f13004ccab53ac35e930643", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_minutes", "parent_id": "xcsh-docs:resources:ike2:reference", "path": "documentation/resources/ike2/properties/ike_keylifetime_minutes/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3100020103332033-2200320203233111-0013310012201310-1000300323210330-3312223323313112-2110101022322212-2102301010112130-3110302112233330", "registry_path": "docs/guides/resources--ike2--reference--group-001.md", "relationships": [{"anchor": "schema-ike_keylifetime_minutes--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_minutes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ike_keylifetime_minutes"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-ike_keylifetime_minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_minutes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ike_keylifetime_minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/properties/ike_keylifetime_minutes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Set IKE Key Lifetime in minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_minutes

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
