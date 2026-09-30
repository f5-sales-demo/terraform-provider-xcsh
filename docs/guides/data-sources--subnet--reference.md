---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 9913, "body_sha256": "sha256:292a1c417a092235e29c69feabbda63c4344bf8cca3a6b71cc6cd738ad43a68f", "canonical_id": "xcsh-docs:data-sources:subnet:reference", "child_ids": ["xcsh-docs:data-sources:subnet:properties:connect_to_layer2", "xcsh-docs:data-sources:subnet:properties:connect_to_slo", "xcsh-docs:data-sources:subnet:properties:isolated_nw", "xcsh-docs:data-sources:subnet:properties:site_subnet_params"], "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:reference", "parent_id": "xcsh-docs:data-sources:subnet:fundamentals", "path": "docs/guides/data-sources--subnet--reference.md", "provider_name": "subnet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md)
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

- [connect_to_layer2](data-sources--subnet--properties--connect_to_layer2.md): complete subsection reference.

- [connect_to_slo](data-sources--subnet--properties--connect_to_slo.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Subnet.

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

- [isolated_nw](data-sources--subnet--properties--isolated_nw.md): complete subsection reference.

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

Name of the Subnet.

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

Namespace where the Subnet exists.

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

- [site_subnet_params](data-sources--subnet--properties--site_subnet_params.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--subnet--reference.md#schema-annotations) |
| `connect_to_layer2` | [connect_to_layer2](data-sources--subnet--properties--connect_to_layer2.md#section) |
| `connect_to_layer2.layer2_intf_ref` | [connect_to_layer2.layer2_intf_ref](data-sources--subnet--properties--connect_to_layer2--layer2_intf_ref.md#section) |
| `connect_to_layer2.layer2_intf_ref.name` | [connect_to_layer2.layer2_intf_ref.name](data-sources--subnet--properties--connect_to_layer2--layer2_intf_ref.md#schema-connect_to_layer2--layer2_intf_ref--name) |
| `connect_to_layer2.layer2_intf_ref.namespace` | [connect_to_layer2.layer2_intf_ref.namespace](data-sources--subnet--properties--connect_to_layer2--layer2_intf_ref.md#schema-connect_to_layer2--layer2_intf_ref--namespace) |
| `connect_to_layer2.layer2_intf_ref.tenant` | [connect_to_layer2.layer2_intf_ref.tenant](data-sources--subnet--properties--connect_to_layer2--layer2_intf_ref.md#schema-connect_to_layer2--layer2_intf_ref--tenant) |
| `connect_to_slo` | [connect_to_slo](data-sources--subnet--properties--connect_to_slo.md#section) |
| `description` | [description](data-sources--subnet--reference.md#schema-description) |
| `id` | [id](data-sources--subnet--reference.md#schema-id) |
| `isolated_nw` | [isolated_nw](data-sources--subnet--properties--isolated_nw.md#section) |
| `labels` | [labels](data-sources--subnet--reference.md#schema-labels) |
| `name` | [name](data-sources--subnet--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--subnet--reference.md#schema-namespace) |
| `site_subnet_params` | [site_subnet_params](data-sources--subnet--properties--site_subnet_params.md#section) |
| `site_subnet_params.dhcp` | [site_subnet_params.dhcp](data-sources--subnet--properties--site_subnet_params--dhcp.md#section) |
| `site_subnet_params.site` | [site_subnet_params.site](data-sources--subnet--properties--site_subnet_params--site.md#section) |
| `site_subnet_params.site.name` | [site_subnet_params.site.name](data-sources--subnet--properties--site_subnet_params--site.md#schema-site_subnet_params--site--name) |
| `site_subnet_params.site.namespace` | [site_subnet_params.site.namespace](data-sources--subnet--properties--site_subnet_params--site.md#schema-site_subnet_params--site--namespace) |
| `site_subnet_params.site.tenant` | [site_subnet_params.site.tenant](data-sources--subnet--properties--site_subnet_params--site.md#schema-site_subnet_params--site--tenant) |
| `site_subnet_params.static_ip` | [site_subnet_params.static_ip](data-sources--subnet--properties--site_subnet_params--static_ip.md#section) |
| `site_subnet_params.subnet_dhcp_server_params` | [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--properties--site_subnet_params--subnet_dhcp_server_params.md#section) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](data-sources--subnet--properties--site_subnet_params--subnet_dhcp_server_params--dhcp_networks.md#section) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix](data-sources--subnet--properties--site_subnet_params--subnet_dhcp_server_params--dhcp_networks.md#schema-site_subnet_params--subnet_dhcp_server_params--dhcp_networks--network_prefix) |

## Next pages

- [connect_to_layer2](data-sources--subnet--properties--connect_to_layer2.md)
- [connect_to_slo](data-sources--subnet--properties--connect_to_slo.md)
- [isolated_nw](data-sources--subnet--properties--isolated_nw.md)
- [site_subnet_params](data-sources--subnet--properties--site_subnet_params.md)
- [xcsh_subnet](../data-sources/subnet.md)
