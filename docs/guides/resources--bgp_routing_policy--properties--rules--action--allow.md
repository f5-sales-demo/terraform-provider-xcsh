---
page_title: "rules.action.allow"
subcategory: ""
description: "rules.action.allow for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 964, "body_sha256": "sha256:bed578beb7c3551aa839466cb3e5476d7a161fe1237068387e3228f060244cfc", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "path": "docs/guides/resources--bgp_routing_policy--properties--rules--action--allow.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "action", "allow"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/action/allow/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action.allow for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.action.allow

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
- [Property reference](resources--bgp_routing_policy--reference.md)
- [rules](resources--bgp_routing_policy--properties--rules.md)
- [rules.action](resources--bgp_routing_policy--properties--rules--action.md)
- rules.action.allow

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

- [rules.action](resources--bgp_routing_policy--properties--rules--action.md)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
