---
page_title: "policy_based_challenge.rule_list"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1166, "body_sha256": "sha256:0f40ead91ca21af79eb86643f0530404987749744326be185e1260c95e56307a", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](data-sources--cdn_loadbalancer--properties--policy_based_challenge.md)
- policy_based_challenge.rule_list

<a id="section"></a>

Type: `"single"`. Computed.

List of challenge rules to be used in policy based challenge.

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

- [rules](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge](data-sources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
