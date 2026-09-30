---
page_title: "policy_based_challenge.rule_list.rules.spec.ip_matcher"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.ip_matcher for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2283, "body_sha256": "sha256:c2a71b8e18b82d5e2125511a51591c310e062b30eb528886f4048d85099e4979", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher:prefix_sets"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.ip_matcher for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_based_challenge.rule_list.rules.spec.ip_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](data-sources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="schema-policy_based_challenge--rule_list--rules--spec--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher--prefix_sets.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher--prefix_sets.md)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
