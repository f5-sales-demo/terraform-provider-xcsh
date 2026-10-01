---
page_title: "segment_policy.src_any"
subcategory: ""
description: "segment_policy.src_any for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:f2144b15f162aef268f9a726ff6f6287af7e3cf49474660292f1d2969732c778", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "docs/guides/resources--service_policy_rule--properties--segment_policy--src_any.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_policy", "src_any"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/src_any/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_policy.src_any for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy.src_any

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [segment_policy](resources--service_policy_rule--properties--segment_policy.md)
- segment_policy.src_any

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
src_any = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [segment_policy](resources--service_policy_rule--properties--segment_policy.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
