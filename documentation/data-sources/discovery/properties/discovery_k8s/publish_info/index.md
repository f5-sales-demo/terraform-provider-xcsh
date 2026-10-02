---
page_title: "discovery_k8s.publish_info"
subcategory: ""
description: "K8s Configuration to publish VIPs."
xcsh_docs: {"aliases": ["discovery k8s publish info"], "body_bytes": 2629, "body_sha256": "sha256:8fedf5ed9c8f05258ca350cc1c39258f8b5a0ecd075ca0dde23998c12ce994fa", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish_fqdns"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "documentation/data-sources/discovery/properties/discovery_k8s/publish_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns delegation"], "anchor": "section", "description": "Configuration parameter for dns delegation.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "syntax": "attribute", "type": "object"}, {"aliases": ["publish"], "anchor": "section", "description": "K8SPublishType.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "publish"], "syntax": "attribute", "type": "object"}, {"aliases": ["publish fqdns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "publish_fqdns"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/publish_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "K8s Configuration to publish VIPs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

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

## Next pages

- [discovery_k8s.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/disable_spec/)
- [discovery_k8s.publish_info.dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/dns_delegation/)
- [discovery_k8s.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/publish/)
- [discovery_k8s.publish_info.publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
