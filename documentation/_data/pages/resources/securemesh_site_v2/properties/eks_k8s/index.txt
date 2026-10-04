---
page_title: "eks_k8s"
subcategory: ""
description: "Kubernetes Provider Type."
xcsh_docs: {"aliases": ["eks k8s"], "body_bytes": 6797, "body_sha256": "sha256:17c037210c026f83bd542a0de0447b8f2364055d8652edcbc90d0ab1e8586492", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/eks_k8s/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s:ConflictingObjectAttributes:disable_anti_affinity,enable_anti_affinity", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s:ConflictingObjectAttributes:disable_anti_affinity,enable_anti_affinity", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s"], "schema_version": 1, "sections": [{"aliases": ["eks k8s deployment size"], "anchor": "schema-eks_k8s--deployment_size", "description": "Enum for Kubernetes deployment size OPTIONS - KUBERNETES_DEPLOYMENT_SIZE_MEDIUM: Medium Medium deployment size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most deployments. - KUBERNETES_DEPLOYMENT_SIZE_LARGE: Large Large deployment size with higher resource requirements (16 vCPU, 64 GB", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "deployment_size"], "syntax": "attribute", "type": "string"}, {"aliases": ["eks k8s disable anti affinity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "disable_anti_affinity"], "syntax": "attribute", "type": "object"}, {"aliases": ["eks k8s enable anti affinity"], "anchor": "section", "description": "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}], "schema_path": ["eks_k8s", "enable_anti_affinity"], "syntax": "block", "type": "object"}, {"aliases": ["eks k8s labels"], "anchor": "schema-eks_k8s--labels", "description": "Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses Kubernetes nodeSelector to schedule pods only on nodes with matching labels.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["eks k8s not managed"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:not_managed", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["eks_k8s", "not_managed"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Kubernetes Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":253,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"253\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"63\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":63,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 253,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "253",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "63",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 63,
      "minLength": 1,
      "type": "string"
    }
  },
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
