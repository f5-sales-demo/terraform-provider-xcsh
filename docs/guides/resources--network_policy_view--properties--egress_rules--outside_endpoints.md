---
page_title: "egress_rules.outside_endpoints"
subcategory: ""
description: "egress_rules.outside_endpoints for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1041, "body_sha256": "sha256:5b355cdd2ee74bd132bf7ffeb7093688e7da0bbe1d242615fb224cb8f6519a0a", "canonical_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "parent_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules", "path": "docs/guides/resources--network_policy_view--properties--egress_rules--outside_endpoints.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_rules", "outside_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/egress_rules/outside_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_rules.outside_endpoints for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules.outside_endpoints

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md)
- [Property reference](resources--network_policy_view--reference.md)
- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- egress_rules.outside_endpoints

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
outside_endpoints = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- [xcsh_network_policy_view](../resources/network_policy_view.md)
