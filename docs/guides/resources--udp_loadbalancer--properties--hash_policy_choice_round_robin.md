---
page_title: "hash_policy_choice_round_robin"
subcategory: ""
description: "hash_policy_choice_round_robin for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 990, "body_sha256": "sha256:a9756252caeb0f8ac79e3a1f184968a1350f2f988b31eaa08972c495a83910f6", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:hash_policy_choice_round_robin", "child_ids": [], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:hash_policy_choice_round_robin", "parent_id": "xcsh-docs:resources:udp_loadbalancer:reference", "path": "docs/guides/resources--udp_loadbalancer--properties--hash_policy_choice_round_robin.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["hash_policy_choice_round_robin"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/hash_policy_choice_round_robin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "hash_policy_choice_round_robin for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hash_policy_choice_round_robin

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- hash_policy_choice_round_robin

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for hash policy choice round robin.

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
hash_policy_choice_round_robin = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--udp_loadbalancer--reference.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
