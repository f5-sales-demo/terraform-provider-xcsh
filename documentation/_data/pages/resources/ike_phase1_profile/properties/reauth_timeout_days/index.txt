---
page_title: "reauth_timeout_days"
subcategory: ""
description: "Set Duration in days."
xcsh_docs: {"aliases": ["duration", "operation timeout", "reauth timeout days"], "body_bytes": 2306, "body_sha256": "sha256:38c9a5b95b02eafd59b6072d5a54ab97aeaccc4e13cb195428ab6b3568807557", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_timeout_days", "parent_id": "xcsh-docs:resources:ike_phase1_profile:reference", "path": "documentation/resources/ike_phase1_profile/properties/reauth_timeout_days/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3022232010212211-3311312110222112-2212001323132301-3322123113101312-3021320311332100-3233231203303320-2011010223222322-1200322130012222", "registry_path": "docs/guides/resources--ike_phase1_profile--reference--group-001.md", "relationships": [{"anchor": "schema-reauth_timeout_days--duration", "enforcement": "provider-schema", "group": "reauth_timeout_days:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_timeout_days", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_timeout_days"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-reauth_timeout_days--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_timeout_days", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["reauth_timeout_days", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/properties/reauth_timeout_days/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Set Duration in days.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_timeout_days

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- reauth_timeout_days

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

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
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-reauth_timeout_days--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
