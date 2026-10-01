---
page_title: "data_guard_rules.any_domain"
subcategory: "Load Balancing"
description: "data_guard_rules.any_domain for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:710214399b51d319f3be1c192223b6e041c0d6f09fbd96bdaa8eca22ea48295c", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "path": "docs/guides/resources--cdn_loadbalancer--properties--data_guard_rules--any_domain.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["data_guard_rules", "any_domain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/data_guard_rules/any_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "data_guard_rules.any_domain for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_guard_rules.any_domain

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [data_guard_rules](resources--cdn_loadbalancer--properties--data_guard_rules.md)
- data_guard_rules.any_domain

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
any_domain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [data_guard_rules](resources--cdn_loadbalancer--properties--data_guard_rules.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
