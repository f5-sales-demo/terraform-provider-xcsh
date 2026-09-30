---
page_title: "waf_exclusion"
subcategory: "Load Balancing"
description: "waf_exclusion for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1980, "body_sha256": "sha256:f2f7992007dee57d0d349df0e9aa969fd6b41cb8ec67756b85c2a5d17e4603bd", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/waf_exclusion/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["waf_exclusion"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/waf_exclusion/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# waf_exclusion

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- waf_exclusion

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

## Direct properties

- [waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/): complete subsection reference.

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_policy/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/)
- [waf_exclusion.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
