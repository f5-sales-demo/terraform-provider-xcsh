---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 76908, "body_sha256": "sha256:0472f2fd634e1509d1a624e4db39ac23696fdd9cdbc47148c78c6ba6c4000919", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "xcsh-docs:data-sources:aws_tgw_site:properties:block_all_services", "xcsh-docs:data-sources:aws_tgw_site:properties:blocked_services", "xcsh-docs:data-sources:aws_tgw_site:properties:coordinates", "xcsh-docs:data-sources:aws_tgw_site:properties:custom_dns", "xcsh-docs:data-sources:aws_tgw_site:properties:default_blocked_services", "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_disabled", "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled", "xcsh-docs:data-sources:aws_tgw_site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:aws_tgw_site:properties:log_receiver", "xcsh-docs:data-sources:aws_tgw_site:properties:logs_streaming_disabled", "xcsh-docs:data-sources:aws_tgw_site:properties:offline_survivability_mode", "xcsh-docs:data-sources:aws_tgw_site:properties:os", "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode", "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity", "xcsh-docs:data-sources:aws_tgw_site:properties:sw", "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments", "xcsh-docs:data-sources:aws_tgw_site:properties:waf_signatures"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:reference", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:fundamentals", "path": "docs/guides/data-sources--aws_tgw_site--reference.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
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

- [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md): complete subsection reference.

- [block_all_services](data-sources--aws_tgw_site--properties--block_all_services.md): complete subsection reference.

- [blocked_services](data-sources--aws_tgw_site--properties--blocked_services.md): complete subsection reference.

- [coordinates](data-sources--aws_tgw_site--properties--coordinates.md): complete subsection reference.

- [custom_dns](data-sources--aws_tgw_site--properties--custom_dns.md): complete subsection reference.

- [default_blocked_services](data-sources--aws_tgw_site--properties--default_blocked_services.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AWSTGWSite.

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

- [direct_connect_disabled](data-sources--aws_tgw_site--properties--direct_connect_disabled.md): complete subsection reference.

- [direct_connect_enabled](data-sources--aws_tgw_site--properties--direct_connect_enabled.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](data-sources--aws_tgw_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](data-sources--aws_tgw_site--properties--logs_streaming_disabled.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AWSTGWSite.

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

Namespace where the AWSTGWSite exists.

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

- [offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](data-sources--aws_tgw_site--properties--os.md): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_tgw_site--properties--performance_enhancement_mode.md): complete subsection reference.

- [private_connectivity](data-sources--aws_tgw_site--properties--private_connectivity.md): complete subsection reference.

- [sw](data-sources--aws_tgw_site--properties--sw.md): complete subsection reference.

<a id="schema-tags"></a>

### tags property

Type: `["map", "string"]`. Computed.

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

- [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md): complete subsection reference.

- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md): complete subsection reference.

- [vpc_attachments](data-sources--aws_tgw_site--properties--vpc_attachments.md): complete subsection reference.

- [waf_signatures](data-sources--aws_tgw_site--properties--waf_signatures.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--aws_tgw_site--reference.md#schema-annotations) |
| `aws_parameters` | [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md#section) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](data-sources--aws_tgw_site--properties--aws_parameters--admin_password.md#section) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#section) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#schema-aws_parameters--admin_password--blindfold_secret_info--decryption_provider) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#schema-aws_parameters--admin_password--blindfold_secret_info--location) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--blindfold_secret_info.md#schema-aws_parameters--admin_password--blindfold_secret_info--store_provider) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md#section) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md#schema-aws_parameters--admin_password--clear_secret_info--provider_ref) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](data-sources--aws_tgw_site--properties--aws_parameters--admin_password--clear_secret_info.md#schema-aws_parameters--admin_password--clear_secret_info--url) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](data-sources--aws_tgw_site--properties--aws_parameters--aws_cred.md#section) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](data-sources--aws_tgw_site--properties--aws_parameters--aws_cred.md#schema-aws_parameters--aws_cred--name) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](data-sources--aws_tgw_site--properties--aws_parameters--aws_cred.md#schema-aws_parameters--aws_cred--namespace) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](data-sources--aws_tgw_site--properties--aws_parameters--aws_cred.md#schema-aws_parameters--aws_cred--tenant) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--aws_region) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes.md#section) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes.md#schema-aws_parameters--az_nodes--aws_az_name) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet.md#section) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet.md#schema-aws_parameters--az_nodes--inside_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet--subnet_param.md#section) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet--subnet_param.md#schema-aws_parameters--az_nodes--inside_subnet--subnet_param--ipv4) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet.md#section) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet.md#schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet--subnet_param.md#section) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet--subnet_param.md#schema-aws_parameters--az_nodes--outside_subnet--subnet_param--ipv4) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--reserved_inside_subnet.md#section) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet.md#section) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet.md#schema-aws_parameters--az_nodes--workload_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet--subnet_param.md#section) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet--subnet_param.md#schema-aws_parameters--az_nodes--workload_subnet--subnet_param--ipv4) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](data-sources--aws_tgw_site--properties--aws_parameters--custom_security_group.md#section) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](data-sources--aws_tgw_site--properties--aws_parameters--custom_security_group.md#schema-aws_parameters--custom_security_group--inside_security_group_id) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](data-sources--aws_tgw_site--properties--aws_parameters--custom_security_group.md#schema-aws_parameters--custom_security_group--outside_security_group_id) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](data-sources--aws_tgw_site--properties--aws_parameters--disable_encryption.md#section) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](data-sources--aws_tgw_site--properties--aws_parameters--disable_internet_vip.md#section) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--disk_size) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](data-sources--aws_tgw_site--properties--aws_parameters--enable_encryption.md#section) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](data-sources--aws_tgw_site--properties--aws_parameters--enable_encryption.md#schema-aws_parameters--enable_encryption--kms_key_id) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](data-sources--aws_tgw_site--properties--aws_parameters--enable_internet_vip.md#section) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](data-sources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#section) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](data-sources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#schema-aws_parameters--existing_tgw--tgw_asn) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](data-sources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#schema-aws_parameters--existing_tgw--tgw_id) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](data-sources--aws_tgw_site--properties--aws_parameters--existing_tgw.md#schema-aws_parameters--existing_tgw--volterra_site_asn) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](data-sources--aws_tgw_site--properties--aws_parameters--f5xc_security_group.md#section) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--instance_type) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw.md#section) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--system_generated.md#section) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md#section) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md#schema-aws_parameters--new_tgw--user_assigned--tgw_asn) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](data-sources--aws_tgw_site--properties--aws_parameters--new_tgw--user_assigned.md#schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](data-sources--aws_tgw_site--properties--aws_parameters--new_vpc.md#section) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](data-sources--aws_tgw_site--properties--aws_parameters--new_vpc--autogenerate.md#section) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](data-sources--aws_tgw_site--properties--aws_parameters--new_vpc.md#schema-aws_parameters--new_vpc--name_tag) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](data-sources--aws_tgw_site--properties--aws_parameters--new_vpc.md#schema-aws_parameters--new_vpc--primary_ipv4) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](data-sources--aws_tgw_site--properties--aws_parameters--no_worker_nodes.md#section) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--nodes_per_az) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](data-sources--aws_tgw_site--properties--aws_parameters--reserved_tgw_cidr.md#section) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--ssh_key) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](data-sources--aws_tgw_site--properties--aws_parameters--tgw_cidr.md#section) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](data-sources--aws_tgw_site--properties--aws_parameters--tgw_cidr.md#schema-aws_parameters--tgw_cidr--ipv4) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--total_nodes) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](data-sources--aws_tgw_site--properties--aws_parameters.md#schema-aws_parameters--vpc_id) |
| `block_all_services` | [block_all_services](data-sources--aws_tgw_site--properties--block_all_services.md#section) |
| `blocked_services` | [blocked_services](data-sources--aws_tgw_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--aws_tgw_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--aws_tgw_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--aws_tgw_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--aws_tgw_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--aws_tgw_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `coordinates` | [coordinates](data-sources--aws_tgw_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--aws_tgw_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--aws_tgw_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](data-sources--aws_tgw_site--properties--custom_dns.md#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--aws_tgw_site--properties--custom_dns.md#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--aws_tgw_site--properties--custom_dns.md#schema-custom_dns--outside_nameserver) |
| `default_blocked_services` | [default_blocked_services](data-sources--aws_tgw_site--properties--default_blocked_services.md#section) |
| `description` | [description](data-sources--aws_tgw_site--reference.md#schema-description) |
| `direct_connect_disabled` | [direct_connect_disabled](data-sources--aws_tgw_site--properties--direct_connect_disabled.md#section) |
| `direct_connect_enabled` | [direct_connect_enabled](data-sources--aws_tgw_site--properties--direct_connect_enabled.md#section) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](data-sources--aws_tgw_site--properties--direct_connect_enabled--auto_asn.md#section) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](data-sources--aws_tgw_site--properties--direct_connect_enabled.md#schema-direct_connect_enabled--custom_asn) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs.md#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md#schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_internet.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#schema-direct_connect_enabled--hosted_vifs--vif_list--other_region) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list--same_as_site_region.md#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](data-sources--aws_tgw_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md#schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](data-sources--aws_tgw_site--properties--direct_connect_enabled--standard_vifs.md#section) |
| `id` | [id](data-sources--aws_tgw_site--reference.md#schema-id) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](data-sources--aws_tgw_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](data-sources--aws_tgw_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](data-sources--aws_tgw_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--aws_tgw_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--aws_tgw_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--aws_tgw_site--properties--logs_streaming_disabled.md#section) |
| `name` | [name](data-sources--aws_tgw_site--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--aws_tgw_site--reference.md#schema-namespace) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](data-sources--aws_tgw_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](data-sources--aws_tgw_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--aws_tgw_site--properties--os.md#schema-os--operating_system_version) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--aws_tgw_site--properties--performance_enhancement_mode.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `private_connectivity` | [private_connectivity](data-sources--aws_tgw_site--properties--private_connectivity.md#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--aws_tgw_site--properties--private_connectivity--cloud_link.md#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--aws_tgw_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--aws_tgw_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--namespace) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--aws_tgw_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--tenant) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--aws_tgw_site--properties--private_connectivity--inside.md#section) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--aws_tgw_site--properties--private_connectivity--outside.md#section) |
| `sw` | [sw](data-sources--aws_tgw_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--aws_tgw_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--aws_tgw_site--properties--sw.md#schema-sw--volterra_software_version) |
| `tags` | [tags](data-sources--aws_tgw_site--reference.md#schema-tags) |
| `tgw_security` | [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md#section) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies.md#section) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#section) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#schema-tgw_security--active_east_west_service_policies--service_policies--name) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#schema-tgw_security--active_east_west_service_policies--service_policies--namespace) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](data-sources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies--service_policies.md#schema-tgw_security--active_east_west_service_policies--service_policies--tenant) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies.md#section) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies.md#section) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--name) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies--forward_proxy_policies.md#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](data-sources--aws_tgw_site--properties--tgw_security--active_network_policies.md#section) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](data-sources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#section) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](data-sources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#schema-tgw_security--active_network_policies--network_policies--name) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](data-sources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#schema-tgw_security--active_network_policies--network_policies--namespace) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](data-sources--aws_tgw_site--properties--tgw_security--active_network_policies--network_policies.md#schema-tgw_security--active_network_policies--network_policies--tenant) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](data-sources--aws_tgw_site--properties--tgw_security--east_west_service_policy_allow_all.md#section) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](data-sources--aws_tgw_site--properties--tgw_security--forward_proxy_allow_all.md#section) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](data-sources--aws_tgw_site--properties--tgw_security--no_east_west_policy.md#section) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](data-sources--aws_tgw_site--properties--tgw_security--no_forward_proxy.md#section) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](data-sources--aws_tgw_site--properties--tgw_security--no_network_policy.md#section) |
| `vn_config` | [vn_config](data-sources--aws_tgw_site--properties--vn_config.md#section) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port.md#section) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port--custom_ports.md#section) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port--custom_ports.md#schema-vn_config--allowed_vip_port--custom_ports--port_ranges) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port--disable_allowed_vip_port.md#section) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port--use_http_https_port.md#section) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port--use_http_port.md#section) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port--use_https_port.md#section) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli.md#section) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--custom_ports.md#section) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--custom_ports.md#schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--disable_allowed_vip_port.md#section) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_https_port.md#section) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_http_port.md#section) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](data-sources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli--use_https_port.md#section) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#section) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#schema-vn_config--dc_cluster_group_inside_vn--name) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#schema-vn_config--dc_cluster_group_inside_vn--namespace) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md#schema-vn_config--dc_cluster_group_inside_vn--tenant) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#section) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#schema-vn_config--dc_cluster_group_outside_vn--name) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#schema-vn_config--dc_cluster_group_outside_vn--namespace) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](data-sources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md#schema-vn_config--dc_cluster_group_outside_vn--tenant) |
| `vn_config.global_network_list` | [vn_config.global_network_list](data-sources--aws_tgw_site--properties--vn_config--global_network_list.md#section) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes.md#section) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](data-sources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list.md#schema-vn_config--inside_static_routes--static_route_list--simple_static_route) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](data-sources--aws_tgw_site--properties--vn_config--no_dc_cluster_group.md#section) |
| `vn_config.no_global_network` | [vn_config.no_global_network](data-sources--aws_tgw_site--properties--vn_config--no_global_network.md#section) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](data-sources--aws_tgw_site--properties--vn_config--no_inside_static_routes.md#section) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](data-sources--aws_tgw_site--properties--vn_config--no_outside_static_routes.md#section) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes.md#section) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md#schema-vn_config--outside_static_routes--static_route_list--simple_static_route) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](data-sources--aws_tgw_site--properties--vn_config--sm_connection_public_ip.md#section) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](data-sources--aws_tgw_site--properties--vn_config--sm_connection_pvt_ip.md#section) |
| `vpc_attachments` | [vpc_attachments](data-sources--aws_tgw_site--properties--vpc_attachments.md#section) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list.md#section) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list--labels.md#section) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list.md#schema-vpc_attachments--vpc_list--vpc_id) |
| `waf_signatures` | [waf_signatures](data-sources--aws_tgw_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--aws_tgw_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--aws_tgw_site--properties--waf_signatures--manual.md#section) |

## Next pages

- [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md)
- [block_all_services](data-sources--aws_tgw_site--properties--block_all_services.md)
- [blocked_services](data-sources--aws_tgw_site--properties--blocked_services.md)
- [coordinates](data-sources--aws_tgw_site--properties--coordinates.md)
- [custom_dns](data-sources--aws_tgw_site--properties--custom_dns.md)
- [default_blocked_services](data-sources--aws_tgw_site--properties--default_blocked_services.md)
- [direct_connect_disabled](data-sources--aws_tgw_site--properties--direct_connect_disabled.md)
- [direct_connect_enabled](data-sources--aws_tgw_site--properties--direct_connect_enabled.md)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](data-sources--aws_tgw_site--properties--log_receiver.md)
- [logs_streaming_disabled](data-sources--aws_tgw_site--properties--logs_streaming_disabled.md)
- [offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode.md)
- [os](data-sources--aws_tgw_site--properties--os.md)
- [performance_enhancement_mode](data-sources--aws_tgw_site--properties--performance_enhancement_mode.md)
- [private_connectivity](data-sources--aws_tgw_site--properties--private_connectivity.md)
- [sw](data-sources--aws_tgw_site--properties--sw.md)
- [tgw_security](data-sources--aws_tgw_site--properties--tgw_security.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- [vpc_attachments](data-sources--aws_tgw_site--properties--vpc_attachments.md)
- [waf_signatures](data-sources--aws_tgw_site--properties--waf_signatures.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
