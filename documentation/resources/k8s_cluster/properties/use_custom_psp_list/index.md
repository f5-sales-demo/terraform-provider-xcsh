---
page_title: "use_custom_psp_list"
subcategory: ""
description: "List of active Pod security policies for a K8s cluster."
xcsh_docs: {"aliases": ["use custom psp list"], "body_bytes": 1653, "body_sha256": "sha256:41b4b02e24da9847b2f72b4f771b64bb676bfa7c9b7842ae9796ac94f6fab878", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/use_custom_psp_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3131020221100132-1003100033303012-3130330301313330-1033003233300201-1331103000203002-0033021330203230-3013312030313301-0011311101333322", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_custom_psp_list:RequiredObjectAttributes:pod_security_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_psp_list"], "schema_version": 1, "sections": [{"aliases": ["use custom psp list pod security policies"], "anchor": "section", "description": "List of active Pod security policies for a K8s cluster.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-use_custom_psp_list--pod_security_policies--name", "enforcement": "provider-schema", "group": "use_custom_psp_list.pod_security_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list:pod_security_policies", "type": "requires"}], "schema_path": ["use_custom_psp_list", "pod_security_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_psp_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of active Pod security policies for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
