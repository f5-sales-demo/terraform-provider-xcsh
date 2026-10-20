---
page_title: "enable"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable"], "body_bytes": 797, "body_sha256": "sha256:fe6f2a01aaaeada4c7ddc4035f3682133ff5d3254edc3c0ae149da6c60d12460", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "id": "xcsh-docs:resources:segment:properties:enable", "parent_id": "xcsh-docs:resources:segment:reference", "path": "documentation/resources/segment/properties/enable/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3100223213111010-2103312210220322-2233203002011122-3223023302011331-2333321113330202-2322330231112103-1333211031112123-0331313121113330", "registry_path": "docs/guides/resources--segment--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/properties/enable/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["segmentCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/properties/)
- enable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
enable = {}
```

This is an empty object or choice marker. It has no direct properties.
