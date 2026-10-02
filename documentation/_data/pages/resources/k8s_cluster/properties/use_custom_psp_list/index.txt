---
page_title: "use_custom_psp_list"
subcategory: ""
description: "List of active Pod security policies for a K8s cluster."
xcsh_docs: {"aliases": ["use custom psp list"], "body_bytes": 2113, "body_sha256": "sha256:e5bfcc376bc59a04f87e05ae4f44c6feca3b3299c3e1b0859ec9a4748d5ab451", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/use_custom_psp_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3131020221100132-1003100033303012-3130330301313330-1033003233300201-1331103000203002-0033021330203230-3013312030313301-0011311101333322", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_custom_psp_list:RequiredObjectAttributes:pod_security_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_psp_list"], "schema_version": 1, "sections": [{"aliases": ["pod security policies"], "anchor": "section", "description": "List of active Pod security policies for a K8s cluster.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-use_custom_psp_list--pod_security_policies--name", "enforcement": "provider-schema", "group": "use_custom_psp_list.pod_security_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "type": "requires"}], "schema_path": ["use_custom_psp_list", "pod_security_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_psp_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of active Pod security policies for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_psp_list

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- use_custom_psp_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Upstream description:

List of active Pod security policies for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("pod_security_policies")}
```

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

- [use_custom_psp_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_psp_list/#section)
- [use_default_psp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_default_psp/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_psp_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [pod_security_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_psp_list/pod_security_policies/): complete subsection reference.

## Next pages

- [use_custom_psp_list.pod_security_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_psp_list/pod_security_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
