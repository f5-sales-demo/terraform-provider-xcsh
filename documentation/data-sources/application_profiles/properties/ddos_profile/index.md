---
page_title: "ddos_profile"
subcategory: ""
description: "BIG-IP DDoS Protection Rules."
xcsh_docs: {"aliases": ["ddos profile"], "body_bytes": 1250, "body_sha256": "sha256:a7cc39fb137a4f796ea3975d48866d699175175296aac813ae2e24d510d251da", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:ddos_profile:disable_ddos_mitigation", "xcsh-docs:data-sources:application_profiles:properties:ddos_profile:enable_ddos_mitigation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:ddos_profile", "parent_id": "xcsh-docs:data-sources:application_profiles:reference", "path": "documentation/data-sources/application_profiles/properties/ddos_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0221310331001010-1231120233222200-3103210111100320-0121030023132112-2012102312231220-0311210333123322-0133000211232001-0223320121031332", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile"], "schema_version": 1, "sections": [{"aliases": ["ddos profile disable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:ddos_profile:disable_ddos_mitigation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "disable_ddos_mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["ddos profile enable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:ddos_profile:enable_ddos_mitigation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/ddos_profile/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "BIG-IP DDoS Protection Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["application_profilesCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- ddos_profile

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Additional upstream details:

BIG-IP DDoS Protection Rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

## Direct properties

- [disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/disable_ddos_mitigation/): complete subsection reference.

- [enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/enable_ddos_mitigation/): complete subsection reference.
