---
page_title: "allow_all_usb"
subcategory: ""
description: "allow_all_usb for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1313, "body_sha256": "sha256:7db7b516a1fcd5392eca981c01a8cd62eec163015e201c397dfeb67e1af57592", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:allow_all_usb", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:allow_all_usb", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "docs/guides/resources--voltstack_site--properties--allow_all_usb.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_all_usb"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/allow_all_usb/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all_usb for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all_usb

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- allow_all_usb

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all\_usb, deny\_all\_usb, usb\_policy\] Configuration parameter for allow all usb.

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

- [allow_all_usb](resources--voltstack_site--properties--allow_all_usb.md#section)
- [deny_all_usb](resources--voltstack_site--properties--deny_all_usb.md#section)
- [usb_policy](resources--voltstack_site--properties--usb_policy.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_usb = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
