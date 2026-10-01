---
page_title: "ipv4"
subcategory: ""
description: "ipv4 for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 911, "body_sha256": "sha256:353e328843b7d118a32e41fe42c75037e19a55a62c2063083eb373583c567ac8", "canonical_id": "xcsh-docs:resources:advertise_policy:properties:ipv4", "child_ids": [], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:ipv4", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "docs/guides/resources--advertise_policy--properties--ipv4.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipv4 for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipv4

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md)
- [Property reference](resources--advertise_policy--reference.md)
- ipv4

<a id="section"></a>

Type: `["object", {}]`. Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

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
ipv4 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--advertise_policy--reference.md)
- [xcsh_advertise_policy](../resources/advertise_policy.md)
