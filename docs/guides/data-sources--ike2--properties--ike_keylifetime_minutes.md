---
page_title: "ike_keylifetime_minutes"
subcategory: ""
description: "ike_keylifetime_minutes for xcsh_ike2."
xcsh_docs: {"aliases": [], "body_bytes": 1651, "body_sha256": "sha256:6e445dd7256236d536377228a3ba0ac9404593c4a16c4739f2a3f444de6ff1bc", "canonical_id": "xcsh-docs:data-sources:ike2:properties:ike_keylifetime_minutes", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike2:properties:ike_keylifetime_minutes", "parent_id": "xcsh-docs:data-sources:ike2:reference", "path": "docs/guides/data-sources--ike2--properties--ike_keylifetime_minutes.md", "provider_name": "ike2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ike_keylifetime_minutes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike2/properties/ike_keylifetime_minutes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ike_keylifetime_minutes for xcsh_ike2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ike_keylifetime_minutes

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md)
- [Property reference](data-sources--ike2--reference.md)
- ike_keylifetime_minutes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

<a id="schema-ike_keylifetime_minutes--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

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

- [Property reference](data-sources--ike2--reference.md)
- [xcsh_ike2](../data-sources/ike2.md)
