---
page_title: "discovery_k8s.publish_info"
subcategory: ""
description: "K8s Configuration to publish VIPs."
xcsh_docs: {"aliases": ["discovery k8s publish info"], "body_bytes": 1696, "body_sha256": "sha256:97b76eb7736a4d52812ef5dd51b6c382a0f5bf8cad57972627df5ea81d88d9c4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish_fqdns"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "documentation/data-sources/discovery/properties/discovery_k8s/publish_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112", "registry_path": "docs/guides/data-sources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s publish info disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s publish info dns delegation"], "anchor": "section", "description": "Configuration parameter for dns delegation.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s publish info publish"], "anchor": "section", "description": "K8SPublishType.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "publish"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s publish info publish fqdns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "publish_fqdns"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/publish_info/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "K8s Configuration to publish VIPs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- discovery_k8s.publish_info

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Additional upstream details:

K8s Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/disable_spec/): complete subsection reference.

- [dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/dns_delegation/): complete subsection reference.

- [publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/publish/): complete subsection reference.

- [publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/): complete subsection reference.
