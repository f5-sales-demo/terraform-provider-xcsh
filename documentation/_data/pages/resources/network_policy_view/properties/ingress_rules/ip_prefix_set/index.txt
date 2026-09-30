---
page_title: "ingress_rules.ip_prefix_set"
subcategory: ""
description: "ingress_rules.ip_prefix_set for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1539, "body_sha256": "sha256:3a71f9d0c881658493220d63d60fa34daaf116a7cc8c8bbc56ff0369bf396d93", "child_ids": ["xcsh-docs:resources:network_policy_view:properties:ingress_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules", "path": "documentation/resources/network_policy_view/properties/ingress_rules/ip_prefix_set/index.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/ingress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_rules.ip_prefix_set for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- ingress_rules.ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [ingress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
