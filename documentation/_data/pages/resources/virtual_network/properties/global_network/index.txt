---
page_title: "global_network"
subcategory: "Networking"
description: "Select the global virtual-network scope for connectivity across participating sites."
xcsh_docs: {"aliases": ["global network"], "body_bytes": 1448, "body_sha256": "sha256:eee3c2b14d29a6e19fefd3083687d6962b139e346de6c2eedf7f9d5e0ab892bd", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:global_network", "parent_id": "xcsh-docs:resources:virtual_network:reference", "path": "documentation/resources/virtual_network/properties/global_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2022100300112003-1220202113121001-1131003021302131-3131310112012213-0021220111013333-0303333111203120-3031131131222030-2012313201111120", "registry_path": "docs/guides/resources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["global_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/global_network/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Select the global virtual-network scope for connectivity across participating sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# global_network

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- global_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

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

- [global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/global_network/#section)
- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/site_local_inside_network/#section)
- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/site_local_network/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_network = {}
```

This is an empty object or choice marker. It has no direct properties.
