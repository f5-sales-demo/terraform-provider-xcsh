---
page_title: "disable_log_anonymization"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable log anonymization"], "body_bytes": 1335, "body_sha256": "sha256:974de0a43be2562b4a1ba0a846ba137e94a3b3c7ccccb3b5850d8704d61b41fb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:disable_log_anonymization", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/disable_log_anonymization/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3010223011012221-3311031211012022-3113130032121323-1100233121200201-0022032322131333-2221122020212113-3102311212332322-0230212232301232", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_log_anonymization"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/disable_log_anonymization/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["fleetCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_log_anonymization

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- disable_log_anonymization

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

- [disable_log_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/disable_log_anonymization/#section)
- [enable_log_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/enable_log_anonymization/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
