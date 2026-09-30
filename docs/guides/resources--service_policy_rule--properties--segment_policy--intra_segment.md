---
page_title: "segment_policy.intra_segment"
subcategory: ""
description: "segment_policy.intra_segment for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 966, "body_sha256": "sha256:5e8aedba860beb19045a17d47ff2b7065abf74e4adba409bb4f8ce1247c0fbe1", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "docs/guides/resources--service_policy_rule--properties--segment_policy--intra_segment.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_policy", "intra_segment"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/intra_segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_policy.intra_segment for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# segment_policy.intra_segment

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [segment_policy](resources--service_policy_rule--properties--segment_policy.md)
- segment_policy.intra_segment

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for intra segment.

Upstream description:

This can be used for messages where no values are needed.

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
intra_segment = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [segment_policy](resources--service_policy_rule--properties--segment_policy.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
