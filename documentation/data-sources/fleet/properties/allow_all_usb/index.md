---
page_title: "allow_all_usb"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["allow all usb"], "body_bytes": 1324, "body_sha256": "sha256:5862283ab99ddab78b2627234954718965136f059a759760d4c31b2f07af4b88", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:allow_all_usb", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/allow_all_usb/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1020303212333102-1221132312100331-1230301230330311-0322102012012202-2302112223332023-1131133311222203-0200222121232321-1312210000123110", "registry_path": "docs/guides/data-sources--fleet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_all_usb"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/allow_all_usb/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all_usb

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- allow_all_usb

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all\_usb, deny\_all\_usb, usb\_policy\] Configuration parameter for allow all usb.

Additional upstream details:

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

- [allow_all_usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/allow_all_usb/#section)
- [deny_all_usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/deny_all_usb/#section)
- [usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/usb_policy/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
