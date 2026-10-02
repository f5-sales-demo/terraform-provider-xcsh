---
page_title: "eks_k8s"
subcategory: ""
description: "Kubernetes Provider Type."
xcsh_docs: {"aliases": ["eks k8s"], "body_bytes": 5459, "body_sha256": "sha256:c81064a2f03665095e9b48b643c246d6ae81228e9e8c6037b0db5c51861c0c61", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/eks_k8s/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s:ConflictingObjectAttributes:disable_anti_affinity,enable_anti_affinity", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s:ConflictingObjectAttributes:disable_anti_affinity,enable_anti_affinity", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s"], "schema_version": 1, "sections": [{"aliases": ["deployment size"], "anchor": "schema-eks_k8s--deployment_size", "description": "Enum for Kubernetes deployment size OPTIONS - KUBERNETES_DEPLOYMENT_SIZE_MEDIUM: Medium Medium deployment size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most deployments. - KUBERNETES_DEPLOYMENT_SIZE_LARGE: Large Large deployment size with higher resource requirements (16 vCPU, 64 GB", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "deployment_size"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable anti affinity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "disable_anti_affinity"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable anti affinity"], "anchor": "section", "description": "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}], "schema_path": ["eks_k8s", "enable_anti_affinity"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-eks_k8s--labels", "description": "Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses Kubernetes nodeSelector to schedule pods only on nodes with matching labels.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["not managed"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["eks_k8s", "not_managed"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Kubernetes Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- eks_k8s

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Kubernetes Provider Type. Kubernetes Provider Type.

Upstream description:

Kubernetes Provider Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_anti_affinity",
    "enable_anti_affinity")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-anti_affinity_choice": "[\"disable_anti_affinity\",\"enable_anti_affinity\"]"
}
```

Terraform syntax:

```terraform
eks_k8s {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-eks_k8s--deployment_size"></a>

### deployment_size property

Type: `"string"`. Optional.

\[Enum: KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM|KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\] Enum for
Kubernetes deployment size OPTIONS - KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium Medium deployment
size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most deployments. -
KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large Large deployment size with higher resource.. Possible
values are \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`, \`KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\`.
Defaults to \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`.

Upstream description:

Enum for Kubernetes deployment size OPTIONS

&#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium

Medium deployment size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most
deployments. &#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large

Large deployment size with higher resource requirements (16 vCPU, 64 GB memory) for demanding
workloads requiring additional performance and capacity.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
  "enum": [
    "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_anti_affinity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/disable_anti_affinity/): complete subsection reference.

- [enable_anti_affinity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/): complete subsection reference.

<a id="schema-eks_k8s--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

Upstream description:

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/): complete subsection reference.

## Next pages

- [eks_k8s.disable_anti_affinity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/disable_anti_affinity/)
- [eks_k8s.enable_anti_affinity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/)
- [eks_k8s.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/not_managed/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
