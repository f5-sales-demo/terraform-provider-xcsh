---
page_title: "mum_action.skip_processing"
subcategory: ""
description: "mum_action.skip_processing for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:ee8bfa51eeee270c5931e4c063a66d4f23205f413689e18f5a193a828e7fd53a", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action", "path": "docs/guides/resources--service_policy_rule--properties--mum_action--skip_processing.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mum_action", "skip_processing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/mum_action/skip_processing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mum_action.skip_processing for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mum_action.skip_processing

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [mum_action](resources--service_policy_rule--properties--mum_action.md)
- mum_action.skip_processing

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
skip_processing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [mum_action](resources--service_policy_rule--properties--mum_action.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
