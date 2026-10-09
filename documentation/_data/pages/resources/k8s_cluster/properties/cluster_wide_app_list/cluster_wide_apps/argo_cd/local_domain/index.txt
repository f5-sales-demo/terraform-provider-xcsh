---
page_title: "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain"
subcategory: ""
description: "Parameters required to enable local access."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps argo cd local domain"], "body_bytes": 4858, "body_sha256": "sha256:e77daf88762ef568e39959f35be6df3f33cbf4a061186f1e930bb9e84635f03b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:default_port", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--port", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain:ConflictingObjectAttributes:default_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain:ConflictingObjectAttributes:default_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:default_port", "type": "conflicts"}, {"anchor": "schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--local_domain", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain:RequiredObjectAttributes:local_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain"], "schema_version": 1, "sections": [{"aliases": ["cluster wide app list cluster wide apps argo cd local domain default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:default_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain", "default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps argo cd local domain local domain"], "anchor": "schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--local_domain", "description": "ArgoCD will be accessible at <site name>.<local domain>.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain", "local_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["cluster wide app list cluster wide apps argo cd local domain password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain", "password"], "syntax": "block", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps argo cd local domain port"], "anchor": "schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--port", "description": "Exclusive with Use custom ArgoCD port. Available port range is less than 65000 except reserved ports.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain", "port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Parameters required to enable local access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/)
- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Parameters required to enable local access.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("local_domain"),
  validators.ConflictingObjectAttributes("default_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
local_domain {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/default_port/): complete subsection reference.

<a id="schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--local_domain"></a>

### local_domain property

Type: `"string"`. Optional.

ArgoCD will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/password/): complete subsection reference.

<a id="schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```
