---
page_title: "voltstack_cluster.storage_class_list.storage_classes"
subcategory: "Infrastructure"
description: "List of custom storage classes."
xcsh_docs: {"aliases": ["voltstack cluster storage class list storage classes"], "body_bytes": 3807, "body_sha256": "sha256:d2544c16fb4818d82574b5674280e3f07d4ecf7b95cd9937a4dfb8cb28192498", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list:storage_classes", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/storage_classes/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3033200112002300-0313122330223123-3101332322122332-3132022123203213-3220101230013330-1233322113303010-1100221002032000-1223012001322112", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-005.md", "relationships": [{"anchor": "schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name", "enforcement": "provider-schema", "group": "voltstack_cluster.storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list:storage_classes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "storage_class_list", "storage_classes"], "schema_version": 1, "sections": [{"aliases": ["voltstack cluster storage class list storage classes default storage class"], "anchor": "schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class", "description": "Make this storage class default storage class for the K8s cluster.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list:storage_classes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "storage_class_list", "storage_classes", "default_storage_class"], "syntax": "attribute", "type": "bool"}, {"aliases": ["voltstack cluster storage class list storage classes storage class name"], "anchor": "schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name", "description": "Name of the storage class as it will appear in K8s.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list:storage_classes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "storage_class_list", "storage_classes", "storage_class_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/storage_classes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of custom storage classes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.storage_class_list.storage_classes

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/)
- voltstack_cluster.storage_class_list.storage_classes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class"></a>

### default_storage_class property

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

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

<a id="schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name"></a>

### storage_class_name property

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
