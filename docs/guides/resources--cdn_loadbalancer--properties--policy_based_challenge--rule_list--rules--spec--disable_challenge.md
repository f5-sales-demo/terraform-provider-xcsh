---
page_title: "policy_based_challenge.rule_list.rules.spec.disable_challenge"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.disable_challenge for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1580, "body_sha256": "sha256:3eb62b8e392cee7b82a890fd9cfa14fdfe71160774263fda843134d65d727ffe", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "docs/guides/resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--disable_challenge.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "disable_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/disable_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.disable_challenge for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.disable_challenge

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
