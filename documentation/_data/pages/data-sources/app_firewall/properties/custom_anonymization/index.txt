---
page_title: "custom_anonymization"
subcategory: "Security"
description: "Anonymization settings which is a list of HTTP headers, parameters and cookies."
xcsh_docs: {"aliases": ["custom anonymization"], "body_bytes": 1579, "body_sha256": "sha256:95a910fae1c09b4c0db1d121d14c6142911dc7f8addf85a7c2d77486615cfbce", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/custom_anonymization/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config"], "anchor": "section", "description": "List of HTTP headers, cookies and query parameters whose values will be masked.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Anonymization settings which is a list of HTTP headers, parameters and cookies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["app_firewallCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- custom_anonymization

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

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

- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/#section)
- [default_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_anonymization/#section)
- [disable_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/disable_anonymization/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/): complete subsection reference.
