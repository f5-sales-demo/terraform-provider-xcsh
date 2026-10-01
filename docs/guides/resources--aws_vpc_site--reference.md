---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 119573, "body_sha256": "sha256:e42487da31e83f5fd9acfdeea3e43655d221a2cacb2f0a9f7f5e745a076b1c6c", "canonical_id": "xcsh-docs:resources:aws_vpc_site:reference", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:admin_password", "xcsh-docs:resources:aws_vpc_site:properties:aws_cred", "xcsh-docs:resources:aws_vpc_site:properties:block_all_services", "xcsh-docs:resources:aws_vpc_site:properties:blocked_services", "xcsh-docs:resources:aws_vpc_site:properties:coordinates", "xcsh-docs:resources:aws_vpc_site:properties:custom_dns", "xcsh-docs:resources:aws_vpc_site:properties:custom_security_group", "xcsh-docs:resources:aws_vpc_site:properties:default_blocked_services", "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_disabled", "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled", "xcsh-docs:resources:aws_vpc_site:properties:disable_encryption", "xcsh-docs:resources:aws_vpc_site:properties:disable_internet_vip", "xcsh-docs:resources:aws_vpc_site:properties:egress_gateway_default", "xcsh-docs:resources:aws_vpc_site:properties:egress_nat_gw", "xcsh-docs:resources:aws_vpc_site:properties:egress_virtual_private_gateway", "xcsh-docs:resources:aws_vpc_site:properties:enable_encryption", "xcsh-docs:resources:aws_vpc_site:properties:enable_internet_vip", "xcsh-docs:resources:aws_vpc_site:properties:f5_orchestrated_routing", "xcsh-docs:resources:aws_vpc_site:properties:f5xc_security_group", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw", "xcsh-docs:resources:aws_vpc_site:properties:kubernetes_upgrade_drain", "xcsh-docs:resources:aws_vpc_site:properties:log_receiver", "xcsh-docs:resources:aws_vpc_site:properties:logs_streaming_disabled", "xcsh-docs:resources:aws_vpc_site:properties:manual_routing", "xcsh-docs:resources:aws_vpc_site:properties:no_worker_nodes", "xcsh-docs:resources:aws_vpc_site:properties:offline_survivability_mode", "xcsh-docs:resources:aws_vpc_site:properties:os", "xcsh-docs:resources:aws_vpc_site:properties:private_connectivity", "xcsh-docs:resources:aws_vpc_site:properties:sw", "xcsh-docs:resources:aws_vpc_site:properties:timeouts", "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster", "xcsh-docs:resources:aws_vpc_site:properties:vpc", "xcsh-docs:resources:aws_vpc_site:properties:waf_signatures"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:reference", "parent_id": "xcsh-docs:resources:aws_vpc_site:fundamentals", "path": "docs/guides/resources--aws_vpc_site--reference.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [admin_password](resources--aws_vpc_site--properties--admin_password.md): complete subsection reference.

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

- [aws_cred](resources--aws_vpc_site--properties--aws_cred.md): complete subsection reference.

<a id="schema-aws_region"></a>

### aws_region property

Type: `"string"`. Required.

AWS Region. Name for AWS Region.

Upstream description:

Name for AWS Region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [block_all_services](resources--aws_vpc_site--properties--block_all_services.md): complete subsection reference.

- [blocked_services](resources--aws_vpc_site--properties--blocked_services.md): complete subsection reference.

- [coordinates](resources--aws_vpc_site--properties--coordinates.md): complete subsection reference.

- [custom_dns](resources--aws_vpc_site--properties--custom_dns.md): complete subsection reference.

- [custom_security_group](resources--aws_vpc_site--properties--custom_security_group.md): complete subsection reference.

- [default_blocked_services](resources--aws_vpc_site--properties--default_blocked_services.md): complete subsection reference.

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

- [direct_connect_disabled](resources--aws_vpc_site--properties--direct_connect_disabled.md): complete subsection reference.

- [direct_connect_enabled](resources--aws_vpc_site--properties--direct_connect_enabled.md): complete subsection reference.

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

- [disable_encryption](resources--aws_vpc_site--properties--disable_encryption.md): complete subsection reference.

- [disable_internet_vip](resources--aws_vpc_site--properties--disable_internet_vip.md): complete subsection reference.

<a id="schema-disk_size"></a>

### disk_size property

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "2048"
  }
}
```

- [egress_gateway_default](resources--aws_vpc_site--properties--egress_gateway_default.md): complete subsection reference.

- [egress_nat_gw](resources--aws_vpc_site--properties--egress_nat_gw.md): complete subsection reference.

- [egress_virtual_private_gateway](resources--aws_vpc_site--properties--egress_virtual_private_gateway.md): complete subsection reference.

- [enable_encryption](resources--aws_vpc_site--properties--enable_encryption.md): complete subsection reference.

- [enable_internet_vip](resources--aws_vpc_site--properties--enable_internet_vip.md): complete subsection reference.

- [f5_orchestrated_routing](resources--aws_vpc_site--properties--f5_orchestrated_routing.md): complete subsection reference.

- [f5xc_security_group](resources--aws_vpc_site--properties--f5xc_security_group.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md): complete subsection reference.

- [ingress_gw](resources--aws_vpc_site--properties--ingress_gw.md): complete subsection reference.

<a id="schema-instance_type"></a>

### instance_type property

Type: `"string"`. Required.

Select Instance size based on performance needed.

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

- [kubernetes_upgrade_drain](resources--aws_vpc_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](resources--aws_vpc_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](resources--aws_vpc_site--properties--logs_streaming_disabled.md): complete subsection reference.

- [manual_routing](resources--aws_vpc_site--properties--manual_routing.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AWS VPC Site. Must be unique within the namespace.

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

Namespace where the AWS VPC Site is created.

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

- [no_worker_nodes](resources--aws_vpc_site--properties--no_worker_nodes.md): complete subsection reference.

<a id="schema-nodes_per_az"></a>

### nodes_per_az property

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](resources--aws_vpc_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](resources--aws_vpc_site--properties--os.md): complete subsection reference.

- [private_connectivity](resources--aws_vpc_site--properties--private_connectivity.md): complete subsection reference.

<a id="schema-ssh_key"></a>

### ssh_key property

Type: `"string"`. Required.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](resources--aws_vpc_site--properties--sw.md): complete subsection reference.

<a id="schema-tags"></a>

### tags property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [timeouts](resources--aws_vpc_site--properties--timeouts.md): complete subsection reference.

<a id="schema-total_nodes"></a>

### total_nodes property

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md): complete subsection reference.

- [vpc](resources--aws_vpc_site--properties--vpc.md): complete subsection reference.

- [waf_signatures](resources--aws_vpc_site--properties--waf_signatures.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--aws_vpc_site--reference.md#schema-address) |
| `admin_password` | [admin_password](resources--aws_vpc_site--properties--admin_password.md#section) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](resources--aws_vpc_site--properties--admin_password--blindfold_secret_info.md#section) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](resources--aws_vpc_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--decryption_provider) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](resources--aws_vpc_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--location) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](resources--aws_vpc_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--store_provider) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](resources--aws_vpc_site--properties--admin_password--clear_secret_info.md#section) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](resources--aws_vpc_site--properties--admin_password--clear_secret_info.md#schema-admin_password--clear_secret_info--provider_ref) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](resources--aws_vpc_site--properties--admin_password--clear_secret_info.md#schema-admin_password--clear_secret_info--url) |
| `annotations` | [annotations](resources--aws_vpc_site--reference.md#schema-annotations) |
| `aws_cred` | [aws_cred](resources--aws_vpc_site--properties--aws_cred.md#section) |
| `aws_cred.name` | [aws_cred.name](resources--aws_vpc_site--properties--aws_cred.md#schema-aws_cred--name) |
| `aws_cred.namespace` | [aws_cred.namespace](resources--aws_vpc_site--properties--aws_cred.md#schema-aws_cred--namespace) |
| `aws_cred.tenant` | [aws_cred.tenant](resources--aws_vpc_site--properties--aws_cred.md#schema-aws_cred--tenant) |
| `aws_region` | [aws_region](resources--aws_vpc_site--reference.md#schema-aws_region) |
| `block_all_services` | [block_all_services](resources--aws_vpc_site--properties--block_all_services.md#section) |
| `blocked_services` | [blocked_services](resources--aws_vpc_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--aws_vpc_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--aws_vpc_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--aws_vpc_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--aws_vpc_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--aws_vpc_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `coordinates` | [coordinates](resources--aws_vpc_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](resources--aws_vpc_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](resources--aws_vpc_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](resources--aws_vpc_site--properties--custom_dns.md#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--aws_vpc_site--properties--custom_dns.md#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--aws_vpc_site--properties--custom_dns.md#schema-custom_dns--outside_nameserver) |
| `custom_security_group` | [custom_security_group](resources--aws_vpc_site--properties--custom_security_group.md#section) |
| `custom_security_group.inside_security_group_id` | [custom_security_group.inside_security_group_id](resources--aws_vpc_site--properties--custom_security_group.md#schema-custom_security_group--inside_security_group_id) |
| `custom_security_group.outside_security_group_id` | [custom_security_group.outside_security_group_id](resources--aws_vpc_site--properties--custom_security_group.md#schema-custom_security_group--outside_security_group_id) |
| `default_blocked_services` | [default_blocked_services](resources--aws_vpc_site--properties--default_blocked_services.md#section) |
| `description` | [description](resources--aws_vpc_site--reference.md#schema-description) |
| `direct_connect_disabled` | [direct_connect_disabled](resources--aws_vpc_site--properties--direct_connect_disabled.md#section) |
| `direct_connect_enabled` | [direct_connect_enabled](resources--aws_vpc_site--properties--direct_connect_enabled.md#section) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](resources--aws_vpc_site--properties--direct_connect_enabled--auto_asn.md#section) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](resources--aws_vpc_site--properties--direct_connect_enabled.md#schema-direct_connect_enabled--custom_asn) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs.md#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md#schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#schema-direct_connect_enabled--hosted_vifs--vif_list--other_region) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list--same_as_site_region.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](resources--aws_vpc_site--properties--direct_connect_enabled--standard_vifs.md#section) |
| `disable` | [disable](resources--aws_vpc_site--reference.md#schema-disable) |
| `disable_encryption` | [disable_encryption](resources--aws_vpc_site--properties--disable_encryption.md#section) |
| `disable_internet_vip` | [disable_internet_vip](resources--aws_vpc_site--properties--disable_internet_vip.md#section) |
| `disk_size` | [disk_size](resources--aws_vpc_site--reference.md#schema-disk_size) |
| `egress_gateway_default` | [egress_gateway_default](resources--aws_vpc_site--properties--egress_gateway_default.md#section) |
| `egress_nat_gw` | [egress_nat_gw](resources--aws_vpc_site--properties--egress_nat_gw.md#section) |
| `egress_nat_gw.nat_gw_id` | [egress_nat_gw.nat_gw_id](resources--aws_vpc_site--properties--egress_nat_gw.md#schema-egress_nat_gw--nat_gw_id) |
| `egress_virtual_private_gateway` | [egress_virtual_private_gateway](resources--aws_vpc_site--properties--egress_virtual_private_gateway.md#section) |
| `egress_virtual_private_gateway.vgw_id` | [egress_virtual_private_gateway.vgw_id](resources--aws_vpc_site--properties--egress_virtual_private_gateway.md#schema-egress_virtual_private_gateway--vgw_id) |
| `enable_encryption` | [enable_encryption](resources--aws_vpc_site--properties--enable_encryption.md#section) |
| `enable_encryption.kms_key_id` | [enable_encryption.kms_key_id](resources--aws_vpc_site--properties--enable_encryption.md#schema-enable_encryption--kms_key_id) |
| `enable_internet_vip` | [enable_internet_vip](resources--aws_vpc_site--properties--enable_internet_vip.md#section) |
| `f5_orchestrated_routing` | [f5_orchestrated_routing](resources--aws_vpc_site--properties--f5_orchestrated_routing.md#section) |
| `f5xc_security_group` | [f5xc_security_group](resources--aws_vpc_site--properties--f5xc_security_group.md#section) |
| `id` | [id](resources--aws_vpc_site--reference.md#schema-id) |
| `ingress_egress_gw` | [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies.md#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies.md#section) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#section) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--name) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw.allowed_vip_port` | [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port.md#section) |
| `ingress_egress_gw.allowed_vip_port.custom_ports` | [ingress_egress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--custom_ports.md#section) |
| `ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--custom_ports.md#schema-ingress_egress_gw--allowed_vip_port--custom_ports--port_ranges) |
| `ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--disable_allowed_vip_port.md#section) |
| `ingress_egress_gw.allowed_vip_port.use_http_https_port` | [ingress_egress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--use_http_https_port.md#section) |
| `ingress_egress_gw.allowed_vip_port.use_http_port` | [ingress_egress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--use_http_port.md#section) |
| `ingress_egress_gw.allowed_vip_port.use_https_port` | [ingress_egress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port--use_https_port.md#section) |
| `ingress_egress_gw.allowed_vip_port_sli` | [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli.md#section) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--custom_ports.md#section) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--custom_ports.md#schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges) |
| `ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--disable_allowed_vip_port.md#section) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_http_https_port.md#section) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_http_port.md#section) |
| `ingress_egress_gw.allowed_vip_port_sli.use_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_https_port.md#section) |
| `ingress_egress_gw.aws_certified_hw` | [ingress_egress_gw.aws_certified_hw](resources--aws_vpc_site--properties--ingress_egress_gw.md#schema-ingress_egress_gw--aws_certified_hw) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes.md#section) |
| `ingress_egress_gw.az_nodes.aws_az_name` | [ingress_egress_gw.az_nodes.aws_az_name](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes.md#schema-ingress_egress_gw--az_nodes--aws_az_name) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet.md#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet.md#schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet_param.md#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet_param.md#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--outside_subnet.md#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--outside_subnet.md#schema-ingress_egress_gw--az_nodes--outside_subnet--existing_subnet_id) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet_param.md#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet_param.md#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.az_nodes.reserved_inside_subnet` | [ingress_egress_gw.az_nodes.reserved_inside_subnet](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--reserved_inside_subnet.md#section) |
| `ingress_egress_gw.az_nodes.workload_subnet` | [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet.md#section) |
| `ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet.md#schema-ingress_egress_gw--az_nodes--workload_subnet--existing_subnet_id) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet--subnet_param.md#section) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--workload_subnet--subnet_param.md#schema-ingress_egress_gw--az_nodes--workload_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#section) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#section) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](resources--aws_vpc_site--properties--ingress_egress_gw--forward_proxy_allow_all.md#section) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](resources--aws_vpc_site--properties--ingress_egress_gw--no_dc_cluster_group.md#section) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](resources--aws_vpc_site--properties--ingress_egress_gw--no_forward_proxy.md#section) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](resources--aws_vpc_site--properties--ingress_egress_gw--no_global_network.md#section) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](resources--aws_vpc_site--properties--ingress_egress_gw--no_inside_static_routes.md#section) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](resources--aws_vpc_site--properties--ingress_egress_gw--no_network_policy.md#section) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](resources--aws_vpc_site--properties--ingress_egress_gw--no_outside_static_routes.md#section) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](resources--aws_vpc_site--properties--ingress_egress_gw--sm_connection_public_ip.md#section) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](resources--aws_vpc_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md#section) |
| `ingress_gw` | [ingress_gw](resources--aws_vpc_site--properties--ingress_gw.md#section) |
| `ingress_gw.allowed_vip_port` | [ingress_gw.allowed_vip_port](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port.md#section) |
| `ingress_gw.allowed_vip_port.custom_ports` | [ingress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port--custom_ports.md#section) |
| `ingress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_gw.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port--custom_ports.md#schema-ingress_gw--allowed_vip_port--custom_ports--port_ranges) |
| `ingress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port--disable_allowed_vip_port.md#section) |
| `ingress_gw.allowed_vip_port.use_http_https_port` | [ingress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port--use_http_https_port.md#section) |
| `ingress_gw.allowed_vip_port.use_http_port` | [ingress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port--use_http_port.md#section) |
| `ingress_gw.allowed_vip_port.use_https_port` | [ingress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--properties--ingress_gw--allowed_vip_port--use_https_port.md#section) |
| `ingress_gw.aws_certified_hw` | [ingress_gw.aws_certified_hw](resources--aws_vpc_site--properties--ingress_gw.md#schema-ingress_gw--aws_certified_hw) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](resources--aws_vpc_site--properties--ingress_gw--az_nodes.md#section) |
| `ingress_gw.az_nodes.aws_az_name` | [ingress_gw.az_nodes.aws_az_name](resources--aws_vpc_site--properties--ingress_gw--az_nodes.md#schema-ingress_gw--az_nodes--aws_az_name) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--properties--ingress_gw--az_nodes--local_subnet.md#section) |
| `ingress_gw.az_nodes.local_subnet.existing_subnet_id` | [ingress_gw.az_nodes.local_subnet.existing_subnet_id](resources--aws_vpc_site--properties--ingress_gw--az_nodes--local_subnet.md#schema-ingress_gw--az_nodes--local_subnet--existing_subnet_id) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--properties--ingress_gw--az_nodes--local_subnet--subnet_param.md#section) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](resources--aws_vpc_site--properties--ingress_gw--az_nodes--local_subnet--subnet_param.md#schema-ingress_gw--az_nodes--local_subnet--subnet_param--ipv4) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `instance_type` | [instance_type](resources--aws_vpc_site--reference.md#schema-instance_type) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--aws_vpc_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](resources--aws_vpc_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](resources--aws_vpc_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](resources--aws_vpc_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--aws_vpc_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--aws_vpc_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--aws_vpc_site--properties--logs_streaming_disabled.md#section) |
| `manual_routing` | [manual_routing](resources--aws_vpc_site--properties--manual_routing.md#section) |
| `name` | [name](resources--aws_vpc_site--reference.md#schema-name) |
| `namespace` | [namespace](resources--aws_vpc_site--reference.md#schema-namespace) |
| `no_worker_nodes` | [no_worker_nodes](resources--aws_vpc_site--properties--no_worker_nodes.md#section) |
| `nodes_per_az` | [nodes_per_az](resources--aws_vpc_site--reference.md#schema-nodes_per_az) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--aws_vpc_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_vpc_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--aws_vpc_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](resources--aws_vpc_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](resources--aws_vpc_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](resources--aws_vpc_site--properties--os.md#schema-os--operating_system_version) |
| `private_connectivity` | [private_connectivity](resources--aws_vpc_site--properties--private_connectivity.md#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--aws_vpc_site--properties--private_connectivity--cloud_link.md#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--aws_vpc_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--aws_vpc_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--namespace) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--aws_vpc_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--tenant) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--aws_vpc_site--properties--private_connectivity--inside.md#section) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--aws_vpc_site--properties--private_connectivity--outside.md#section) |
| `ssh_key` | [ssh_key](resources--aws_vpc_site--reference.md#schema-ssh_key) |
| `sw` | [sw](resources--aws_vpc_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--aws_vpc_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--aws_vpc_site--properties--sw.md#schema-sw--volterra_software_version) |
| `tags` | [tags](resources--aws_vpc_site--reference.md#schema-tags) |
| `timeouts` | [timeouts](resources--aws_vpc_site--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--aws_vpc_site--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--aws_vpc_site--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--aws_vpc_site--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--aws_vpc_site--properties--timeouts.md#schema-timeouts--update) |
| `total_nodes` | [total_nodes](resources--aws_vpc_site--reference.md#schema-total_nodes) |
| `voltstack_cluster` | [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_network_policies.md#section) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](resources--aws_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#section) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](resources--aws_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--name) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](resources--aws_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--namespace) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](resources--aws_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--tenant) |
| `voltstack_cluster.allowed_vip_port` | [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port.md#section) |
| `voltstack_cluster.allowed_vip_port.custom_ports` | [voltstack_cluster.allowed_vip_port.custom_ports](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--custom_ports.md#section) |
| `voltstack_cluster.allowed_vip_port.custom_ports.port_ranges` | [voltstack_cluster.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--custom_ports.md#schema-voltstack_cluster--allowed_vip_port--custom_ports--port_ranges) |
| `voltstack_cluster.allowed_vip_port.disable_allowed_vip_port` | [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--disable_allowed_vip_port.md#section) |
| `voltstack_cluster.allowed_vip_port.use_http_https_port` | [voltstack_cluster.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_http_https_port.md#section) |
| `voltstack_cluster.allowed_vip_port.use_http_port` | [voltstack_cluster.allowed_vip_port.use_http_port](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_http_port.md#section) |
| `voltstack_cluster.allowed_vip_port.use_https_port` | [voltstack_cluster.allowed_vip_port.use_https_port](resources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port--use_https_port.md#section) |
| `voltstack_cluster.aws_certified_hw` | [voltstack_cluster.aws_certified_hw](resources--aws_vpc_site--properties--voltstack_cluster.md#schema-voltstack_cluster--aws_certified_hw) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes.md#section) |
| `voltstack_cluster.az_nodes.aws_az_name` | [voltstack_cluster.az_nodes.aws_az_name](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes.md#schema-voltstack_cluster--az_nodes--aws_az_name) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet.md#section) |
| `voltstack_cluster.az_nodes.local_subnet.existing_subnet_id` | [voltstack_cluster.az_nodes.local_subnet.existing_subnet_id](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet.md#schema-voltstack_cluster--az_nodes--local_subnet--existing_subnet_id) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md#schema-voltstack_cluster--az_nodes--local_subnet--subnet_param--ipv4) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](resources--aws_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#section) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](resources--aws_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--name) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](resources--aws_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--namespace) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](resources--aws_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--tenant) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](resources--aws_vpc_site--properties--voltstack_cluster--default_storage.md#section) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](resources--aws_vpc_site--properties--voltstack_cluster--forward_proxy_allow_all.md#section) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](resources--aws_vpc_site--properties--voltstack_cluster--k8s_cluster.md#section) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](resources--aws_vpc_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--name) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](resources--aws_vpc_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--namespace) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](resources--aws_vpc_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--tenant) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](resources--aws_vpc_site--properties--voltstack_cluster--no_dc_cluster_group.md#section) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](resources--aws_vpc_site--properties--voltstack_cluster--no_forward_proxy.md#section) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](resources--aws_vpc_site--properties--voltstack_cluster--no_global_network.md#section) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](resources--aws_vpc_site--properties--voltstack_cluster--no_k8s_cluster.md#section) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](resources--aws_vpc_site--properties--voltstack_cluster--no_network_policy.md#section) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](resources--aws_vpc_site--properties--voltstack_cluster--no_outside_static_routes.md#section) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md#schema-voltstack_cluster--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](resources--aws_vpc_site--properties--voltstack_cluster--sm_connection_public_ip.md#section) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](resources--aws_vpc_site--properties--voltstack_cluster--sm_connection_pvt_ip.md#section) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](resources--aws_vpc_site--properties--voltstack_cluster--storage_class_list.md#section) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](resources--aws_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#section) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](resources--aws_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](resources--aws_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name) |
| `vpc` | [vpc](resources--aws_vpc_site--properties--vpc.md#section) |
| `vpc.new_vpc` | [vpc.new_vpc](resources--aws_vpc_site--properties--vpc--new_vpc.md#section) |
| `vpc.new_vpc.autogenerate` | [vpc.new_vpc.autogenerate](resources--aws_vpc_site--properties--vpc--new_vpc--autogenerate.md#section) |
| `vpc.new_vpc.name_tag` | [vpc.new_vpc.name_tag](resources--aws_vpc_site--properties--vpc--new_vpc.md#schema-vpc--new_vpc--name_tag) |
| `vpc.new_vpc.primary_ipv4` | [vpc.new_vpc.primary_ipv4](resources--aws_vpc_site--properties--vpc--new_vpc.md#schema-vpc--new_vpc--primary_ipv4) |
| `vpc.vpc_id` | [vpc.vpc_id](resources--aws_vpc_site--properties--vpc.md#schema-vpc--vpc_id) |
| `waf_signatures` | [waf_signatures](resources--aws_vpc_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--aws_vpc_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--aws_vpc_site--properties--waf_signatures--manual.md#section) |

## Next pages

- [admin_password](resources--aws_vpc_site--properties--admin_password.md)
- [aws_cred](resources--aws_vpc_site--properties--aws_cred.md)
- [block_all_services](resources--aws_vpc_site--properties--block_all_services.md)
- [blocked_services](resources--aws_vpc_site--properties--blocked_services.md)
- [coordinates](resources--aws_vpc_site--properties--coordinates.md)
- [custom_dns](resources--aws_vpc_site--properties--custom_dns.md)
- [custom_security_group](resources--aws_vpc_site--properties--custom_security_group.md)
- [default_blocked_services](resources--aws_vpc_site--properties--default_blocked_services.md)
- [direct_connect_disabled](resources--aws_vpc_site--properties--direct_connect_disabled.md)
- [direct_connect_enabled](resources--aws_vpc_site--properties--direct_connect_enabled.md)
- [disable_encryption](resources--aws_vpc_site--properties--disable_encryption.md)
- [disable_internet_vip](resources--aws_vpc_site--properties--disable_internet_vip.md)
- [egress_gateway_default](resources--aws_vpc_site--properties--egress_gateway_default.md)
- [egress_nat_gw](resources--aws_vpc_site--properties--egress_nat_gw.md)
- [egress_virtual_private_gateway](resources--aws_vpc_site--properties--egress_virtual_private_gateway.md)
- [enable_encryption](resources--aws_vpc_site--properties--enable_encryption.md)
- [enable_internet_vip](resources--aws_vpc_site--properties--enable_internet_vip.md)
- [f5_orchestrated_routing](resources--aws_vpc_site--properties--f5_orchestrated_routing.md)
- [f5xc_security_group](resources--aws_vpc_site--properties--f5xc_security_group.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_gw](resources--aws_vpc_site--properties--ingress_gw.md)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](resources--aws_vpc_site--properties--log_receiver.md)
- [logs_streaming_disabled](resources--aws_vpc_site--properties--logs_streaming_disabled.md)
- [manual_routing](resources--aws_vpc_site--properties--manual_routing.md)
- [no_worker_nodes](resources--aws_vpc_site--properties--no_worker_nodes.md)
- [offline_survivability_mode](resources--aws_vpc_site--properties--offline_survivability_mode.md)
- [os](resources--aws_vpc_site--properties--os.md)
- [private_connectivity](resources--aws_vpc_site--properties--private_connectivity.md)
- [sw](resources--aws_vpc_site--properties--sw.md)
- [timeouts](resources--aws_vpc_site--properties--timeouts.md)
- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md)
- [vpc](resources--aws_vpc_site--properties--vpc.md)
- [waf_signatures](resources--aws_vpc_site--properties--waf_signatures.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
