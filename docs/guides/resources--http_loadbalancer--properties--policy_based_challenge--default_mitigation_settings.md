---
page_title: "policy_based_challenge.default_mitigation_settings"
subcategory: "Load Balancing"
description: "policy_based_challenge.default_mitigation_settings for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1117, "body_sha256": "sha256:ce1bea1ef9785c7841d8c00a6a7b00920b471cea574b5b74ff604f5f2b8acd8b", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/resources--http_loadbalancer--properties--policy_based_challenge--default_mitigation_settings.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "default_mitigation_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/default_mitigation_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.default_mitigation_settings for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.default_mitigation_settings

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- policy_based_challenge.default_mitigation_settings

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
default_mitigation_settings = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
