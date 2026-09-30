---
page_title: "reauth_timeout_hours"
subcategory: ""
description: "reauth_timeout_hours for xcsh_ike1."
xcsh_docs: {"aliases": [], "body_bytes": 1907, "body_sha256": "sha256:c0db93245c6eb388a7618f7b77365cdd48567ed55a7ddb5ba772abb581e507d4", "canonical_id": "xcsh-docs:resources:ike1:properties:reauth_timeout_hours", "child_ids": [], "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:reauth_timeout_hours", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "docs/guides/resources--ike1--properties--reauth_timeout_hours.md", "provider_name": "ike1", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["reauth_timeout_hours"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/reauth_timeout_hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "reauth_timeout_hours for xcsh_ike1.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# reauth_timeout_hours

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md)
- [Property reference](resources--ike1--reference.md)
- reauth_timeout_hours

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout hours.

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

Terraform syntax:

```terraform
reauth_timeout_hours {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-reauth_timeout_hours--duration"></a>

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

- [Property reference](resources--ike1--reference.md)
- [xcsh_ike1](../resources/ike1.md)
