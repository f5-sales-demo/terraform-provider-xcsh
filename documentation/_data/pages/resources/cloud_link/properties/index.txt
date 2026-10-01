---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 20119, "body_sha256": "sha256:38204696a4fe79c82a0c3ce030912066cc96e87a7a188d34dfbb78d514d2ed3a", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws", "xcsh-docs:resources:cloud_link:properties:disabled", "xcsh-docs:resources:cloud_link:properties:enabled", "xcsh-docs:resources:cloud_link:properties:gcp", "xcsh-docs:resources:cloud_link:properties:timeouts"], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:reference", "parent_id": "xcsh-docs:resources:cloud_link:fundamentals", "path": "documentation/resources/cloud_link/properties/index.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
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

- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/): complete subsection reference.

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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/disabled/): complete subsection reference.

- [enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/enabled/): complete subsection reference.

- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/): complete subsection reference.

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

Name of the Cloud Link. Must be unique within the namespace.

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

Namespace where the Cloud Link is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-annotations) |
| `aws` | [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/#section) |
| `aws.aws_cred` | [aws.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/aws_cred/#section) |
| `aws.aws_cred.name` | [aws.aws_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/aws_cred/#schema-aws--aws_cred--name) |
| `aws.aws_cred.namespace` | [aws.aws_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/aws_cred/#schema-aws--aws_cred--namespace) |
| `aws.aws_cred.tenant` | [aws.aws_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/aws_cred/#schema-aws--aws_cred--tenant) |
| `aws.byoc` | [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/#section) |
| `aws.byoc.connections` | [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#section) |
| `aws.byoc.connections.auth_key` | [aws.byoc.connections.auth_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/#section) |
| `aws.byoc.connections.auth_key.blindfold_secret_info` | [aws.byoc.connections.auth_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/#section) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/#schema-aws--byoc--connections--auth_key--blindfold_secret_info--decryption_provider) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.location` | [aws.byoc.connections.auth_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/#schema-aws--byoc--connections--auth_key--blindfold_secret_info--location) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/#schema-aws--byoc--connections--auth_key--blindfold_secret_info--store_provider) |
| `aws.byoc.connections.auth_key.clear_secret_info` | [aws.byoc.connections.auth_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/#section) |
| `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` | [aws.byoc.connections.auth_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/#schema-aws--byoc--connections--auth_key--clear_secret_info--provider_ref) |
| `aws.byoc.connections.auth_key.clear_secret_info.url` | [aws.byoc.connections.auth_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/#schema-aws--byoc--connections--auth_key--clear_secret_info--url) |
| `aws.byoc.connections.bgp_asn` | [aws.byoc.connections.bgp_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--bgp_asn) |
| `aws.byoc.connections.connection_id` | [aws.byoc.connections.connection_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--connection_id) |
| `aws.byoc.connections.ipv4` | [aws.byoc.connections.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/ipv4/#section) |
| `aws.byoc.connections.ipv4.aws_router_peer_address` | [aws.byoc.connections.ipv4.aws_router_peer_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/ipv4/#schema-aws--byoc--connections--ipv4--aws_router_peer_address) |
| `aws.byoc.connections.ipv4.router_peer_address` | [aws.byoc.connections.ipv4.router_peer_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/ipv4/#schema-aws--byoc--connections--ipv4--router_peer_address) |
| `aws.byoc.connections.metadata` | [aws.byoc.connections.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/metadata/#section) |
| `aws.byoc.connections.metadata.description_spec` | [aws.byoc.connections.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/metadata/#schema-aws--byoc--connections--metadata--description_spec) |
| `aws.byoc.connections.metadata.name` | [aws.byoc.connections.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/metadata/#schema-aws--byoc--connections--metadata--name) |
| `aws.byoc.connections.region` | [aws.byoc.connections.region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--region) |
| `aws.byoc.connections.system_generated_name` | [aws.byoc.connections.system_generated_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/system_generated_name/#section) |
| `aws.byoc.connections.tags` | [aws.byoc.connections.tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--tags) |
| `aws.byoc.connections.user_assigned_name` | [aws.byoc.connections.user_assigned_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--user_assigned_name) |
| `aws.byoc.connections.virtual_interface_type` | [aws.byoc.connections.virtual_interface_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--virtual_interface_type) |
| `aws.byoc.connections.vlan` | [aws.byoc.connections.vlan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/#schema-aws--byoc--connections--vlan) |
| `aws.custom_asn` | [aws.custom_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/#schema-aws--custom_asn) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-disable) |
| `disabled` | [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/disabled/#section) |
| `enabled` | [enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/enabled/#section) |
| `enabled.cloudlink_network_name` | [enabled.cloudlink_network_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/enabled/#schema-enabled--cloudlink_network_name) |
| `gcp` | [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/#section) |
| `gcp.byoc` | [gcp.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/#section) |
| `gcp.byoc.connections` | [gcp.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/#section) |
| `gcp.byoc.connections.interconnect_attachment_name` | [gcp.byoc.connections.interconnect_attachment_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/#schema-gcp--byoc--connections--interconnect_attachment_name) |
| `gcp.byoc.connections.metadata` | [gcp.byoc.connections.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/metadata/#section) |
| `gcp.byoc.connections.metadata.description_spec` | [gcp.byoc.connections.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/metadata/#schema-gcp--byoc--connections--metadata--description_spec) |
| `gcp.byoc.connections.metadata.name` | [gcp.byoc.connections.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/metadata/#schema-gcp--byoc--connections--metadata--name) |
| `gcp.byoc.connections.project` | [gcp.byoc.connections.project](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/#schema-gcp--byoc--connections--project) |
| `gcp.byoc.connections.region` | [gcp.byoc.connections.region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/#schema-gcp--byoc--connections--region) |
| `gcp.byoc.connections.same_as_credential` | [gcp.byoc.connections.same_as_credential](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/same_as_credential/#section) |
| `gcp.gcp_cred` | [gcp.gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/gcp_cred/#section) |
| `gcp.gcp_cred.name` | [gcp.gcp_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/gcp_cred/#schema-gcp--gcp_cred--name) |
| `gcp.gcp_cred.namespace` | [gcp.gcp_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/gcp_cred/#schema-gcp--gcp_cred--namespace) |
| `gcp.gcp_cred.tenant` | [gcp.gcp_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/gcp_cred/#schema-gcp--gcp_cred--tenant) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/disabled/)
- [enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/enabled/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/timeouts/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
