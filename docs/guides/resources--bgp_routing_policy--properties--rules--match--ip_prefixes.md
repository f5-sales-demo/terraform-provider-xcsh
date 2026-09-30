---
page_title: "rules.match.ip_prefixes"
subcategory: ""
description: "rules.match.ip_prefixes for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1318, "body_sha256": "sha256:e976452d5f4cc186c27511f1414e17d15458b0e11d593d94eaefc88ed1274823", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes"], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "path": "docs/guides/resources--bgp_routing_policy--properties--rules--match--ip_prefixes.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "match", "ip_prefixes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match.ip_prefixes for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.match.ip_prefixes

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
- [Property reference](resources--bgp_routing_policy--reference.md)
- [rules](resources--bgp_routing_policy--properties--rules.md)
- [rules.match](resources--bgp_routing_policy--properties--rules--match.md)
- rules.match.ip_prefixes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IP prefix and prefix length range match condition.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefixes")}
```

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
ip_prefixes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [prefixes](resources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md): complete subsection reference.

## Next pages

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md)
- [rules.match](resources--bgp_routing_policy--properties--rules--match.md)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
