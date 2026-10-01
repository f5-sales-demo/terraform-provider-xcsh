---
page_title: "policy_based_challenge.default_temporary_blocking_parameters"
subcategory: "Load Balancing"
description: "policy_based_challenge.default_temporary_blocking_parameters for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1140, "body_sha256": "sha256:fda4f4a2e63e94b61811e01bede9cce4d0027c2fcec90237ca413c2b4c7064fe", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:default_temporary_blocking_parameters", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:default_temporary_blocking_parameters", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/resources--cdn_loadbalancer--properties--policy_based_challenge--default_temporary_blocking_parameters.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "default_temporary_blocking_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/default_temporary_blocking_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.default_temporary_blocking_parameters for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.default_temporary_blocking_parameters

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- policy_based_challenge.default_temporary_blocking_parameters

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
default_temporary_blocking_parameters = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
