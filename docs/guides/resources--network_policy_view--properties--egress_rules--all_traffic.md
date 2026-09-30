---
page_title: "egress_rules.all_traffic"
subcategory: ""
description: "egress_rules.all_traffic for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 946, "body_sha256": "sha256:77813c85d3a2f85bc9eebdb04c9c912fff4e6f65956673fe7fcc6e0dc6ad6dca", "canonical_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "parent_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules", "path": "docs/guides/resources--network_policy_view--properties--egress_rules--all_traffic.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_rules", "all_traffic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/egress_rules/all_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_rules.all_traffic for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# egress_rules.all_traffic

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md)
- [Property reference](resources--network_policy_view--reference.md)
- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- egress_rules.all_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- [xcsh_network_policy_view](../resources/network_policy_view.md)
