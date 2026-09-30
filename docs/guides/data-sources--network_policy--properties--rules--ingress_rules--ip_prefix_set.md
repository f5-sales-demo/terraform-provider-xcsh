---
page_title: "rules.ingress_rules.ip_prefix_set"
subcategory: "Security"
description: "rules.ingress_rules.ip_prefix_set for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1158, "body_sha256": "sha256:0cc4e534c285988eddf79145f60087e3afafdeb4dfe6a355d7e145caa9ec1f49", "canonical_id": "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:network_policy:properties:rules:ingress_rules", "path": "docs/guides/data-sources--network_policy--properties--rules--ingress_rules--ip_prefix_set.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ingress_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/ingress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ingress_rules.ip_prefix_set for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.ingress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md)
- [Property reference](data-sources--network_policy--reference.md)
- [rules](data-sources--network_policy--properties--rules.md)
- [rules.ingress_rules](data-sources--network_policy--properties--rules--ingress_rules.md)
- rules.ingress_rules.ip_prefix_set

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [ref](data-sources--network_policy--properties--rules--ingress_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [rules.ingress_rules.ip_prefix_set.ref](data-sources--network_policy--properties--rules--ingress_rules--ip_prefix_set--ref.md)
- [rules.ingress_rules](data-sources--network_policy--properties--rules--ingress_rules.md)
- [xcsh_network_policy](../data-sources/network_policy.md)
