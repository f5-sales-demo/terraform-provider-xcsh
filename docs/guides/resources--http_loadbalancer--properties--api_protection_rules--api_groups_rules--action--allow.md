---
page_title: "api_protection_rules.api_groups_rules.action.allow"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules.action.allow for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1406, "body_sha256": "sha256:84aab8ccd1d8178aacd18188d589d5890245fdc4ee7d7345d60381af1e50a355", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "path": "docs/guides/resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action--allow.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "action", "allow"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/allow/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules.action.allow for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.action.allow

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action.md)
- api_protection_rules.api_groups_rules.action.allow

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
allow = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
