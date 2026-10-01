---
page_title: "waf_exclusion_rules.waf_skip_processing"
subcategory: ""
description: "waf_exclusion_rules.waf_skip_processing for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1096, "body_sha256": "sha256:2dddd96b7cc34b53b89a57c22657ad6403080f176828a11bd17d416208528de5", "canonical_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:waf_skip_processing", "child_ids": [], "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:waf_skip_processing", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules", "path": "docs/guides/resources--waf_exclusion_policy--properties--waf_exclusion_rules--waf_skip_processing.md", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion_rules", "waf_skip_processing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/properties/waf_exclusion_rules/waf_skip_processing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion_rules.waf_skip_processing for xcsh_waf_exclusion_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules.waf_skip_processing

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
- [Property reference](resources--waf_exclusion_policy--reference.md)
- [waf_exclusion_rules](resources--waf_exclusion_policy--properties--waf_exclusion_rules.md)
- waf_exclusion_rules.waf_skip_processing

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
waf_skip_processing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [waf_exclusion_rules](resources--waf_exclusion_policy--properties--waf_exclusion_rules.md)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
