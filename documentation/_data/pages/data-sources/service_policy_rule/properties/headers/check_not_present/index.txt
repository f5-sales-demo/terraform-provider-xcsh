---
page_title: "headers.check_not_present"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["headers check not present"], "body_bytes": 978, "body_sha256": "sha256:ae51cdb6c53572a11d412830adb2999d06b6829d4da57a43e2b334af607fd827", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:headers:check_not_present", "parent_id": "xcsh-docs:data-sources:service_policy_rule:properties:headers", "path": "documentation/data-sources/service_policy_rule/properties/headers/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2300201011320213-2033232121332232-1111030330032001-0020303322123320-2031120222332230-1010323130120303-1213320023310111-0201220122320212", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["headers", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/headers/check_not_present/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# headers.check_not_present

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/)
- headers.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

This is an empty object or choice marker. It has no direct properties.
