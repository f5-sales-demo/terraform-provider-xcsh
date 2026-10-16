---
page_title: "where.virtual_site.disable_internet_vip"
subcategory: "Networking"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["where virtual site disable internet vip"], "body_bytes": 1114, "body_sha256": "sha256:08031cbf5430d95144821342c4f1784100e389727cd50edf43bc1174feace2b7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:where:virtual_site:disable_internet_vip", "parent_id": "xcsh-docs:resources:endpoint:properties:where:virtual_site", "path": "documentation/resources/endpoint/properties/where/virtual_site/disable_internet_vip/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0321313020021211-2232003203110323-2200133123320100-3103231201001313-1021103202102332-0223223222112102-3221333020303003-2111211323022001", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_site", "disable_internet_vip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/where/virtual_site/disable_internet_vip/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["endpointCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_site.disable_internet_vip

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_site/)
- where.virtual_site.disable_internet_vip

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.
