---
page_title: "allow_all_usb"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["allow all usb"], "body_bytes": 1593, "body_sha256": "sha256:64ed472df92318137ecf2b55e950d02d74faaef289c830b650d3b24d9a1ac5e3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:allow_all_usb", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/allow_all_usb/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0233020232022221-0012000111130132-3331103020110010-1132313232302301-3302110232202113-3033221010233312-1302312102121301-2003110120333123", "registry_path": "docs/guides/resources--fleet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_all_usb"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/allow_all_usb/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all_usb

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
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

- [allow_all_usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/allow_all_usb/#section)
- [deny_all_usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/deny_all_usb/#section)
- [usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/usb_policy/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_usb = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
