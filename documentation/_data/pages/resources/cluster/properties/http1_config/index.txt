---
page_title: "http1_config"
subcategory: ""
description: "HTTP/1.1 Protocol OPTIONS for upstream connections."
xcsh_docs: {"aliases": ["http1 config"], "body_bytes": 950, "body_sha256": "sha256:d6e74e40855ccd63d9326cd0556b66c0baaa613d1ae4588e2ab04ea99f488ef2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:cluster:properties:http1_config:header_transformation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/http1_config/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0212332320322030-3122311023101131-3132013323032230-0111223003223223-2130302302311213-0223311332033230-1210013301010022-2212200332003133", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http1_config"], "schema_version": 1, "sections": [{"aliases": ["http1 config header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http1_config", "header_transformation"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "HTTP/1.1 Protocol OPTIONS for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["clusterCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- http1_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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
http1_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/): complete subsection reference.
