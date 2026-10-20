---
page_title: "advanced_options.disable_outlier_detection"
subcategory: "Load Balancing"
description: "Disables outlier detection for the origin pool."
xcsh_docs: {"aliases": ["advanced options disable outlier detection", "disable outlier detection"], "body_bytes": 1137, "body_sha256": "sha256:c0ddfc2e6e53fcae3f2cbce0bb9baf41c41b8c3b26c78989c6d104ae16bfcad3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule", "reviewed-summary"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_outlier_detection", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "documentation/resources/origin_pool/properties/advanced_options/disable_outlier_detection/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0303111031121330-3030232011000130-3000031202321210-3313200311012320-0233221123302133-3200321002133322-3222232213001122-3301202011203131", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "disable_outlier_detection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/disable_outlier_detection/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Disables outlier detection for the origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.disable_outlier_detection

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- advanced_options.disable_outlier_detection

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

Terraform syntax:

```terraform
disable_outlier_detection = {}
```

This is an empty object or choice marker. It has no direct properties.
