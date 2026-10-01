---
page_title: "policy_based_challenge.rule_list.rules.spec.asn_matcher"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.asn_matcher for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1974, "body_sha256": "sha256:76073a34f6b3e939f0db4d6f297a56e9f4c54cd95bda091b7215dcc0e6e7bb26", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "docs/guides/resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.asn_matcher for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.asn_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher--asn_sets.md)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
