---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 108153, "body_sha256": "sha256:b823c7d67bebe4cfd4c983f9c901b95f48f9c0d7ce42530c32d8984af0bc14c9", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:admin_password", "xcsh-docs:data-sources:gcp_vpc_site:properties:block_all_services", "xcsh-docs:data-sources:gcp_vpc_site:properties:blocked_services", "xcsh-docs:data-sources:gcp_vpc_site:properties:cloud_credentials", "xcsh-docs:data-sources:gcp_vpc_site:properties:coordinates", "xcsh-docs:data-sources:gcp_vpc_site:properties:custom_dns", "xcsh-docs:data-sources:gcp_vpc_site:properties:default_blocked_services", "xcsh-docs:data-sources:gcp_vpc_site:properties:disable_encryption", "xcsh-docs:data-sources:gcp_vpc_site:properties:enable_encryption", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw", "xcsh-docs:data-sources:gcp_vpc_site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:gcp_vpc_site:properties:log_receiver", "xcsh-docs:data-sources:gcp_vpc_site:properties:logs_streaming_disabled", "xcsh-docs:data-sources:gcp_vpc_site:properties:offline_survivability_mode", "xcsh-docs:data-sources:gcp_vpc_site:properties:os", "xcsh-docs:data-sources:gcp_vpc_site:properties:private_connect_disabled", "xcsh-docs:data-sources:gcp_vpc_site:properties:private_connectivity", "xcsh-docs:data-sources:gcp_vpc_site:properties:sw", "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster", "xcsh-docs:data-sources:gcp_vpc_site:properties:waf_signatures"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:fundamentals", "path": "docs/guides/data-sources--gcp_vpc_site--reference.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

- [admin_password](data-sources--gcp_vpc_site--properties--admin_password.md): complete subsection reference.

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

- [block_all_services](data-sources--gcp_vpc_site--properties--block_all_services.md): complete subsection reference.

- [blocked_services](data-sources--gcp_vpc_site--properties--blocked_services.md): complete subsection reference.

- [cloud_credentials](data-sources--gcp_vpc_site--properties--cloud_credentials.md): complete subsection reference.

- [coordinates](data-sources--gcp_vpc_site--properties--coordinates.md): complete subsection reference.

- [custom_dns](data-sources--gcp_vpc_site--properties--custom_dns.md): complete subsection reference.

- [default_blocked_services](data-sources--gcp_vpc_site--properties--default_blocked_services.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the GCPVPCSite.

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

- [disable_encryption](data-sources--gcp_vpc_site--properties--disable_encryption.md): complete subsection reference.

<a id="schema-disk_size"></a>

### disk_size property

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](data-sources--gcp_vpc_site--properties--enable_encryption.md): complete subsection reference.

<a id="schema-gcp_labels"></a>

### gcp_labels property

Type: `["map", "string"]`. Computed.

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

Upstream description:

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

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

<a id="schema-gcp_region"></a>

### gcp_region property

Type: `"string"`. Computed.

GCP Region. Name for GCP Region.

Upstream description:

Name for GCP Region.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md): complete subsection reference.

- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md): complete subsection reference.

<a id="schema-instance_type"></a>

### instance_type property

Type: `"string"`. Computed.

Select Instance size based on performance needed.

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

- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](data-sources--gcp_vpc_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](data-sources--gcp_vpc_site--properties--logs_streaming_disabled.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the GCPVPCSite.

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

Namespace where the GCPVPCSite exists.

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

- [offline_survivability_mode](data-sources--gcp_vpc_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](data-sources--gcp_vpc_site--properties--os.md): complete subsection reference.

- [private_connect_disabled](data-sources--gcp_vpc_site--properties--private_connect_disabled.md): complete subsection reference.

- [private_connectivity](data-sources--gcp_vpc_site--properties--private_connectivity.md): complete subsection reference.

<a id="schema-ssh_key"></a>

### ssh_key property

Type: `"string"`. Computed.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

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

- [sw](data-sources--gcp_vpc_site--properties--sw.md): complete subsection reference.

- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md): complete subsection reference.

- [waf_signatures](data-sources--gcp_vpc_site--properties--waf_signatures.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--gcp_vpc_site--reference.md#schema-address) |
| `admin_password` | [admin_password](data-sources--gcp_vpc_site--properties--admin_password.md#section) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](data-sources--gcp_vpc_site--properties--admin_password--blindfold_secret_info.md#section) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](data-sources--gcp_vpc_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--decryption_provider) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](data-sources--gcp_vpc_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--location) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](data-sources--gcp_vpc_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--store_provider) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](data-sources--gcp_vpc_site--properties--admin_password--clear_secret_info.md#section) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](data-sources--gcp_vpc_site--properties--admin_password--clear_secret_info.md#schema-admin_password--clear_secret_info--provider_ref) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](data-sources--gcp_vpc_site--properties--admin_password--clear_secret_info.md#schema-admin_password--clear_secret_info--url) |
| `annotations` | [annotations](data-sources--gcp_vpc_site--reference.md#schema-annotations) |
| `block_all_services` | [block_all_services](data-sources--gcp_vpc_site--properties--block_all_services.md#section) |
| `blocked_services` | [blocked_services](data-sources--gcp_vpc_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `cloud_credentials` | [cloud_credentials](data-sources--gcp_vpc_site--properties--cloud_credentials.md#section) |
| `cloud_credentials.name` | [cloud_credentials.name](data-sources--gcp_vpc_site--properties--cloud_credentials.md#schema-cloud_credentials--name) |
| `cloud_credentials.namespace` | [cloud_credentials.namespace](data-sources--gcp_vpc_site--properties--cloud_credentials.md#schema-cloud_credentials--namespace) |
| `cloud_credentials.tenant` | [cloud_credentials.tenant](data-sources--gcp_vpc_site--properties--cloud_credentials.md#schema-cloud_credentials--tenant) |
| `coordinates` | [coordinates](data-sources--gcp_vpc_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--gcp_vpc_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--gcp_vpc_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](data-sources--gcp_vpc_site--properties--custom_dns.md#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--gcp_vpc_site--properties--custom_dns.md#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--gcp_vpc_site--properties--custom_dns.md#schema-custom_dns--outside_nameserver) |
| `default_blocked_services` | [default_blocked_services](data-sources--gcp_vpc_site--properties--default_blocked_services.md#section) |
| `description` | [description](data-sources--gcp_vpc_site--reference.md#schema-description) |
| `disable_encryption` | [disable_encryption](data-sources--gcp_vpc_site--properties--disable_encryption.md#section) |
| `disk_size` | [disk_size](data-sources--gcp_vpc_site--reference.md#schema-disk_size) |
| `enable_encryption` | [enable_encryption](data-sources--gcp_vpc_site--properties--enable_encryption.md#section) |
| `enable_encryption.kms_key_resource_id` | [enable_encryption.kms_key_resource_id](data-sources--gcp_vpc_site--properties--enable_encryption.md#schema-enable_encryption--kms_key_resource_id) |
| `enable_encryption.kms_key_ring_id` | [enable_encryption.kms_key_ring_id](data-sources--gcp_vpc_site--properties--enable_encryption.md#schema-enable_encryption--kms_key_ring_id) |
| `gcp_labels` | [gcp_labels](data-sources--gcp_vpc_site--reference.md#schema-gcp_labels) |
| `gcp_region` | [gcp_region](data-sources--gcp_vpc_site--reference.md#schema-gcp_region) |
| `id` | [id](data-sources--gcp_vpc_site--reference.md#schema-id) |
| `ingress_egress_gw` | [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies.md#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies.md#section) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#section) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--name) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#section) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#section) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](data-sources--gcp_vpc_site--properties--ingress_egress_gw--forward_proxy_allow_all.md#section) |
| `ingress_egress_gw.gcp_certified_hw` | [ingress_egress_gw.gcp_certified_hw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md#schema-ingress_egress_gw--gcp_certified_hw) |
| `ingress_egress_gw.gcp_zone_names` | [ingress_egress_gw.gcp_zone_names](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md#schema-ingress_egress_gw--gcp_zone_names) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.inside_network` | [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network.md#section) |
| `ingress_egress_gw.inside_network.existing_network` | [ingress_egress_gw.inside_network.existing_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--existing_network.md#section) |
| `ingress_egress_gw.inside_network.existing_network.name` | [ingress_egress_gw.inside_network.existing_network.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--existing_network.md#schema-ingress_egress_gw--inside_network--existing_network--name) |
| `ingress_egress_gw.inside_network.new_network` | [ingress_egress_gw.inside_network.new_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network.md#section) |
| `ingress_egress_gw.inside_network.new_network.name` | [ingress_egress_gw.inside_network.new_network.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network.md#schema-ingress_egress_gw--inside_network--new_network--name) |
| `ingress_egress_gw.inside_network.new_network_autogenerate` | [ingress_egress_gw.inside_network.new_network_autogenerate](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network_autogenerate.md#section) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.inside_subnet` | [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet.md#section) |
| `ingress_egress_gw.inside_subnet.existing_subnet` | [ingress_egress_gw.inside_subnet.existing_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--existing_subnet.md#section) |
| `ingress_egress_gw.inside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--existing_subnet.md#schema-ingress_egress_gw--inside_subnet--existing_subnet--subnet_name) |
| `ingress_egress_gw.inside_subnet.new_subnet` | [ingress_egress_gw.inside_subnet.new_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--new_subnet.md#section) |
| `ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--new_subnet.md#schema-ingress_egress_gw--inside_subnet--new_subnet--primary_ipv4) |
| `ingress_egress_gw.inside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--new_subnet.md#schema-ingress_egress_gw--inside_subnet--new_subnet--subnet_name) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_dc_cluster_group.md#section) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_forward_proxy.md#section) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_global_network.md#section) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_inside_static_routes.md#section) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_network_policy.md#section) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_outside_static_routes.md#section) |
| `ingress_egress_gw.node_number` | [ingress_egress_gw.node_number](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md#schema-ingress_egress_gw--node_number) |
| `ingress_egress_gw.outside_network` | [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network.md#section) |
| `ingress_egress_gw.outside_network.existing_network` | [ingress_egress_gw.outside_network.existing_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--existing_network.md#section) |
| `ingress_egress_gw.outside_network.existing_network.name` | [ingress_egress_gw.outside_network.existing_network.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--existing_network.md#schema-ingress_egress_gw--outside_network--existing_network--name) |
| `ingress_egress_gw.outside_network.new_network` | [ingress_egress_gw.outside_network.new_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network.md#section) |
| `ingress_egress_gw.outside_network.new_network.name` | [ingress_egress_gw.outside_network.new_network.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network.md#schema-ingress_egress_gw--outside_network--new_network--name) |
| `ingress_egress_gw.outside_network.new_network_autogenerate` | [ingress_egress_gw.outside_network.new_network_autogenerate](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network_autogenerate.md#section) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.outside_subnet` | [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet.md#section) |
| `ingress_egress_gw.outside_subnet.existing_subnet` | [ingress_egress_gw.outside_subnet.existing_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet--existing_subnet.md#section) |
| `ingress_egress_gw.outside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet--existing_subnet.md#schema-ingress_egress_gw--outside_subnet--existing_subnet--subnet_name) |
| `ingress_egress_gw.outside_subnet.new_subnet` | [ingress_egress_gw.outside_subnet.new_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet--new_subnet.md#section) |
| `ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet--new_subnet.md#schema-ingress_egress_gw--outside_subnet--new_subnet--primary_ipv4) |
| `ingress_egress_gw.outside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet--new_subnet.md#schema-ingress_egress_gw--outside_subnet--new_subnet--subnet_name) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](data-sources--gcp_vpc_site--properties--ingress_egress_gw--sm_connection_public_ip.md#section) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](data-sources--gcp_vpc_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md#section) |
| `ingress_gw` | [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md#section) |
| `ingress_gw.gcp_certified_hw` | [ingress_gw.gcp_certified_hw](data-sources--gcp_vpc_site--properties--ingress_gw.md#schema-ingress_gw--gcp_certified_hw) |
| `ingress_gw.gcp_zone_names` | [ingress_gw.gcp_zone_names](data-sources--gcp_vpc_site--properties--ingress_gw.md#schema-ingress_gw--gcp_zone_names) |
| `ingress_gw.local_network` | [ingress_gw.local_network](data-sources--gcp_vpc_site--properties--ingress_gw--local_network.md#section) |
| `ingress_gw.local_network.existing_network` | [ingress_gw.local_network.existing_network](data-sources--gcp_vpc_site--properties--ingress_gw--local_network--existing_network.md#section) |
| `ingress_gw.local_network.existing_network.name` | [ingress_gw.local_network.existing_network.name](data-sources--gcp_vpc_site--properties--ingress_gw--local_network--existing_network.md#schema-ingress_gw--local_network--existing_network--name) |
| `ingress_gw.local_network.new_network` | [ingress_gw.local_network.new_network](data-sources--gcp_vpc_site--properties--ingress_gw--local_network--new_network.md#section) |
| `ingress_gw.local_network.new_network.name` | [ingress_gw.local_network.new_network.name](data-sources--gcp_vpc_site--properties--ingress_gw--local_network--new_network.md#schema-ingress_gw--local_network--new_network--name) |
| `ingress_gw.local_network.new_network_autogenerate` | [ingress_gw.local_network.new_network_autogenerate](data-sources--gcp_vpc_site--properties--ingress_gw--local_network--new_network_autogenerate.md#section) |
| `ingress_gw.local_subnet` | [ingress_gw.local_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet.md#section) |
| `ingress_gw.local_subnet.existing_subnet` | [ingress_gw.local_subnet.existing_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--existing_subnet.md#section) |
| `ingress_gw.local_subnet.existing_subnet.subnet_name` | [ingress_gw.local_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--existing_subnet.md#schema-ingress_gw--local_subnet--existing_subnet--subnet_name) |
| `ingress_gw.local_subnet.new_subnet` | [ingress_gw.local_subnet.new_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--new_subnet.md#section) |
| `ingress_gw.local_subnet.new_subnet.primary_ipv4` | [ingress_gw.local_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--new_subnet.md#schema-ingress_gw--local_subnet--new_subnet--primary_ipv4) |
| `ingress_gw.local_subnet.new_subnet.subnet_name` | [ingress_gw.local_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--new_subnet.md#schema-ingress_gw--local_subnet--new_subnet--subnet_name) |
| `ingress_gw.node_number` | [ingress_gw.node_number](data-sources--gcp_vpc_site--properties--ingress_gw.md#schema-ingress_gw--node_number) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `instance_type` | [instance_type](data-sources--gcp_vpc_site--reference.md#schema-instance_type) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](data-sources--gcp_vpc_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](data-sources--gcp_vpc_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](data-sources--gcp_vpc_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--gcp_vpc_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--gcp_vpc_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--gcp_vpc_site--properties--logs_streaming_disabled.md#section) |
| `name` | [name](data-sources--gcp_vpc_site--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--gcp_vpc_site--reference.md#schema-namespace) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--gcp_vpc_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--gcp_vpc_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--gcp_vpc_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](data-sources--gcp_vpc_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](data-sources--gcp_vpc_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--gcp_vpc_site--properties--os.md#schema-os--operating_system_version) |
| `private_connect_disabled` | [private_connect_disabled](data-sources--gcp_vpc_site--properties--private_connect_disabled.md#section) |
| `private_connectivity` | [private_connectivity](data-sources--gcp_vpc_site--properties--private_connectivity.md#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--gcp_vpc_site--properties--private_connectivity--cloud_link.md#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--gcp_vpc_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--gcp_vpc_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--namespace) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--gcp_vpc_site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--tenant) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--gcp_vpc_site--properties--private_connectivity--inside.md#section) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--gcp_vpc_site--properties--private_connectivity--outside.md#section) |
| `ssh_key` | [ssh_key](data-sources--gcp_vpc_site--reference.md#schema-ssh_key) |
| `sw` | [sw](data-sources--gcp_vpc_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--gcp_vpc_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--gcp_vpc_site--properties--sw.md#schema-sw--volterra_software_version) |
| `voltstack_cluster` | [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies.md#section) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#section) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--name) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--namespace) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--tenant) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](data-sources--gcp_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#section) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--name) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--namespace) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--tenant) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](data-sources--gcp_vpc_site--properties--voltstack_cluster--default_storage.md#section) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](data-sources--gcp_vpc_site--properties--voltstack_cluster--forward_proxy_allow_all.md#section) |
| `voltstack_cluster.gcp_certified_hw` | [voltstack_cluster.gcp_certified_hw](data-sources--gcp_vpc_site--properties--voltstack_cluster.md#schema-voltstack_cluster--gcp_certified_hw) |
| `voltstack_cluster.gcp_zone_names` | [voltstack_cluster.gcp_zone_names](data-sources--gcp_vpc_site--properties--voltstack_cluster.md#schema-voltstack_cluster--gcp_zone_names) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster--k8s_cluster.md#section) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--name) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--namespace) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--tenant) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](data-sources--gcp_vpc_site--properties--voltstack_cluster--no_dc_cluster_group.md#section) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](data-sources--gcp_vpc_site--properties--voltstack_cluster--no_forward_proxy.md#section) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--no_global_network.md#section) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster--no_k8s_cluster.md#section) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](data-sources--gcp_vpc_site--properties--voltstack_cluster--no_network_policy.md#section) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](data-sources--gcp_vpc_site--properties--voltstack_cluster--no_outside_static_routes.md#section) |
| `voltstack_cluster.node_number` | [voltstack_cluster.node_number](data-sources--gcp_vpc_site--properties--voltstack_cluster.md#schema-voltstack_cluster--node_number) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md#schema-voltstack_cluster--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster.site_local_network` | [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network.md#section) |
| `voltstack_cluster.site_local_network.existing_network` | [voltstack_cluster.site_local_network.existing_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--existing_network.md#section) |
| `voltstack_cluster.site_local_network.existing_network.name` | [voltstack_cluster.site_local_network.existing_network.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--existing_network.md#schema-voltstack_cluster--site_local_network--existing_network--name) |
| `voltstack_cluster.site_local_network.new_network` | [voltstack_cluster.site_local_network.new_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network.md#section) |
| `voltstack_cluster.site_local_network.new_network.name` | [voltstack_cluster.site_local_network.new_network.name](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network.md#schema-voltstack_cluster--site_local_network--new_network--name) |
| `voltstack_cluster.site_local_network.new_network_autogenerate` | [voltstack_cluster.site_local_network.new_network_autogenerate](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network_autogenerate.md#section) |
| `voltstack_cluster.site_local_subnet` | [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet.md#section) |
| `voltstack_cluster.site_local_subnet.existing_subnet` | [voltstack_cluster.site_local_subnet.existing_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--existing_subnet.md#section) |
| `voltstack_cluster.site_local_subnet.existing_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--existing_subnet.md#schema-voltstack_cluster--site_local_subnet--existing_subnet--subnet_name) |
| `voltstack_cluster.site_local_subnet.new_subnet` | [voltstack_cluster.site_local_subnet.new_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--new_subnet.md#section) |
| `voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4` | [voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--new_subnet.md#schema-voltstack_cluster--site_local_subnet--new_subnet--primary_ipv4) |
| `voltstack_cluster.site_local_subnet.new_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--new_subnet.md#schema-voltstack_cluster--site_local_subnet--new_subnet--subnet_name) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](data-sources--gcp_vpc_site--properties--voltstack_cluster--sm_connection_public_ip.md#section) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](data-sources--gcp_vpc_site--properties--voltstack_cluster--sm_connection_pvt_ip.md#section) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list.md#section) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#section) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name) |
| `waf_signatures` | [waf_signatures](data-sources--gcp_vpc_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--gcp_vpc_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--gcp_vpc_site--properties--waf_signatures--manual.md#section) |

## Next pages

- [admin_password](data-sources--gcp_vpc_site--properties--admin_password.md)
- [block_all_services](data-sources--gcp_vpc_site--properties--block_all_services.md)
- [blocked_services](data-sources--gcp_vpc_site--properties--blocked_services.md)
- [cloud_credentials](data-sources--gcp_vpc_site--properties--cloud_credentials.md)
- [coordinates](data-sources--gcp_vpc_site--properties--coordinates.md)
- [custom_dns](data-sources--gcp_vpc_site--properties--custom_dns.md)
- [default_blocked_services](data-sources--gcp_vpc_site--properties--default_blocked_services.md)
- [disable_encryption](data-sources--gcp_vpc_site--properties--disable_encryption.md)
- [enable_encryption](data-sources--gcp_vpc_site--properties--enable_encryption.md)
- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](data-sources--gcp_vpc_site--properties--log_receiver.md)
- [logs_streaming_disabled](data-sources--gcp_vpc_site--properties--logs_streaming_disabled.md)
- [offline_survivability_mode](data-sources--gcp_vpc_site--properties--offline_survivability_mode.md)
- [os](data-sources--gcp_vpc_site--properties--os.md)
- [private_connect_disabled](data-sources--gcp_vpc_site--properties--private_connect_disabled.md)
- [private_connectivity](data-sources--gcp_vpc_site--properties--private_connectivity.md)
- [sw](data-sources--gcp_vpc_site--properties--sw.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [waf_signatures](data-sources--gcp_vpc_site--properties--waf_signatures.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
