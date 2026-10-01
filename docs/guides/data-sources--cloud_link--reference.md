---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 15180, "body_sha256": "sha256:59c141bfec44de9baf7a66aaef65d6baeae34eec6010fd09aeed9550656f7121", "canonical_id": "xcsh-docs:data-sources:cloud_link:reference", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:aws", "xcsh-docs:data-sources:cloud_link:properties:disabled", "xcsh-docs:data-sources:cloud_link:properties:enabled", "xcsh-docs:data-sources:cloud_link:properties:gcp"], "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:reference", "parent_id": "xcsh-docs:data-sources:cloud_link:fundamentals", "path": "docs/guides/data-sources--cloud_link--reference.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md)
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

- [aws](data-sources--cloud_link--properties--aws.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the CloudLink.

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

- [disabled](data-sources--cloud_link--properties--disabled.md): complete subsection reference.

- [enabled](data-sources--cloud_link--properties--enabled.md): complete subsection reference.

- [gcp](data-sources--cloud_link--properties--gcp.md): complete subsection reference.

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

Name of the CloudLink.

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

Namespace where the CloudLink exists.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_link--reference.md#schema-annotations) |
| `aws` | [aws](data-sources--cloud_link--properties--aws.md#section) |
| `aws.aws_cred` | [aws.aws_cred](data-sources--cloud_link--properties--aws--aws_cred.md#section) |
| `aws.aws_cred.name` | [aws.aws_cred.name](data-sources--cloud_link--properties--aws--aws_cred.md#schema-aws--aws_cred--name) |
| `aws.aws_cred.namespace` | [aws.aws_cred.namespace](data-sources--cloud_link--properties--aws--aws_cred.md#schema-aws--aws_cred--namespace) |
| `aws.aws_cred.tenant` | [aws.aws_cred.tenant](data-sources--cloud_link--properties--aws--aws_cred.md#schema-aws--aws_cred--tenant) |
| `aws.byoc` | [aws.byoc](data-sources--cloud_link--properties--aws--byoc.md#section) |
| `aws.byoc.connections` | [aws.byoc.connections](data-sources--cloud_link--properties--aws--byoc--connections.md#section) |
| `aws.byoc.connections.auth_key` | [aws.byoc.connections.auth_key](data-sources--cloud_link--properties--aws--byoc--connections--auth_key.md#section) |
| `aws.byoc.connections.auth_key.blindfold_secret_info` | [aws.byoc.connections.auth_key.blindfold_secret_info](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--blindfold_secret_info.md#section) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--blindfold_secret_info.md#schema-aws--byoc--connections--auth_key--blindfold_secret_info--decryption_provider) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.location` | [aws.byoc.connections.auth_key.blindfold_secret_info.location](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--blindfold_secret_info.md#schema-aws--byoc--connections--auth_key--blindfold_secret_info--location) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.store_provider](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--blindfold_secret_info.md#schema-aws--byoc--connections--auth_key--blindfold_secret_info--store_provider) |
| `aws.byoc.connections.auth_key.clear_secret_info` | [aws.byoc.connections.auth_key.clear_secret_info](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--clear_secret_info.md#section) |
| `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` | [aws.byoc.connections.auth_key.clear_secret_info.provider_ref](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--clear_secret_info.md#schema-aws--byoc--connections--auth_key--clear_secret_info--provider_ref) |
| `aws.byoc.connections.auth_key.clear_secret_info.url` | [aws.byoc.connections.auth_key.clear_secret_info.url](data-sources--cloud_link--properties--aws--byoc--connections--auth_key--clear_secret_info.md#schema-aws--byoc--connections--auth_key--clear_secret_info--url) |
| `aws.byoc.connections.bgp_asn` | [aws.byoc.connections.bgp_asn](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--bgp_asn) |
| `aws.byoc.connections.connection_id` | [aws.byoc.connections.connection_id](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--connection_id) |
| `aws.byoc.connections.ipv4` | [aws.byoc.connections.ipv4](data-sources--cloud_link--properties--aws--byoc--connections--ipv4.md#section) |
| `aws.byoc.connections.ipv4.aws_router_peer_address` | [aws.byoc.connections.ipv4.aws_router_peer_address](data-sources--cloud_link--properties--aws--byoc--connections--ipv4.md#schema-aws--byoc--connections--ipv4--aws_router_peer_address) |
| `aws.byoc.connections.ipv4.router_peer_address` | [aws.byoc.connections.ipv4.router_peer_address](data-sources--cloud_link--properties--aws--byoc--connections--ipv4.md#schema-aws--byoc--connections--ipv4--router_peer_address) |
| `aws.byoc.connections.metadata` | [aws.byoc.connections.metadata](data-sources--cloud_link--properties--aws--byoc--connections--metadata.md#section) |
| `aws.byoc.connections.metadata.description_spec` | [aws.byoc.connections.metadata.description_spec](data-sources--cloud_link--properties--aws--byoc--connections--metadata.md#schema-aws--byoc--connections--metadata--description_spec) |
| `aws.byoc.connections.metadata.name` | [aws.byoc.connections.metadata.name](data-sources--cloud_link--properties--aws--byoc--connections--metadata.md#schema-aws--byoc--connections--metadata--name) |
| `aws.byoc.connections.region` | [aws.byoc.connections.region](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--region) |
| `aws.byoc.connections.system_generated_name` | [aws.byoc.connections.system_generated_name](data-sources--cloud_link--properties--aws--byoc--connections--system_generated_name.md#section) |
| `aws.byoc.connections.tags` | [aws.byoc.connections.tags](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--tags) |
| `aws.byoc.connections.user_assigned_name` | [aws.byoc.connections.user_assigned_name](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--user_assigned_name) |
| `aws.byoc.connections.virtual_interface_type` | [aws.byoc.connections.virtual_interface_type](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--virtual_interface_type) |
| `aws.byoc.connections.vlan` | [aws.byoc.connections.vlan](data-sources--cloud_link--properties--aws--byoc--connections.md#schema-aws--byoc--connections--vlan) |
| `aws.custom_asn` | [aws.custom_asn](data-sources--cloud_link--properties--aws.md#schema-aws--custom_asn) |
| `description` | [description](data-sources--cloud_link--reference.md#schema-description) |
| `disabled` | [disabled](data-sources--cloud_link--properties--disabled.md#section) |
| `enabled` | [enabled](data-sources--cloud_link--properties--enabled.md#section) |
| `enabled.cloudlink_network_name` | [enabled.cloudlink_network_name](data-sources--cloud_link--properties--enabled.md#schema-enabled--cloudlink_network_name) |
| `gcp` | [gcp](data-sources--cloud_link--properties--gcp.md#section) |
| `gcp.byoc` | [gcp.byoc](data-sources--cloud_link--properties--gcp--byoc.md#section) |
| `gcp.byoc.connections` | [gcp.byoc.connections](data-sources--cloud_link--properties--gcp--byoc--connections.md#section) |
| `gcp.byoc.connections.interconnect_attachment_name` | [gcp.byoc.connections.interconnect_attachment_name](data-sources--cloud_link--properties--gcp--byoc--connections.md#schema-gcp--byoc--connections--interconnect_attachment_name) |
| `gcp.byoc.connections.metadata` | [gcp.byoc.connections.metadata](data-sources--cloud_link--properties--gcp--byoc--connections--metadata.md#section) |
| `gcp.byoc.connections.metadata.description_spec` | [gcp.byoc.connections.metadata.description_spec](data-sources--cloud_link--properties--gcp--byoc--connections--metadata.md#schema-gcp--byoc--connections--metadata--description_spec) |
| `gcp.byoc.connections.metadata.name` | [gcp.byoc.connections.metadata.name](data-sources--cloud_link--properties--gcp--byoc--connections--metadata.md#schema-gcp--byoc--connections--metadata--name) |
| `gcp.byoc.connections.project` | [gcp.byoc.connections.project](data-sources--cloud_link--properties--gcp--byoc--connections.md#schema-gcp--byoc--connections--project) |
| `gcp.byoc.connections.region` | [gcp.byoc.connections.region](data-sources--cloud_link--properties--gcp--byoc--connections.md#schema-gcp--byoc--connections--region) |
| `gcp.byoc.connections.same_as_credential` | [gcp.byoc.connections.same_as_credential](data-sources--cloud_link--properties--gcp--byoc--connections--same_as_credential.md#section) |
| `gcp.gcp_cred` | [gcp.gcp_cred](data-sources--cloud_link--properties--gcp--gcp_cred.md#section) |
| `gcp.gcp_cred.name` | [gcp.gcp_cred.name](data-sources--cloud_link--properties--gcp--gcp_cred.md#schema-gcp--gcp_cred--name) |
| `gcp.gcp_cred.namespace` | [gcp.gcp_cred.namespace](data-sources--cloud_link--properties--gcp--gcp_cred.md#schema-gcp--gcp_cred--namespace) |
| `gcp.gcp_cred.tenant` | [gcp.gcp_cred.tenant](data-sources--cloud_link--properties--gcp--gcp_cred.md#schema-gcp--gcp_cred--tenant) |
| `id` | [id](data-sources--cloud_link--reference.md#schema-id) |
| `labels` | [labels](data-sources--cloud_link--reference.md#schema-labels) |
| `name` | [name](data-sources--cloud_link--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--cloud_link--reference.md#schema-namespace) |

## Next pages

- [aws](data-sources--cloud_link--properties--aws.md)
- [disabled](data-sources--cloud_link--properties--disabled.md)
- [enabled](data-sources--cloud_link--properties--enabled.md)
- [gcp](data-sources--cloud_link--properties--gcp.md)
- [xcsh_cloud_link](../data-sources/cloud_link.md)
