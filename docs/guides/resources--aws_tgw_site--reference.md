---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 77589, "body_sha256": "sha256:a60a68d6029412ae0d9ea162b13a0700a2fff7b381aa4e1745bb5f54f8b4393d", "canonical_id": "xcsh-docs:resources:aws_tgw_site:reference", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "xcsh-docs:resources:aws_tgw_site:properties:block_all_services", "xcsh-docs:resources:aws_tgw_site:properties:blocked_services", "xcsh-docs:resources:aws_tgw_site:properties:coordinates", "xcsh-docs:resources:aws_tgw_site:properties:custom_dns", "xcsh-docs:resources:aws_tgw_site:properties:default_blocked_services", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_disabled", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain", "xcsh-docs:resources:aws_tgw_site:properties:log_receiver", "xcsh-docs:resources:aws_tgw_site:properties:logs_streaming_disabled", "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode", "xcsh-docs:resources:aws_tgw_site:properties:os", "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode", "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity", "xcsh-docs:resources:aws_tgw_site:properties:sw", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "xcsh-docs:resources:aws_tgw_site:properties:timeouts", "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:reference", "parent_id": "xcsh-docs:resources:aws_tgw_site:fundamentals", "path": "docs/guides/resources--aws_tgw_site--reference.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
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

- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md): complete subsection reference.

- [block_all_services](resources--aws_tgw_site--properties--block_all_services.md): complete subsection reference.

- [blocked_services](resources--aws_tgw_site--properties--blocked_services.md): complete subsection reference.

- [coordinates](resources--aws_tgw_site--properties--coordinates.md): complete subsection reference.

- [custom_dns](resources--aws_tgw_site--properties--custom_dns.md): complete subsection reference.

- [default_blocked_services](resources--aws_tgw_site--properties--default_blocked_services.md): complete subsection reference.

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

- [direct_connect_disabled](resources--aws_tgw_site--properties--direct_connect_disabled.md): complete subsection reference.

- [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md): complete subsection reference.

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

- [kubernetes_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](resources--aws_tgw_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](resources--aws_tgw_site--properties--logs_streaming_disabled.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AWS TGW Site. Must be unique within the namespace.

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

Namespace where the AWS TGW Site is created.

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

- [offline_survivability_mode](resources--aws_tgw_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](resources--aws_tgw_site--properties--os.md): complete subsection reference.

- [performance_enhancement_mode](resources--aws_tgw_site--properties--performance_enhancement_mode.md): complete subsection reference.

- [private_connectivity](resources--aws_tgw_site--properties--private_connectivity.md): complete subsection reference.

- [sw](resources--aws_tgw_site--properties--sw.md): complete subsection reference.

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

- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md): complete subsection reference.

- [timeouts](resources--aws_tgw_site--properties--timeouts.md): complete subsection reference.

- [vn_config](resources--aws_tgw_site--properties--vn_config.md): complete subsection reference.

- [vpc_attachments](resources--aws_tgw_site--properties--vpc_attachments.md): complete subsection reference.

- [waf_signatures](resources--aws_tgw_site--properties--waf_signatures.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--aws_tgw_site--reference.md#schema-annotations) |
| `aws_parameters` | [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md#section) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](resources--aws_tgw_site--properties--aws_parameters--admin_password.md#section) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](resources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#section) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](resources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#schema-aws_parameters--admin_password--blindfold_secret_info--decryption_provider) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](resources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#schema-aws_parameters--admin_password--blindfold_secret_info--location) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](resources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#schema-aws_parameters--admin_password--blindfold_secret_info--store_provider) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](resources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md#section) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](resources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md#schema-aws_parameters--admin_password--clear_secret_info--provider_ref) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](resources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md#schema-aws_parameters--admin_password--clear_secret_info--url) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](resources--aws_tgw_site--properties--aws_parameters--aws_cred.md#section) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](resources--aws_tgw_site--properties--aws_parameters--aws_cred.md#schema-aws_parameters--aws_cred--name) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](resources--aws_tgw_site--properties--aws_parameters--aws_cred.md#schema-aws_parameters--aws_cred--namespace) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](resources--aws_tgw_site--properties--aws_parameters--aws_cred.md#schema-aws_parameters--aws_cred--tenant) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--aws_region) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](resources--aws_tgw_site--properties--aws_parameters--az_nodes.md#section) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](resources--aws_tgw_site--properties--aws_parameters--az_nodes.md#schema-aws_parameters--az_nodes--aws_az_name) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet.md#section) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](resources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet.md#schema-aws_parameters--az_nodes--inside_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](resources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet--subnet_param.md#section) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](resources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet--subnet_param.md#schema-aws_parameters--az_nodes--inside_subnet--subnet_param--ipv4) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet.md#section) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](resources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet.md#schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](resources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet--subnet_param.md#section) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](resources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet--subnet_param.md#schema-aws_parameters--az_nodes--outside_subnet--subnet_param--ipv4) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--reserved_inside_subnet.md#section) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet.md#section) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](resources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet.md#schema-aws_parameters--az_nodes--workload_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](resources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet--subnet_param.md#section) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](resources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet--subnet_param.md#schema-aws_parameters--az_nodes--workload_subnet--subnet_param--ipv4) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](resources--aws_tgw_site--properties--aws_parameters--custom_security_group.md#section) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](resources--aws_tgw_site--properties--aws_parameters--custom_security_group.md#schema-aws_parameters--custom_security_group--inside_security_group_id) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](resources--aws_tgw_site--properties--aws_parameters--custom_security_group.md#schema-aws_parameters--custom_security_group--outside_security_group_id) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](resources--aws_tgw_site--properties--aws_parameters--disable_encryption.md#section) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](resources--aws_tgw_site--properties--aws_parameters--disable_internet_vip.md#section) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--disk_size) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](resources--aws_tgw_site--properties--aws_parameters--enable_encryption.md#section) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](resources--aws_tgw_site--properties--aws_parameters--enable_encryption.md#schema-aws_parameters--enable_encryption--kms_key_id) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](resources--aws_tgw_site--properties--aws_parameters--enable_internet_vip.md#section) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](resources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#section) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](resources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#schema-aws_parameters--existing_tgw--tgw_asn) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](resources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#schema-aws_parameters--existing_tgw--tgw_id) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](resources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#schema-aws_parameters--existing_tgw--volterra_site_asn) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](resources--aws_tgw_site--properties--aws_parameters--f5xc_security_group.md#section) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--instance_type) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](resources--aws_tgw_site--properties--aws_parameters--new_tgw.md#section) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](resources--aws_tgw_site--properties--aws_parameters--new_tgw--system_generated.md#section) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](resources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md#section) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](resources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md#schema-aws_parameters--new_tgw--user_assigned--tgw_asn) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](resources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md#schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](resources--aws_tgw_site--properties--aws_parameters--new_vpc.md#section) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](resources--aws_tgw_site--properties--aws_parameters--new_vpc--autogenerate.md#section) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](resources--aws_tgw_site--properties--aws_parameters--new_vpc.md#schema-aws_parameters--new_vpc--name_tag) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](resources--aws_tgw_site--properties--aws_parameters--new_vpc.md#schema-aws_parameters--new_vpc--primary_ipv4) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](resources--aws_tgw_site--properties--aws_parameters--no_worker_nodes.md#section) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--nodes_per_az) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](resources--aws_tgw_site--properties--aws_parameters--reserved_tgw_cidr.md#section) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--ssh_key) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](resources--aws_tgw_site--properties--aws_parameters--tgw_cidr.md#section) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](resources--aws_tgw_site--properties--aws_parameters--tgw_cidr.md#schema-aws_parameters--tgw_cidr--ipv4) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--total_nodes) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](resources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--vpc_id) |
| `block_all_services` | [block_all_services](resources--aws_tgw_site--properties--block_all_services.md#section) |
| `blocked_services` | [blocked_services](resources--aws_tgw_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--aws_tgw_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--aws_tgw_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--aws_tgw_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--aws_tgw_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--aws_tgw_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `coordinates` | [coordinates](resources--aws_tgw_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](resources--aws_tgw_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](resources--aws_tgw_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](resources--aws_tgw_site--properties--custom_dns.md#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--aws_tgw_site--properties--custom_dns.md#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--aws_tgw_site--properties--custom_dns.md#schema-custom_dns--outside_nameserver) |
| `default_blocked_services` | [default_blocked_services](resources--aws_tgw_site--properties--default_blocked_services.md#section) |
| `description` | [description](resources--aws_tgw_site--reference.md#schema-description) |
| `direct_connect_disabled` | [direct_connect_disabled](resources--aws_tgw_site--properties--direct_connect_disabled.md#section) |
| `direct_connect_enabled` | [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md#section) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](resources--aws_tgw_site--properties--direct_connect_enabled--auto_asn.md#section) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](resources--aws_tgw_site--properties--direct_connect_enabled.md#schema-direct_connect_enabled--custom_asn) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs.md#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md#schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#schema-direct_connect_enabled--hosted_vifs--vif_list--other_region) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list--same_as_site_region.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](resources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](resources--aws_tgw_site--properties--direct_connect_enabled--standard_vifs.md#section) |
| `disable` | [disable](resources--aws_tgw_site--reference.md#schema-disable) |
| `id` | [id](resources--aws_tgw_site--reference.md#schema-id) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](resources--aws_tgw_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](resources--aws_tgw_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](resources--aws_tgw_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--aws_tgw_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--aws_tgw_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--aws_tgw_site--properties--logs_streaming_disabled.md#section) |
| `name` | [name](resources--aws_tgw_site--reference.md#schema-name) |
| `namespace` | [namespace](resources--aws_tgw_site--reference.md#schema-namespace) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--aws_tgw_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_tgw_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--aws_tgw_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](resources--aws_tgw_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](resources--aws_tgw_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](resources--aws_tgw_site--properties--os.md#schema-os--operating_system_version) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--aws_tgw_site--properties--performance_enhancement_mode.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `private_connectivity` | [private_connectivity](resources--aws_tgw_site--properties--private_connectivity.md#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--aws_tgw_site--properties--private_connectivity--cloud_link.md#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--aws_tgw_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--aws_tgw_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--namespace) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--aws_tgw_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--tenant) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--aws_tgw_site--properties--private_connectivity--inside.md#section) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--aws_tgw_site--properties--private_connectivity--outside.md#section) |
| `sw` | [sw](resources--aws_tgw_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--aws_tgw_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--aws_tgw_site--properties--sw.md#schema-sw--volterra_software_version) |
| `tags` | [tags](resources--aws_tgw_site--reference.md#schema-tags) |
| `tgw_security` | [tgw_security](resources--aws_tgw_site--properties--tgw_security.md#section) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies.md#section) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#section) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#schema-tgw_security--active_east_west_service_policies--service_policies--name) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#schema-tgw_security--active_east_west_service_policies--service_policies--namespace) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#schema-tgw_security--active_east_west_service_policies--service_policies--tenant) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies.md#section) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies.md#section) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--name) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](resources--aws_tgw_site--properties--tgw_security--active_network_policies.md#section) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](resources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#section) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](resources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#schema-tgw_security--active_network_policies--network_policies--name) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](resources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#schema-tgw_security--active_network_policies--network_policies--namespace) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](resources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#schema-tgw_security--active_network_policies--network_policies--tenant) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](resources--aws_tgw_site--properties--tgw_security--east_west_service_policy_allow_all.md#section) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](resources--aws_tgw_site--properties--tgw_security--forward_proxy_allow_all.md#section) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](resources--aws_tgw_site--properties--tgw_security--no_east_west_policy.md#section) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](resources--aws_tgw_site--properties--tgw_security--no_forward_proxy.md#section) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](resources--aws_tgw_site--properties--tgw_security--no_network_policy.md#section) |
| `timeouts` | [timeouts](resources--aws_tgw_site--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--aws_tgw_site--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--aws_tgw_site--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--aws_tgw_site--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--aws_tgw_site--properties--timeouts.md#schema-timeouts--update) |
| `vn_config` | [vn_config](resources--aws_tgw_site--properties--vn_config.md#section) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port.md#section) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](resources--aws_tgw_site--properties--vn_config--allowed_vip_port--custom_ports.md#section) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](resources--aws_tgw_site--properties--vn_config--allowed_vip_port--custom_ports.md#schema-vn_config--allowed_vip_port--custom_ports--port_ranges) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port--disable_allowed_vip_port.md#section) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port--use_http_https_port.md#section) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port--use_http_port.md#section) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port--use_https_port.md#section) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli.md#section) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--custom_ports.md#section) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--custom_ports.md#schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--disable_allowed_vip_port.md#section) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_https_port.md#section) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_port.md#section) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_https_port.md#section) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#section) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#schema-vn_config--dc_cluster_group_inside_vn--name) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#schema-vn_config--dc_cluster_group_inside_vn--namespace) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#schema-vn_config--dc_cluster_group_inside_vn--tenant) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#section) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#schema-vn_config--dc_cluster_group_outside_vn--name) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#schema-vn_config--dc_cluster_group_outside_vn--namespace) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#schema-vn_config--dc_cluster_group_outside_vn--tenant) |
| `vn_config.global_network_list` | [vn_config.global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md#section) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](resources--aws_tgw_site--properties--vn_config--inside_static_routes.md#section) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list.md#schema-vn_config--inside_static_routes--static_route_list--simple_static_route) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](resources--aws_tgw_site--properties--vn_config--no_dc_cluster_group.md#section) |
| `vn_config.no_global_network` | [vn_config.no_global_network](resources--aws_tgw_site--properties--vn_config--no_global_network.md#section) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](resources--aws_tgw_site--properties--vn_config--no_inside_static_routes.md#section) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](resources--aws_tgw_site--properties--vn_config--no_outside_static_routes.md#section) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](resources--aws_tgw_site--properties--vn_config--outside_static_routes.md#section) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md#schema-vn_config--outside_static_routes--static_route_list--simple_static_route) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](resources--aws_tgw_site--properties--vn_config--sm_connection_public_ip.md#section) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](resources--aws_tgw_site--properties--vn_config--sm_connection_pvt_ip.md#section) |
| `vpc_attachments` | [vpc_attachments](resources--aws_tgw_site--properties--vpc_attachments.md#section) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](resources--aws_tgw_site--properties--vpc_attachments--vpc_list.md#section) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](resources--aws_tgw_site--properties--vpc_attachments--vpc_list--labels.md#section) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](resources--aws_tgw_site--properties--vpc_attachments--vpc_list.md#schema-vpc_attachments--vpc_list--vpc_id) |
| `waf_signatures` | [waf_signatures](resources--aws_tgw_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--aws_tgw_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--aws_tgw_site--properties--waf_signatures--manual.md#section) |

## Next pages

- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
- [block_all_services](resources--aws_tgw_site--properties--block_all_services.md)
- [blocked_services](resources--aws_tgw_site--properties--blocked_services.md)
- [coordinates](resources--aws_tgw_site--properties--coordinates.md)
- [custom_dns](resources--aws_tgw_site--properties--custom_dns.md)
- [default_blocked_services](resources--aws_tgw_site--properties--default_blocked_services.md)
- [direct_connect_disabled](resources--aws_tgw_site--properties--direct_connect_disabled.md)
- [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](resources--aws_tgw_site--properties--log_receiver.md)
- [logs_streaming_disabled](resources--aws_tgw_site--properties--logs_streaming_disabled.md)
- [offline_survivability_mode](resources--aws_tgw_site--properties--offline_survivability_mode.md)
- [os](resources--aws_tgw_site--properties--os.md)
- [performance_enhancement_mode](resources--aws_tgw_site--properties--performance_enhancement_mode.md)
- [private_connectivity](resources--aws_tgw_site--properties--private_connectivity.md)
- [sw](resources--aws_tgw_site--properties--sw.md)
- [tgw_security](resources--aws_tgw_site--properties--tgw_security.md)
- [timeouts](resources--aws_tgw_site--properties--timeouts.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vpc_attachments](resources--aws_tgw_site--properties--vpc_attachments.md)
- [waf_signatures](resources--aws_tgw_site--properties--waf_signatures.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
