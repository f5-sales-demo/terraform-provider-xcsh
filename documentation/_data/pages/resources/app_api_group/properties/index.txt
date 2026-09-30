---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 14151, "body_sha256": "sha256:734b11c8f3a530ea75afaf91263b7b0ec431f96d805f9badd0acd96a3314b40f", "child_ids": ["xcsh-docs:resources:app_api_group:properties:bigip_virtual_server", "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer", "xcsh-docs:resources:app_api_group:properties:elements", "xcsh-docs:resources:app_api_group:properties:http_loadbalancer", "xcsh-docs:resources:app_api_group:properties:timeouts"], "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:reference", "parent_id": "xcsh-docs:resources:app_api_group:fundamentals", "path": "documentation/resources/app_api_group/properties/index.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
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

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/): complete subsection reference.

- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/): complete subsection reference.

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

- [elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/elements/): complete subsection reference.

- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/): complete subsection reference.

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

Name of the App API Group. Must be unique within the namespace.

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

Namespace where the App API Group is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-annotations) |
| `bigip_virtual_server` | [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/#section) |
| `bigip_virtual_server.bigip_virtual_server` | [bigip_virtual_server.bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/#section) |
| `bigip_virtual_server.bigip_virtual_server.name` | [bigip_virtual_server.bigip_virtual_server.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/#schema-bigip_virtual_server--bigip_virtual_server--name) |
| `bigip_virtual_server.bigip_virtual_server.namespace` | [bigip_virtual_server.bigip_virtual_server.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/#schema-bigip_virtual_server--bigip_virtual_server--namespace) |
| `bigip_virtual_server.bigip_virtual_server.tenant` | [bigip_virtual_server.bigip_virtual_server.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/#schema-bigip_virtual_server--bigip_virtual_server--tenant) |
| `cdn_loadbalancer` | [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/#section) |
| `cdn_loadbalancer.cdn_loadbalancer` | [cdn_loadbalancer.cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/#section) |
| `cdn_loadbalancer.cdn_loadbalancer.name` | [cdn_loadbalancer.cdn_loadbalancer.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/#schema-cdn_loadbalancer--cdn_loadbalancer--name) |
| `cdn_loadbalancer.cdn_loadbalancer.namespace` | [cdn_loadbalancer.cdn_loadbalancer.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/#schema-cdn_loadbalancer--cdn_loadbalancer--namespace) |
| `cdn_loadbalancer.cdn_loadbalancer.tenant` | [cdn_loadbalancer.cdn_loadbalancer.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/#schema-cdn_loadbalancer--cdn_loadbalancer--tenant) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-disable) |
| `elements` | [elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/elements/#section) |
| `elements.methods` | [elements.methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/elements/#schema-elements--methods) |
| `elements.path_regex` | [elements.path_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/elements/#schema-elements--path_regex) |
| `http_loadbalancer` | [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/#section) |
| `http_loadbalancer.http_loadbalancer` | [http_loadbalancer.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/http_loadbalancer/#section) |
| `http_loadbalancer.http_loadbalancer.name` | [http_loadbalancer.http_loadbalancer.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/http_loadbalancer/#schema-http_loadbalancer--http_loadbalancer--name) |
| `http_loadbalancer.http_loadbalancer.namespace` | [http_loadbalancer.http_loadbalancer.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/http_loadbalancer/#schema-http_loadbalancer--http_loadbalancer--namespace) |
| `http_loadbalancer.http_loadbalancer.tenant` | [http_loadbalancer.http_loadbalancer.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/http_loadbalancer/#schema-http_loadbalancer--http_loadbalancer--tenant) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/bigip_virtual_server/)
- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/)
- [elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/elements/)
- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/http_loadbalancer/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/timeouts/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
