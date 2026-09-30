---
page_title: "protocol_policer"
subcategory: ""
description: "protocol_policer for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1928, "body_sha256": "sha256:9b649f9d48017c0a4da2c02c03210ab6ec07196571fc77ec1a08538129ed4107", "canonical_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer", "child_ids": ["xcsh-docs:resources:protocol_policer:properties:protocol_policer:policer", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol"], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer", "parent_id": "xcsh-docs:resources:protocol_policer:reference", "path": "docs/guides/resources--protocol_policer--properties--protocol_policer.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protocol_policer

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md)
- [Property reference](resources--protocol_policer--reference.md)
- protocol_policer

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of L4 protocol match condition and associated traffic rate limits.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("policer")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protocol_policer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policer](resources--protocol_policer--properties--protocol_policer--policer.md): complete subsection reference.

- [protocol](resources--protocol_policer--properties--protocol_policer--protocol.md): complete subsection reference.

## Next pages

- [protocol_policer.policer](resources--protocol_policer--properties--protocol_policer--policer.md)
- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- [Property reference](resources--protocol_policer--reference.md)
- [xcsh_protocol_policer](../resources/protocol_policer.md)
