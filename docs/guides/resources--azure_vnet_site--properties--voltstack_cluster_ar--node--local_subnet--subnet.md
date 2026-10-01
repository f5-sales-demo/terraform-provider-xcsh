---
page_title: "voltstack_cluster_ar.node.local_subnet.subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.node.local_subnet.subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 4184, "body_sha256": "sha256:3b7bbe3a1e35e65b9e0d0188ee2c3c85b44fb004d3b38e3eaa993dae9584552c", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet:vnet_resource_group"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "node", "local_subnet", "subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.node.local_subnet.subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.node.local_subnet.subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--properties--voltstack_cluster_ar--node.md)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet.md)
- voltstack_cluster_ar.node.local_subnet.subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster_ar--node--local_subnet--subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-voltstack_cluster_ar--node--local_subnet--subnet--subnet_resource_grp"></a>

### subnet_resource_grp property

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet--vnet_resource_group.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet--vnet_resource_group.md)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
