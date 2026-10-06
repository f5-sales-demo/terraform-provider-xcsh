---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "Consul Configuration to publish VIPs."
xcsh_docs: {"aliases": ["discovery consul publish info"], "body_bytes": 1310, "body_sha256": "sha256:4aa5f669890b2c34faab2a1c4bb881d775fb0ff11351c22029b04eabb030fa60", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:publish"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "path": "documentation/data-sources/discovery/properties/discovery_consul/publish_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3022323202130331-1300101212130033-0001310133112233-0230110021230003-0022213333130131-1230010030020111-1232120213113231-3303010023321313", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "sections": [{"aliases": ["discovery consul publish info disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "publish_info", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery consul publish info publish"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info:publish", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "publish_info", "publish"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Consul Configuration to publish VIPs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.publish_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- discovery_consul.publish_info

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Additional upstream details:

Consul Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/disable_spec/): complete subsection reference.

- [publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/publish/): complete subsection reference.
