---
page_title: "dualstack"
subcategory: ""
description: "dualstack for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1205, "body_sha256": "sha256:62c406c1bf9091400f8cee192b957774c39aadff27557f645b1bf752c5c76f6e", "canonical_id": "xcsh-docs:data-sources:advertise_policy:properties:dualstack", "child_ids": [], "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:dualstack", "parent_id": "xcsh-docs:data-sources:advertise_policy:reference", "path": "docs/guides/data-sources--advertise_policy--properties--dualstack.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dualstack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/dualstack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dualstack for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dualstack

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
- [Property reference](data-sources--advertise_policy--reference.md)
- dualstack

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dualstack, ipv4, ipv6\] Enable this option

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

OneOf alternatives in this subsection:

- [dualstack](data-sources--advertise_policy--properties--dualstack.md#section)
- [ipv4](data-sources--advertise_policy--properties--ipv4.md#section)
- [ipv6](data-sources--advertise_policy--properties--ipv6.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--advertise_policy--reference.md)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
