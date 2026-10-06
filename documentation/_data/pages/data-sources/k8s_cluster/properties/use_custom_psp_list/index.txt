---
page_title: "use_custom_psp_list"
subcategory: ""
description: "List of active Pod security policies for a K8s cluster."
xcsh_docs: {"aliases": ["use custom psp list"], "body_bytes": 1350, "body_sha256": "sha256:38eafa0baeff5b0a705ef225c17715efddc5ed74bbc3f4c2f07bfd9eb43b44ca", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_psp_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/use_custom_psp_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0222103111213102-1022030213200111-2301201033101101-2013300120212213-2320221103200012-2233132210011313-3020211000121131-1213103021100330", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_psp_list"], "schema_version": 1, "sections": [{"aliases": ["use custom psp list pod security policies"], "anchor": "section", "description": "List of active Pod security policies for a K8s cluster.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["use_custom_psp_list", "pod_security_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/use_custom_psp_list/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of active Pod security policies for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_psp_list

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- use_custom_psp_list

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

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

- [use_custom_psp_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_psp_list/#section)
- [use_default_psp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_default_psp/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [pod_security_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/use_custom_psp_list/pod_security_policies/): complete subsection reference.
