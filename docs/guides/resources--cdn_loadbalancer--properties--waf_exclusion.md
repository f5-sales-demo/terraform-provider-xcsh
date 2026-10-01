---
page_title: "waf_exclusion"
subcategory: "Load Balancing"
description: "waf_exclusion for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1671, "body_sha256": "sha256:d17ba59b88c99716e2a1b4f3f81ab416cdd423a96237b01f327c89a04971e896", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--waf_exclusion.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/waf_exclusion/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
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

- [waf_exclusion_inline_rules](resources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md): complete subsection reference.

- [waf_exclusion_policy](resources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_policy.md): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md)
- [waf_exclusion.waf_exclusion_policy](resources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_policy.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
