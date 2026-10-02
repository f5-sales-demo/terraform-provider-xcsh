---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_connect."
xcsh_docs: {"aliases": ["cloud connect"], "body_bytes": 22791, "body_sha256": "sha256:45f37e1a814bc9ceed48d836f02eff78b692f7f0ac01d9a8a522548b66535512", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site", "xcsh-docs:resources:cloud_connect:properties:segment", "xcsh-docs:resources:cloud_connect:properties:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:reference", "parent_id": "xcsh-docs:resources:cloud_connect:fundamentals", "path": "documentation/resources/cloud_connect/properties/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["aws provider"], "anchor": "section", "description": "Cloud Connect with AWS.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider"], "syntax": "block", "type": "object"}, {"aliases": ["azure vnet site"], "anchor": "section", "description": "Cloud Connect Azure VNet Site Type.", "document_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_vnet_site"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:cloud_connect:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["segment"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cloud_connect:properties:segment", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-segment--name", "enforcement": "provider-schema", "group": "segment:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:segment", "type": "requires"}], "schema_path": ["segment"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:cloud_connect:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_cloud_connect.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/): complete subsection reference.

- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Cloud Connect. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Cloud Connect is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/segment/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-annotations) |
| `aws_provider` | [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/#section) |
| `aws_provider.aws_tgw_site` | [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/#section) |
| `aws_provider.aws_tgw_site.cred` | [aws_provider.aws_tgw_site.cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/#section) |
| `aws_provider.aws_tgw_site.cred.name` | [aws_provider.aws_tgw_site.cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/#schema-aws_provider--aws_tgw_site--cred--name) |
| `aws_provider.aws_tgw_site.cred.namespace` | [aws_provider.aws_tgw_site.cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/#schema-aws_provider--aws_tgw_site--cred--namespace) |
| `aws_provider.aws_tgw_site.cred.tenant` | [aws_provider.aws_tgw_site.cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/#schema-aws_provider--aws_tgw_site--cred--tenant) |
| `aws_provider.aws_tgw_site.site` | [aws_provider.aws_tgw_site.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/site/#section) |
| `aws_provider.aws_tgw_site.site.name` | [aws_provider.aws_tgw_site.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/site/#schema-aws_provider--aws_tgw_site--site--name) |
| `aws_provider.aws_tgw_site.site.namespace` | [aws_provider.aws_tgw_site.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/site/#schema-aws_provider--aws_tgw_site--site--namespace) |
| `aws_provider.aws_tgw_site.site.tenant` | [aws_provider.aws_tgw_site.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/site/#schema-aws_provider--aws_tgw_site--site--tenant) |
| `aws_provider.aws_tgw_site.vpc_attachments` | [aws_provider.aws_tgw_site.vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--route_table_id) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--static_routes) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/all_route_tables/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/selective_route_tables/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/selective_route_tables/#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--selective_route_tables--route_table_id) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/labels/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/manual_routing/#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--vpc_id) |
| `azure_vnet_site` | [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/#section) |
| `azure_vnet_site.site` | [azure_vnet_site.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/site/#section) |
| `azure_vnet_site.site.name` | [azure_vnet_site.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/site/#schema-azure_vnet_site--site--name) |
| `azure_vnet_site.site.namespace` | [azure_vnet_site.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/site/#schema-azure_vnet_site--site--namespace) |
| `azure_vnet_site.site.tenant` | [azure_vnet_site.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/site/#schema-azure_vnet_site--site--tenant) |
| `azure_vnet_site.vnet_attachments` | [azure_vnet_site.vnet_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list` | [azure_vnet_site.vnet_attachments.vnet_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/route_tables/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/route_tables/#schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--route_table_id) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/custom_routing/route_tables/#schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--static_routes) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route` | [azure_vnet_site.vnet_attachments.vnet_list.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/all_route_tables/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/selective_route_tables/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/default_route/selective_route_tables/#schema-azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables--route_table_id) |
| `azure_vnet_site.vnet_attachments.vnet_list.labels` | [azure_vnet_site.vnet_attachments.vnet_list.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/labels/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.manual_routing` | [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/manual_routing/#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.subscription_id` | [azure_vnet_site.vnet_attachments.vnet_list.subscription_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/#schema-azure_vnet_site--vnet_attachments--vnet_list--subscription_id) |
| `azure_vnet_site.vnet_attachments.vnet_list.vnet_id` | [azure_vnet_site.vnet_attachments.vnet_list.vnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/#schema-azure_vnet_site--vnet_attachments--vnet_list--vnet_id) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/#schema-namespace) |
| `segment` | [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/segment/#section) |
| `segment.name` | [segment.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/segment/#schema-segment--name) |
| `segment.namespace` | [segment.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/segment/#schema-segment--namespace) |
| `segment.tenant` | [segment.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/segment/#schema-segment--tenant) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/)
- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/segment/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/timeouts/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
