---
page_title: "allow_all_requests"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["allow all requests"], "body_bytes": 1705, "body_sha256": "sha256:10cf5d5315e70faa6b62bda1e1ef76745aed31bc6ed8a1d04b6a868888a9d8c6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:allow_all_requests", "parent_id": "xcsh-docs:data-sources:service_policy:reference", "path": "documentation/data-sources/service_policy/properties/allow_all_requests/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1310133010132113-2321211101112030-2032301102311223-3121100330132020-0312130331122112-2311320303111131-3110000220310101-2300003320213301", "registry_path": "docs/guides/data-sources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_all_requests"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/allow_all_requests/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["service_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all_requests

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- allow_all_requests

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

- [allow_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/allow_all_requests/#section)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/allow_list/#section)
- [deny_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/deny_all_requests/#section)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/deny_list/#section)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
