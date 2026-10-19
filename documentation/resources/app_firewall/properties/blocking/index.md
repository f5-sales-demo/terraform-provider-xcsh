---
page_title: "blocking"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["blocking"], "body_bytes": 1204, "body_sha256": "sha256:642e0f429e9fb5f30e70cacffd04c4825da580c79adb64db315a077d8b5c107c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:blocking", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/blocking/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1221032300100122-3100233131013023-3021323101202122-2031013300301101-0212103322111331-0232232121030201-0010103312223320-0223001012200312", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocking"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/blocking/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["app_firewallCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocking

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- blocking

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: blocking, monitoring\] Enable this option

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

- [blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/blocking/#section)
- [monitoring](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/monitoring/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocking = {}
```

This is an empty object or choice marker. It has no direct properties.
