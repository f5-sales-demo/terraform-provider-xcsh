---
page_title: "rules.segment"
subcategory: ""
description: "rules.segment for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1483, "body_sha256": "sha256:f1df173ae3ddd9fc72a233c193bf73c16d760f44dcc8d4ea15a6c3fc173c51c8", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:segment:refs"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "documentation/resources/nat_policy/properties/rules/segment/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["rules", "segment"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.segment for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.segment

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- rules.segment

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
segment {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/): complete subsection reference.

## Next pages

- [rules.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
