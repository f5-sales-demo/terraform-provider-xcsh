---
page_title: "advanced_options.disable_lb_source_ip_persistence"
subcategory: "Load Balancing"
description: "Disables load balancer source IP persistence for the origin pool."
xcsh_docs: {"aliases": ["advanced options disable lb source ip persistence", "disable lb source ip persistence", "disable source ip persistence"], "body_bytes": 1278, "body_sha256": "sha256:df3bdaf71f37db581f82c42b49b69ae004b93ccbc8f081682ea76404ef34b40d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule", "reviewed-summary"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "documentation/resources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1110033310111221-0311103030303223-3211303210220123-3033011120133100-2031102313201310-2302221330013212-2211323331302113-2330023011311230", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "disable_lb_source_ip_persistence"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Disables load balancer source IP persistence for the origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.disable_lb_source_ip_persistence

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- advanced_options.disable_lb_source_ip_persistence

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

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
disable_lb_source_ip_persistence = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
