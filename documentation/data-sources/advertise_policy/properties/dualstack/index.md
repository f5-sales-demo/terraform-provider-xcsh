---
page_title: "dualstack"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dualstack"], "body_bytes": 1299, "body_sha256": "sha256:dd589db9d7538b45c06c782ddb9c308c2010f793bb5cc3bc10b8c150d390b4fa", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:dualstack", "parent_id": "xcsh-docs:data-sources:advertise_policy:reference", "path": "documentation/data-sources/advertise_policy/properties/dualstack/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1300321132021101-2312230222232310-0310313232330230-3132121011011210-3012022120112200-0112311123223231-0322331101121122-2111101311031301", "registry_path": "docs/guides/data-sources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dualstack"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/dualstack/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dualstack

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- dualstack

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dualstack, ipv4, ipv6\] Enable this option

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

- [dualstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/dualstack/#section)
- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv4/#section)
- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/ipv6/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
