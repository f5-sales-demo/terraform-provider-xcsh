---
page_title: "voltstack_cluster.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["voltstack cluster active forward proxy policies"], "body_bytes": 1593, "body_sha256": "sha256:961f69fe07ec8e5f54ac1d177fdabb14cd8f38d3d0dc6a7a5a4f0179194b084a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "path": "documentation/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3032122300113332-1102211212202213-1102312313001033-3213100122003112-3330003230032132-2313203233312312-0132210133012311-1000233033322202", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["voltstack_cluster", "active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.active_forward_proxy_policies

<a id="section"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

## Direct properties

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.

## Next pages

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
