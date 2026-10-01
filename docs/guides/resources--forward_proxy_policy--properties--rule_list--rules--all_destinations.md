---
page_title: "rule_list.rules.all_destinations"
subcategory: "Security"
description: "rule_list.rules.all_destinations for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1165, "body_sha256": "sha256:ceb52903673f29550bfdba509854694454627bcb83b45075674f8b74149dfca7", "canonical_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:all_destinations", "child_ids": [], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:all_destinations", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules", "path": "docs/guides/resources--forward_proxy_policy--properties--rule_list--rules--all_destinations.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "all_destinations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/rule_list/rules/all_destinations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.all_destinations for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.all_destinations

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
- [Property reference](resources--forward_proxy_policy--reference.md)
- [rule_list](resources--forward_proxy_policy--properties--rule_list.md)
- [rule_list.rules](resources--forward_proxy_policy--properties--rule_list--rules.md)
- rule_list.rules.all_destinations

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rule_list.rules](resources--forward_proxy_policy--properties--rule_list--rules.md)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
