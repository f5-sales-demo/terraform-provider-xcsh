---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 17896, "body_sha256": "sha256:cee9a777e1ff1fc44d31ea523c27ac84e1a2001b34b39ead0cedf694260ee8e1", "canonical_id": "xcsh-docs:data-sources:cloud_connect:reference", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:aws_provider", "xcsh-docs:data-sources:cloud_connect:properties:azure_vnet_site", "xcsh-docs:data-sources:cloud_connect:properties:segment"], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:reference", "parent_id": "xcsh-docs:data-sources:cloud_connect:fundamentals", "path": "docs/guides/data-sources--cloud_connect--reference.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [aws_provider](data-sources--cloud_connect--properties--aws_provider.md): complete subsection reference.

- [azure_vnet_site](data-sources--cloud_connect--properties--azure_vnet_site.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the CloudConnect.

Upstream description:

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

Name of the CloudConnect.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the CloudConnect exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [segment](data-sources--cloud_connect--properties--segment.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_connect--reference.md#schema-annotations) |
| `aws_provider` | [aws_provider](data-sources--cloud_connect--properties--aws_provider.md#section) |
| `aws_provider.aws_tgw_site` | [aws_provider.aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md#section) |
| `aws_provider.aws_tgw_site.cred` | [aws_provider.aws_tgw_site.cred](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--cred.md#section) |
| `aws_provider.aws_tgw_site.cred.name` | [aws_provider.aws_tgw_site.cred.name](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--cred.md#schema-aws_provider--aws_tgw_site--cred--name) |
| `aws_provider.aws_tgw_site.cred.namespace` | [aws_provider.aws_tgw_site.cred.namespace](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--cred.md#schema-aws_provider--aws_tgw_site--cred--namespace) |
| `aws_provider.aws_tgw_site.cred.tenant` | [aws_provider.aws_tgw_site.cred.tenant](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--cred.md#schema-aws_provider--aws_tgw_site--cred--tenant) |
| `aws_provider.aws_tgw_site.site` | [aws_provider.aws_tgw_site.site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--site.md#section) |
| `aws_provider.aws_tgw_site.site.name` | [aws_provider.aws_tgw_site.site.name](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--site.md#schema-aws_provider--aws_tgw_site--site--name) |
| `aws_provider.aws_tgw_site.site.namespace` | [aws_provider.aws_tgw_site.site.namespace](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--site.md#schema-aws_provider--aws_tgw_site--site--namespace) |
| `aws_provider.aws_tgw_site.site.tenant` | [aws_provider.aws_tgw_site.site.tenant](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--site.md#schema-aws_provider--aws_tgw_site--site--tenant) |
| `aws_provider.aws_tgw_site.vpc_attachments` | [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables.md#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--route_table_id) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables.md#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--static_routes) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--all_route_tables.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--selective_route_tables.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--selective_route_tables.md#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--selective_route_tables--route_table_id) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--manual_routing.md#section) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md#schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--vpc_id) |
| `azure_vnet_site` | [azure_vnet_site](data-sources--cloud_connect--properties--azure_vnet_site.md#section) |
| `azure_vnet_site.site` | [azure_vnet_site.site](data-sources--cloud_connect--properties--azure_vnet_site--site.md#section) |
| `azure_vnet_site.site.name` | [azure_vnet_site.site.name](data-sources--cloud_connect--properties--azure_vnet_site--site.md#schema-azure_vnet_site--site--name) |
| `azure_vnet_site.site.namespace` | [azure_vnet_site.site.namespace](data-sources--cloud_connect--properties--azure_vnet_site--site.md#schema-azure_vnet_site--site--namespace) |
| `azure_vnet_site.site.tenant` | [azure_vnet_site.site.tenant](data-sources--cloud_connect--properties--azure_vnet_site--site.md#schema-azure_vnet_site--site--tenant) |
| `azure_vnet_site.vnet_attachments` | [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list` | [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables.md#schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--route_table_id) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables.md#schema-azure_vnet_site--vnet_attachments--vnet_list--custom_routing--route_tables--static_routes) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route` | [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--all_route_tables.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables.md#schema-azure_vnet_site--vnet_attachments--vnet_list--default_route--selective_route_tables--route_table_id) |
| `azure_vnet_site.vnet_attachments.vnet_list.labels` | [azure_vnet_site.vnet_attachments.vnet_list.labels](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--labels.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.manual_routing` | [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--manual_routing.md#section) |
| `azure_vnet_site.vnet_attachments.vnet_list.subscription_id` | [azure_vnet_site.vnet_attachments.vnet_list.subscription_id](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md#schema-azure_vnet_site--vnet_attachments--vnet_list--subscription_id) |
| `azure_vnet_site.vnet_attachments.vnet_list.vnet_id` | [azure_vnet_site.vnet_attachments.vnet_list.vnet_id](data-sources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md#schema-azure_vnet_site--vnet_attachments--vnet_list--vnet_id) |
| `description` | [description](data-sources--cloud_connect--reference.md#schema-description) |
| `id` | [id](data-sources--cloud_connect--reference.md#schema-id) |
| `labels` | [labels](data-sources--cloud_connect--reference.md#schema-labels) |
| `name` | [name](data-sources--cloud_connect--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--cloud_connect--reference.md#schema-namespace) |
| `segment` | [segment](data-sources--cloud_connect--properties--segment.md#section) |
| `segment.name` | [segment.name](data-sources--cloud_connect--properties--segment.md#schema-segment--name) |
| `segment.namespace` | [segment.namespace](data-sources--cloud_connect--properties--segment.md#schema-segment--namespace) |
| `segment.tenant` | [segment.tenant](data-sources--cloud_connect--properties--segment.md#schema-segment--tenant) |

## Next pages

- [aws_provider](data-sources--cloud_connect--properties--aws_provider.md)
- [azure_vnet_site](data-sources--cloud_connect--properties--azure_vnet_site.md)
- [segment](data-sources--cloud_connect--properties--segment.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
