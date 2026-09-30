---
page_title: "rules.ingress_rules.all_traffic"
subcategory: "Security"
description: "rules.ingress_rules.all_traffic for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1014, "body_sha256": "sha256:2244f2cb850f33b7a06ffb16eee36817dd1fb09256d6c8918642a107f9e82213", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "path": "docs/guides/resources--network_policy--properties--rules--ingress_rules--all_traffic.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ingress_rules", "all_traffic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/ingress_rules/all_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ingress_rules.all_traffic for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.ingress_rules.all_traffic

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- [rules](resources--network_policy--properties--rules.md)
- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- rules.ingress_rules.all_traffic

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

- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- [xcsh_network_policy](../resources/network_policy.md)
