---
page_title: "f5_big_ip_aws_service.nodes.reserved_mgmt_subnet"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["f5 big ip aws service nodes reserved mgmt subnet"], "body_bytes": 1228, "body_sha256": "sha256:b813c685a3edc69bd0b22c5e83a51456be57b28126609848f6d846d5de8adb6b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/reserved_mgmt_subnet/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-2010110231023102-2320103003311321-1331301001120003-0133203323300021-2331102311120021-1311331010010103-1221300133122101-2113331113111133", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes", "reserved_mgmt_subnet"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/reserved_mgmt_subnet/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.nodes.reserved_mgmt_subnet

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- [f5_big_ip_aws_service.nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/)
- f5_big_ip_aws_service.nodes.reserved_mgmt_subnet

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved mgmt subnet.

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
reserved_mgmt_subnet = {}
```

This is an empty object or choice marker. It has no direct properties.
