---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 230307, "body_sha256": "sha256:d779a6144fb51e42d68424682e352c8f58016e3dc52c137cb4a8b4c1810eadde", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:admin_password", "xcsh-docs:data-sources:azure_vnet_site:properties:azure_cred", "xcsh-docs:data-sources:azure_vnet_site:properties:block_all_services", "xcsh-docs:data-sources:azure_vnet_site:properties:blocked_services", "xcsh-docs:data-sources:azure_vnet_site:properties:coordinates", "xcsh-docs:data-sources:azure_vnet_site:properties:custom_dns", "xcsh-docs:data-sources:azure_vnet_site:properties:default_blocked_services", "xcsh-docs:data-sources:azure_vnet_site:properties:disable_encryption", "xcsh-docs:data-sources:azure_vnet_site:properties:enable_encryption", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar", "xcsh-docs:data-sources:azure_vnet_site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:azure_vnet_site:properties:log_receiver", "xcsh-docs:data-sources:azure_vnet_site:properties:logs_streaming_disabled", "xcsh-docs:data-sources:azure_vnet_site:properties:no_worker_nodes", "xcsh-docs:data-sources:azure_vnet_site:properties:offline_survivability_mode", "xcsh-docs:data-sources:azure_vnet_site:properties:os", "xcsh-docs:data-sources:azure_vnet_site:properties:sw", "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar", "xcsh-docs:data-sources:azure_vnet_site:properties:waf_signatures"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:reference", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:fundamentals", "path": "docs/guides/data-sources--azure_vnet_site--reference.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
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

- [admin_password](data-sources--azure_vnet_site--properties--admin_password.md): complete subsection reference.

<a id="schema-alternate_region"></a>

### alternate_region property

Type: `"string"`. Computed.

\[OneOf: alternate\_region, azure\_region\] Exclusive with \[azure\_region\] Name of the Azure
region which does not support availability zones.

Upstream description:

Exclusive with \[azure\_region\] Name of the Azure region which does not support availability zones.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

OneOf alternatives in this subsection:

- [alternate_region](data-sources--azure_vnet_site--reference.md#schema-alternate_region)
- [azure_region](data-sources--azure_vnet_site--reference.md#schema-azure_region)

Select alternatives according to the provider validators above.

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

- [azure_cred](data-sources--azure_vnet_site--properties--azure_cred.md): complete subsection reference.

<a id="schema-azure_region"></a>

### azure_region property

Type: `"string"`. Computed.

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Upstream description:

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [block_all_services](data-sources--azure_vnet_site--properties--block_all_services.md): complete subsection reference.

- [blocked_services](data-sources--azure_vnet_site--properties--blocked_services.md): complete subsection reference.

- [coordinates](data-sources--azure_vnet_site--properties--coordinates.md): complete subsection reference.

- [custom_dns](data-sources--azure_vnet_site--properties--custom_dns.md): complete subsection reference.

- [default_blocked_services](data-sources--azure_vnet_site--properties--default_blocked_services.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AzureVNETSite.

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

- [disable_encryption](data-sources--azure_vnet_site--properties--disable_encryption.md): complete subsection reference.

<a id="schema-disk_size"></a>

### disk_size property

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB. Server applies default when omitted.

Upstream description:

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

- [enable_encryption](data-sources--azure_vnet_site--properties--enable_encryption.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md): complete subsection reference.

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md): complete subsection reference.

- [ingress_gw](data-sources--azure_vnet_site--properties--ingress_gw.md): complete subsection reference.

- [ingress_gw_ar](data-sources--azure_vnet_site--properties--ingress_gw_ar.md): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](data-sources--azure_vnet_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](data-sources--azure_vnet_site--properties--logs_streaming_disabled.md): complete subsection reference.

<a id="schema-machine_type"></a>

### machine_type property

Type: `"string"`. Computed.

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Upstream description:

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AzureVNETSite.

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

Type: `"string"`. Optional, Computed.

Namespace where the AzureVNETSite exists.

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

- [no_worker_nodes](data-sources--azure_vnet_site--properties--no_worker_nodes.md): complete subsection reference.

<a id="schema-nodes_per_az"></a>

### nodes_per_az property

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

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

- [offline_survivability_mode](data-sources--azure_vnet_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](data-sources--azure_vnet_site--properties--os.md): complete subsection reference.

<a id="schema-resource_group"></a>

### resource_group property

Type: `"string"`. Computed.

Azure resource group for resources that will be created.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

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

- [sw](data-sources--azure_vnet_site--properties--sw.md): complete subsection reference.

<a id="schema-tags"></a>

### tags property

Type: `["map", "string"]`. Computed.

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console.

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

<a id="schema-total_nodes"></a>

### total_nodes property

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

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

- [vnet](data-sources--azure_vnet_site--properties--vnet.md): complete subsection reference.

- [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md): complete subsection reference.

- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md): complete subsection reference.

- [waf_signatures](data-sources--azure_vnet_site--properties--waf_signatures.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--azure_vnet_site--reference.md#schema-address) |
| `admin_password` | [admin_password](data-sources--azure_vnet_site--properties--admin_password.md#section) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](data-sources--azure_vnet_site--properties--admin_password--blindfold_secret_info.md#section) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](data-sources--azure_vnet_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--decryption_provider) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](data-sources--azure_vnet_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--location) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](data-sources--azure_vnet_site--properties--admin_password--blindfold_secret_info.md#schema-admin_password--blindfold_secret_info--store_provider) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](data-sources--azure_vnet_site--properties--admin_password--clear_secret_info.md#section) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](data-sources--azure_vnet_site--properties--admin_password--clear_secret_info.md#schema-admin_password--clear_secret_info--provider_ref) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](data-sources--azure_vnet_site--properties--admin_password--clear_secret_info.md#schema-admin_password--clear_secret_info--url) |
| `alternate_region` | [alternate_region](data-sources--azure_vnet_site--reference.md#schema-alternate_region) |
| `annotations` | [annotations](data-sources--azure_vnet_site--reference.md#schema-annotations) |
| `azure_cred` | [azure_cred](data-sources--azure_vnet_site--properties--azure_cred.md#section) |
| `azure_cred.name` | [azure_cred.name](data-sources--azure_vnet_site--properties--azure_cred.md#schema-azure_cred--name) |
| `azure_cred.namespace` | [azure_cred.namespace](data-sources--azure_vnet_site--properties--azure_cred.md#schema-azure_cred--namespace) |
| `azure_cred.tenant` | [azure_cred.tenant](data-sources--azure_vnet_site--properties--azure_cred.md#schema-azure_cred--tenant) |
| `azure_region` | [azure_region](data-sources--azure_vnet_site--reference.md#schema-azure_region) |
| `block_all_services` | [block_all_services](data-sources--azure_vnet_site--properties--block_all_services.md#section) |
| `blocked_services` | [blocked_services](data-sources--azure_vnet_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--azure_vnet_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--azure_vnet_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--azure_vnet_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--azure_vnet_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--azure_vnet_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `coordinates` | [coordinates](data-sources--azure_vnet_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--azure_vnet_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--azure_vnet_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](data-sources--azure_vnet_site--properties--custom_dns.md#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--azure_vnet_site--properties--custom_dns.md#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--azure_vnet_site--properties--custom_dns.md#schema-custom_dns--outside_nameserver) |
| `default_blocked_services` | [default_blocked_services](data-sources--azure_vnet_site--properties--default_blocked_services.md#section) |
| `description` | [description](data-sources--azure_vnet_site--reference.md#schema-description) |
| `disable_encryption` | [disable_encryption](data-sources--azure_vnet_site--properties--disable_encryption.md#section) |
| `disk_size` | [disk_size](data-sources--azure_vnet_site--reference.md#schema-disk_size) |
| `enable_encryption` | [enable_encryption](data-sources--azure_vnet_site--properties--enable_encryption.md#section) |
| `enable_encryption.disk_encryption_set_id` | [enable_encryption.disk_encryption_set_id](data-sources--azure_vnet_site--properties--enable_encryption.md#schema-enable_encryption--disk_encryption_set_id) |
| `enable_encryption.resource_group` | [enable_encryption.resource_group](data-sources--azure_vnet_site--properties--enable_encryption.md#schema-enable_encryption--resource_group) |
| `id` | [id](data-sources--azure_vnet_site--reference.md#schema-id) |
| `ingress_egress_gw` | [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md#section) |
| `ingress_egress_gw.accelerated_networking` | [ingress_egress_gw.accelerated_networking](data-sources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking.md#section) |
| `ingress_egress_gw.accelerated_networking.disable_spec` | [ingress_egress_gw.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking--disable_spec.md#section) |
| `ingress_egress_gw.accelerated_networking.enable` | [ingress_egress_gw.accelerated_networking.enable](data-sources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking--enable.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies.md#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies.md#section) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#section) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--name) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies--network_policies.md#schema-ingress_egress_gw--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md#section) |
| `ingress_egress_gw.az_nodes.azure_az` | [ingress_egress_gw.az_nodes.azure_az](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md#schema-ingress_egress_gw--az_nodes--azure_az) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet.md#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet` | [ingress_egress_gw.az_nodes.inside_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet.md#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_name` | [ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet.md#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet--subnet_name) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet.md#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.az_nodes.inside_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet_param.md#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet_param.md#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet.md#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet` | [ingress_egress_gw.az_nodes.outside_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet.md#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_name` | [ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet.md#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet--subnet_name) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet.md#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet_param.md#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--az_nodes--outside_subnet--subnet_param.md#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.azure_certified_hw` | [ingress_egress_gw.azure_certified_hw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md#schema-ingress_egress_gw--azure_certified_hw) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#section) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#section) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](data-sources--azure_vnet_site--properties--ingress_egress_gw--forward_proxy_allow_all.md#section) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.hub` | [ingress_egress_gw.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub.md#section) |
| `ingress_egress_gw.hub.express_route_disabled` | [ingress_egress_gw.hub.express_route_disabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_disabled.md#section) |
| `ingress_egress_gw.hub.express_route_enabled` | [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.advertise_to_route_server` | [ingress_egress_gw.hub.express_route_enabled.advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--advertise_to_route_server.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.auto_asn` | [ingress_egress_gw.hub.express_route_enabled.auto_asn](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--auto_asn.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections` | [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.circuit_id` | [ingress_egress_gw.hub.express_route_enabled.connections.circuit_id](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--circuit_id) |
| `ingress_egress_gw.hub.express_route_enabled.connections.metadata` | [ingress_egress_gw.hub.express_route_enabled.connections.metadata](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--metadata.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.metadata.description_spec` | [ingress_egress_gw.hub.express_route_enabled.connections.metadata.description_spec](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--metadata.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--metadata--description_spec) |
| `ingress_egress_gw.hub.express_route_enabled.connections.metadata.name` | [ingress_egress_gw.hub.express_route_enabled.connections.metadata.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--metadata.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--metadata--name) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--decryption_provider) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--store_provider) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--provider_ref) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--url) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.circuit_id` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.circuit_id](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections--other_subscription.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--circuit_id) |
| `ingress_egress_gw.hub.express_route_enabled.connections.weight` | [ingress_egress_gw.hub.express_route_enabled.connections.weight](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--connections.md#schema-ingress_egress_gw--hub--express_route_enabled--connections--weight) |
| `ingress_egress_gw.hub.express_route_enabled.custom_asn` | [ingress_egress_gw.hub.express_route_enabled.custom_asn](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md#schema-ingress_egress_gw--hub--express_route_enabled--custom_asn) |
| `ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server` | [ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--do_not_advertise_to_route_server.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--auto.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet.md#schema-ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet_param.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet_param.md#schema-ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet--auto.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet.md#schema-ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet_param.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet_param.md#schema-ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route` | [ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--site_registration_over_express_route.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name` | [ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--site_registration_over_express_route.md#schema-ingress_egress_gw--hub--express_route_enabled--site_registration_over_express_route--cloudlink_network_name) |
| `ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet` | [ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--site_registration_over_internet.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_ergw1az` | [ingress_egress_gw.hub.express_route_enabled.sku_ergw1az](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--sku_ergw1az.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_ergw2az` | [ingress_egress_gw.hub.express_route_enabled.sku_ergw2az](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--sku_ergw2az.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_high_perf` | [ingress_egress_gw.hub.express_route_enabled.sku_high_perf](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--sku_high_perf.md#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_standard` | [ingress_egress_gw.hub.express_route_enabled.sku_standard](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--sku_standard.md#section) |
| `ingress_egress_gw.hub.spoke_vnets` | [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.auto` | [ingress_egress_gw.hub.spoke_vnets.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--auto.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.labels` | [ingress_egress_gw.hub.spoke_vnets.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--labels.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.manual` | [ingress_egress_gw.hub.spoke_vnets.manual](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--manual.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet` | [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing` | [ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet--f5_orchestrated_routing.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing` | [ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet--manual_routing.md#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.resource_group` | [ingress_egress_gw.hub.spoke_vnets.vnet.resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet.md#schema-ingress_egress_gw--hub--spoke_vnets--vnet--resource_group) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.vnet_name` | [ingress_egress_gw.hub.spoke_vnets.vnet.vnet_name](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet.md#schema-ingress_egress_gw--hub--spoke_vnets--vnet--vnet_name) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md#schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](data-sources--azure_vnet_site--properties--ingress_egress_gw--no_dc_cluster_group.md#section) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](data-sources--azure_vnet_site--properties--ingress_egress_gw--no_forward_proxy.md#section) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](data-sources--azure_vnet_site--properties--ingress_egress_gw--no_global_network.md#section) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw--no_inside_static_routes.md#section) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](data-sources--azure_vnet_site--properties--ingress_egress_gw--no_network_policy.md#section) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw--no_outside_static_routes.md#section) |
| `ingress_egress_gw.not_hub` | [ingress_egress_gw.not_hub](data-sources--azure_vnet_site--properties--ingress_egress_gw--not_hub.md#section) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes--static_route_list.md#schema-ingress_egress_gw--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](data-sources--azure_vnet_site--properties--ingress_egress_gw--sm_connection_public_ip.md#section) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](data-sources--azure_vnet_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md#section) |
| `ingress_egress_gw_ar` | [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md#section) |
| `ingress_egress_gw_ar.accelerated_networking` | [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--accelerated_networking.md#section) |
| `ingress_egress_gw_ar.accelerated_networking.disable_spec` | [ingress_egress_gw_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--accelerated_networking--disable_spec.md#section) |
| `ingress_egress_gw_ar.accelerated_networking.enable` | [ingress_egress_gw_ar.accelerated_networking.enable](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--accelerated_networking--enable.md#section) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies` | [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_enhanced_firewall_policies.md#section) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw_ar.active_forward_proxy_policies` | [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_forward_proxy_policies.md#section) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies.md#schema-ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw_ar.active_network_policies` | [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_network_policies.md#section) |
| `ingress_egress_gw_ar.active_network_policies.network_policies` | [ingress_egress_gw_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_network_policies--network_policies.md#section) |
| `ingress_egress_gw_ar.active_network_policies.network_policies.name` | [ingress_egress_gw_ar.active_network_policies.network_policies.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_network_policies--network_policies.md#schema-ingress_egress_gw_ar--active_network_policies--network_policies--name) |
| `ingress_egress_gw_ar.active_network_policies.network_policies.namespace` | [ingress_egress_gw_ar.active_network_policies.network_policies.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_network_policies--network_policies.md#schema-ingress_egress_gw_ar--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw_ar.active_network_policies.network_policies.tenant` | [ingress_egress_gw_ar.active_network_policies.network_policies.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--active_network_policies--network_policies.md#schema-ingress_egress_gw_ar--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw_ar.azure_certified_hw` | [ingress_egress_gw_ar.azure_certified_hw](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md#schema-ingress_egress_gw_ar--azure_certified_hw) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_inside_vn.md#section) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn.name` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw_ar--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw_ar--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_inside_vn.md#schema-ingress_egress_gw_ar--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_outside_vn.md#section) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn.name` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw_ar--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw_ar--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--dc_cluster_group_outside_vn.md#schema-ingress_egress_gw_ar--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw_ar.forward_proxy_allow_all` | [ingress_egress_gw_ar.forward_proxy_allow_all](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--forward_proxy_allow_all.md#section) |
| `ingress_egress_gw_ar.global_network_list` | [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list.md#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections` | [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections.md#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw_ar.hub` | [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md#section) |
| `ingress_egress_gw_ar.hub.express_route_disabled` | [ingress_egress_gw_ar.hub.express_route_disabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_disabled.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled` | [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server` | [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--advertise_to_route_server.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.auto_asn` | [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--auto_asn.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections` | [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.circuit_id` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.circuit_id](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--circuit_id) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.description_spec` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.description_spec](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata--description_spec) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.name` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata--name) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--decryption_provider) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--store_provider) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--provider_ref) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--url) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.circuit_id` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.circuit_id](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--circuit_id) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.weight` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.weight](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--weight) |
| `ingress_egress_gw_ar.hub.express_route_enabled.custom_asn` | [ingress_egress_gw_ar.hub.express_route_enabled.custom_asn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--custom_asn) |
| `ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server` | [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--do_not_advertise_to_route_server.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--auto.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet_param.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet_param.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--auto.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet_param.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet_param.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route` | [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_express_route.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name` | [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_express_route.md#schema-ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_express_route--cloudlink_network_name) |
| `ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet` | [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_internet.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_ergw1az.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_ergw2az.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_high_perf.md#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_standard` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_standard.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets` | [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.auto` | [ingress_egress_gw_ar.hub.spoke_vnets.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--auto.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.labels` | [ingress_egress_gw_ar.hub.spoke_vnets.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--labels.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.manual` | [ingress_egress_gw_ar.hub.spoke_vnets.manual](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--manual.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet--f5_orchestrated_routing.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet--manual_routing.md#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.resource_group` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet.md#schema-ingress_egress_gw_ar--hub--spoke_vnets--vnet--resource_group) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.vnet_name` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.vnet_name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--vnet.md#schema-ingress_egress_gw_ar--hub--spoke_vnets--vnet--vnet_name) |
| `ingress_egress_gw_ar.inside_static_routes` | [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list` | [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.simple_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list.md#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw_ar.no_dc_cluster_group` | [ingress_egress_gw_ar.no_dc_cluster_group](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--no_dc_cluster_group.md#section) |
| `ingress_egress_gw_ar.no_forward_proxy` | [ingress_egress_gw_ar.no_forward_proxy](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--no_forward_proxy.md#section) |
| `ingress_egress_gw_ar.no_global_network` | [ingress_egress_gw_ar.no_global_network](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--no_global_network.md#section) |
| `ingress_egress_gw_ar.no_inside_static_routes` | [ingress_egress_gw_ar.no_inside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--no_inside_static_routes.md#section) |
| `ingress_egress_gw_ar.no_network_policy` | [ingress_egress_gw_ar.no_network_policy](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--no_network_policy.md#section) |
| `ingress_egress_gw_ar.no_outside_static_routes` | [ingress_egress_gw_ar.no_outside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--no_outside_static_routes.md#section) |
| `ingress_egress_gw_ar.node` | [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node.md#section) |
| `ingress_egress_gw_ar.node.fault_domain` | [ingress_egress_gw_ar.node.fault_domain](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node.md#schema-ingress_egress_gw_ar--node--fault_domain) |
| `ingress_egress_gw_ar.node.inside_subnet` | [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet.md#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet` | [ingress_egress_gw_ar.node.inside_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet--subnet.md#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_name` | [ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet--subnet.md#schema-ingress_egress_gw_ar--node--inside_subnet--subnet--subnet_name) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet--subnet.md#schema-ingress_egress_gw_ar--node--inside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet_param` | [ingress_egress_gw_ar.node.inside_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet--subnet_param.md#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.node.inside_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--inside_subnet--subnet_param.md#schema-ingress_egress_gw_ar--node--inside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.node.node_number` | [ingress_egress_gw_ar.node.node_number](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node.md#schema-ingress_egress_gw_ar--node--node_number) |
| `ingress_egress_gw_ar.node.outside_subnet` | [ingress_egress_gw_ar.node.outside_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet.md#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet` | [ingress_egress_gw_ar.node.outside_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet--subnet.md#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_name` | [ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet--subnet.md#schema-ingress_egress_gw_ar--node--outside_subnet--subnet--subnet_name) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet--subnet.md#schema-ingress_egress_gw_ar--node--outside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet_param` | [ingress_egress_gw_ar.node.outside_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet--subnet_param.md#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.node.outside_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node--outside_subnet--subnet_param.md#schema-ingress_egress_gw_ar--node--outside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.node.update_domain` | [ingress_egress_gw_ar.node.update_domain](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--node.md#schema-ingress_egress_gw_ar--node--update_domain) |
| `ingress_egress_gw_ar.not_hub` | [ingress_egress_gw_ar.not_hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--not_hub.md#section) |
| `ingress_egress_gw_ar.outside_static_routes` | [ingress_egress_gw_ar.outside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list` | [ingress_egress_gw_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.simple_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list.md#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw_ar.performance_enhancement_mode` | [ingress_egress_gw_ar.performance_enhancement_mode](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode.md#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `ingress_egress_gw_ar.sm_connection_public_ip` | [ingress_egress_gw_ar.sm_connection_public_ip](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--sm_connection_public_ip.md#section) |
| `ingress_egress_gw_ar.sm_connection_pvt_ip` | [ingress_egress_gw_ar.sm_connection_pvt_ip](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--sm_connection_pvt_ip.md#section) |
| `ingress_gw` | [ingress_gw](data-sources--azure_vnet_site--properties--ingress_gw.md#section) |
| `ingress_gw.accelerated_networking` | [ingress_gw.accelerated_networking](data-sources--azure_vnet_site--properties--ingress_gw--accelerated_networking.md#section) |
| `ingress_gw.accelerated_networking.disable_spec` | [ingress_gw.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--ingress_gw--accelerated_networking--disable_spec.md#section) |
| `ingress_gw.accelerated_networking.enable` | [ingress_gw.accelerated_networking.enable](data-sources--azure_vnet_site--properties--ingress_gw--accelerated_networking--enable.md#section) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes.md#section) |
| `ingress_gw.az_nodes.azure_az` | [ingress_gw.az_nodes.azure_az](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes.md#schema-ingress_gw--az_nodes--azure_az) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet.md#section) |
| `ingress_gw.az_nodes.local_subnet.subnet` | [ingress_gw.az_nodes.local_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet--subnet.md#section) |
| `ingress_gw.az_nodes.local_subnet.subnet.subnet_name` | [ingress_gw.az_nodes.local_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet--subnet.md#schema-ingress_gw--az_nodes--local_subnet--subnet--subnet_name) |
| `ingress_gw.az_nodes.local_subnet.subnet.subnet_resource_grp` | [ingress_gw.az_nodes.local_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet--subnet.md#schema-ingress_gw--az_nodes--local_subnet--subnet--subnet_resource_grp) |
| `ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group` | [ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet--subnet_param.md#section) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_gw--az_nodes--local_subnet--subnet_param.md#schema-ingress_gw--az_nodes--local_subnet--subnet_param--ipv4) |
| `ingress_gw.azure_certified_hw` | [ingress_gw.azure_certified_hw](data-sources--azure_vnet_site--properties--ingress_gw.md#schema-ingress_gw--azure_certified_hw) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `ingress_gw_ar` | [ingress_gw_ar](data-sources--azure_vnet_site--properties--ingress_gw_ar.md#section) |
| `ingress_gw_ar.accelerated_networking` | [ingress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking.md#section) |
| `ingress_gw_ar.accelerated_networking.disable_spec` | [ingress_gw_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking--disable_spec.md#section) |
| `ingress_gw_ar.accelerated_networking.enable` | [ingress_gw_ar.accelerated_networking.enable](data-sources--azure_vnet_site--properties--ingress_gw_ar--accelerated_networking--enable.md#section) |
| `ingress_gw_ar.azure_certified_hw` | [ingress_gw_ar.azure_certified_hw](data-sources--azure_vnet_site--properties--ingress_gw_ar.md#schema-ingress_gw_ar--azure_certified_hw) |
| `ingress_gw_ar.node` | [ingress_gw_ar.node](data-sources--azure_vnet_site--properties--ingress_gw_ar--node.md#section) |
| `ingress_gw_ar.node.fault_domain` | [ingress_gw_ar.node.fault_domain](data-sources--azure_vnet_site--properties--ingress_gw_ar--node.md#schema-ingress_gw_ar--node--fault_domain) |
| `ingress_gw_ar.node.local_subnet` | [ingress_gw_ar.node.local_subnet](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet.md#section) |
| `ingress_gw_ar.node.local_subnet.subnet` | [ingress_gw_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet.md#section) |
| `ingress_gw_ar.node.local_subnet.subnet.subnet_name` | [ingress_gw_ar.node.local_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet.md#schema-ingress_gw_ar--node--local_subnet--subnet--subnet_name) |
| `ingress_gw_ar.node.local_subnet.subnet.subnet_resource_grp` | [ingress_gw_ar.node.local_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet.md#schema-ingress_gw_ar--node--local_subnet--subnet--subnet_resource_grp) |
| `ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group` | [ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet--vnet_resource_group.md#section) |
| `ingress_gw_ar.node.local_subnet.subnet_param` | [ingress_gw_ar.node.local_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet_param.md#section) |
| `ingress_gw_ar.node.local_subnet.subnet_param.ipv4` | [ingress_gw_ar.node.local_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--ingress_gw_ar--node--local_subnet--subnet_param.md#schema-ingress_gw_ar--node--local_subnet--subnet_param--ipv4) |
| `ingress_gw_ar.node.node_number` | [ingress_gw_ar.node.node_number](data-sources--azure_vnet_site--properties--ingress_gw_ar--node.md#schema-ingress_gw_ar--node--node_number) |
| `ingress_gw_ar.node.update_domain` | [ingress_gw_ar.node.update_domain](data-sources--azure_vnet_site--properties--ingress_gw_ar--node.md#schema-ingress_gw_ar--node--update_domain) |
| `ingress_gw_ar.performance_enhancement_mode` | [ingress_gw_ar.performance_enhancement_mode](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode.md#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](data-sources--azure_vnet_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](data-sources--azure_vnet_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](data-sources--azure_vnet_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--azure_vnet_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--azure_vnet_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--azure_vnet_site--properties--logs_streaming_disabled.md#section) |
| `machine_type` | [machine_type](data-sources--azure_vnet_site--reference.md#schema-machine_type) |
| `name` | [name](data-sources--azure_vnet_site--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--azure_vnet_site--reference.md#schema-namespace) |
| `no_worker_nodes` | [no_worker_nodes](data-sources--azure_vnet_site--properties--no_worker_nodes.md#section) |
| `nodes_per_az` | [nodes_per_az](data-sources--azure_vnet_site--reference.md#schema-nodes_per_az) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--azure_vnet_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--azure_vnet_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--azure_vnet_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](data-sources--azure_vnet_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](data-sources--azure_vnet_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--azure_vnet_site--properties--os.md#schema-os--operating_system_version) |
| `resource_group` | [resource_group](data-sources--azure_vnet_site--reference.md#schema-resource_group) |
| `ssh_key` | [ssh_key](data-sources--azure_vnet_site--reference.md#schema-ssh_key) |
| `sw` | [sw](data-sources--azure_vnet_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--azure_vnet_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--azure_vnet_site--properties--sw.md#schema-sw--volterra_software_version) |
| `tags` | [tags](data-sources--azure_vnet_site--reference.md#schema-tags) |
| `total_nodes` | [total_nodes](data-sources--azure_vnet_site--reference.md#schema-total_nodes) |
| `vnet` | [vnet](data-sources--azure_vnet_site--properties--vnet.md#section) |
| `vnet.existing_vnet` | [vnet.existing_vnet](data-sources--azure_vnet_site--properties--vnet--existing_vnet.md#section) |
| `vnet.existing_vnet.f5_orchestrated_routing` | [vnet.existing_vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--properties--vnet--existing_vnet--f5_orchestrated_routing.md#section) |
| `vnet.existing_vnet.manual_routing` | [vnet.existing_vnet.manual_routing](data-sources--azure_vnet_site--properties--vnet--existing_vnet--manual_routing.md#section) |
| `vnet.existing_vnet.resource_group` | [vnet.existing_vnet.resource_group](data-sources--azure_vnet_site--properties--vnet--existing_vnet.md#schema-vnet--existing_vnet--resource_group) |
| `vnet.existing_vnet.vnet_name` | [vnet.existing_vnet.vnet_name](data-sources--azure_vnet_site--properties--vnet--existing_vnet.md#schema-vnet--existing_vnet--vnet_name) |
| `vnet.new_vnet` | [vnet.new_vnet](data-sources--azure_vnet_site--properties--vnet--new_vnet.md#section) |
| `vnet.new_vnet.autogenerate` | [vnet.new_vnet.autogenerate](data-sources--azure_vnet_site--properties--vnet--new_vnet--autogenerate.md#section) |
| `vnet.new_vnet.name` | [vnet.new_vnet.name](data-sources--azure_vnet_site--properties--vnet--new_vnet.md#schema-vnet--new_vnet--name) |
| `vnet.new_vnet.primary_ipv4` | [vnet.new_vnet.primary_ipv4](data-sources--azure_vnet_site--properties--vnet--new_vnet.md#schema-vnet--new_vnet--primary_ipv4) |
| `voltstack_cluster` | [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md#section) |
| `voltstack_cluster.accelerated_networking` | [voltstack_cluster.accelerated_networking](data-sources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking.md#section) |
| `voltstack_cluster.accelerated_networking.disable_spec` | [voltstack_cluster.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking--disable_spec.md#section) |
| `voltstack_cluster.accelerated_networking.enable` | [voltstack_cluster.accelerated_networking.enable](data-sources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking--enable.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_forward_proxy_policies.md#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--azure_vnet_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_network_policies.md#section) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](data-sources--azure_vnet_site--properties--voltstack_cluster--active_network_policies--network_policies.md#section) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](data-sources--azure_vnet_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--name) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--namespace) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--active_network_policies--network_policies.md#schema-voltstack_cluster--active_network_policies--network_policies--tenant) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes.md#section) |
| `voltstack_cluster.az_nodes.azure_az` | [voltstack_cluster.az_nodes.azure_az](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes.md#schema-voltstack_cluster--az_nodes--azure_az) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet.md#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet` | [voltstack_cluster.az_nodes.local_subnet.subnet](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet.md#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet.subnet_name` | [voltstack_cluster.az_nodes.local_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet.md#schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_name) |
| `voltstack_cluster.az_nodes.local_subnet.subnet.subnet_resource_grp` | [voltstack_cluster.az_nodes.local_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet.md#schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_resource_grp) |
| `voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group` | [voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet--vnet_resource_group.md#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md#schema-voltstack_cluster--az_nodes--local_subnet--subnet_param--ipv4) |
| `voltstack_cluster.azure_certified_hw` | [voltstack_cluster.azure_certified_hw](data-sources--azure_vnet_site--properties--voltstack_cluster.md#schema-voltstack_cluster--azure_certified_hw) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](data-sources--azure_vnet_site--properties--voltstack_cluster--dc_cluster_group.md#section) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](data-sources--azure_vnet_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--name) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--namespace) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--dc_cluster_group.md#schema-voltstack_cluster--dc_cluster_group--tenant) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](data-sources--azure_vnet_site--properties--voltstack_cluster--default_storage.md#section) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](data-sources--azure_vnet_site--properties--voltstack_cluster--forward_proxy_allow_all.md#section) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster--k8s_cluster.md#section) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](data-sources--azure_vnet_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--name) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--namespace) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--k8s_cluster.md#schema-voltstack_cluster--k8s_cluster--tenant) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](data-sources--azure_vnet_site--properties--voltstack_cluster--no_dc_cluster_group.md#section) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](data-sources--azure_vnet_site--properties--voltstack_cluster--no_forward_proxy.md#section) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](data-sources--azure_vnet_site--properties--voltstack_cluster--no_global_network.md#section) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster--no_k8s_cluster.md#section) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](data-sources--azure_vnet_site--properties--voltstack_cluster--no_network_policy.md#section) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](data-sources--azure_vnet_site--properties--voltstack_cluster--no_outside_static_routes.md#section) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md#schema-voltstack_cluster--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](data-sources--azure_vnet_site--properties--voltstack_cluster--sm_connection_public_ip.md#section) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](data-sources--azure_vnet_site--properties--voltstack_cluster--sm_connection_pvt_ip.md#section) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](data-sources--azure_vnet_site--properties--voltstack_cluster--storage_class_list.md#section) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](data-sources--azure_vnet_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#section) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](data-sources--azure_vnet_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](data-sources--azure_vnet_site--properties--voltstack_cluster--storage_class_list--storage_classes.md#schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name) |
| `voltstack_cluster_ar` | [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md#section) |
| `voltstack_cluster_ar.accelerated_networking` | [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--accelerated_networking.md#section) |
| `voltstack_cluster_ar.accelerated_networking.disable_spec` | [voltstack_cluster_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--accelerated_networking--disable_spec.md#section) |
| `voltstack_cluster_ar.accelerated_networking.enable` | [voltstack_cluster_ar.accelerated_networking.enable](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--accelerated_networking--enable.md#section) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies` | [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_enhanced_firewall_policies.md#section) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster_ar.active_forward_proxy_policies` | [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies.md#section) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies.md#schema-voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster_ar.active_network_policies` | [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies.md#section) |
| `voltstack_cluster_ar.active_network_policies.network_policies` | [voltstack_cluster_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md#section) |
| `voltstack_cluster_ar.active_network_policies.network_policies.name` | [voltstack_cluster_ar.active_network_policies.network_policies.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md#schema-voltstack_cluster_ar--active_network_policies--network_policies--name) |
| `voltstack_cluster_ar.active_network_policies.network_policies.namespace` | [voltstack_cluster_ar.active_network_policies.network_policies.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md#schema-voltstack_cluster_ar--active_network_policies--network_policies--namespace) |
| `voltstack_cluster_ar.active_network_policies.network_policies.tenant` | [voltstack_cluster_ar.active_network_policies.network_policies.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--active_network_policies--network_policies.md#schema-voltstack_cluster_ar--active_network_policies--network_policies--tenant) |
| `voltstack_cluster_ar.azure_certified_hw` | [voltstack_cluster_ar.azure_certified_hw](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md#schema-voltstack_cluster_ar--azure_certified_hw) |
| `voltstack_cluster_ar.dc_cluster_group` | [voltstack_cluster_ar.dc_cluster_group](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--dc_cluster_group.md#section) |
| `voltstack_cluster_ar.dc_cluster_group.name` | [voltstack_cluster_ar.dc_cluster_group.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--dc_cluster_group.md#schema-voltstack_cluster_ar--dc_cluster_group--name) |
| `voltstack_cluster_ar.dc_cluster_group.namespace` | [voltstack_cluster_ar.dc_cluster_group.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--dc_cluster_group.md#schema-voltstack_cluster_ar--dc_cluster_group--namespace) |
| `voltstack_cluster_ar.dc_cluster_group.tenant` | [voltstack_cluster_ar.dc_cluster_group.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--dc_cluster_group.md#schema-voltstack_cluster_ar--dc_cluster_group--tenant) |
| `voltstack_cluster_ar.default_storage` | [voltstack_cluster_ar.default_storage](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--default_storage.md#section) |
| `voltstack_cluster_ar.forward_proxy_allow_all` | [voltstack_cluster_ar.forward_proxy_allow_all](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--forward_proxy_allow_all.md#section) |
| `voltstack_cluster_ar.global_network_list` | [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list.md#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections` | [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections.md#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster_ar.k8s_cluster` | [voltstack_cluster_ar.k8s_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--k8s_cluster.md#section) |
| `voltstack_cluster_ar.k8s_cluster.name` | [voltstack_cluster_ar.k8s_cluster.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--k8s_cluster.md#schema-voltstack_cluster_ar--k8s_cluster--name) |
| `voltstack_cluster_ar.k8s_cluster.namespace` | [voltstack_cluster_ar.k8s_cluster.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--k8s_cluster.md#schema-voltstack_cluster_ar--k8s_cluster--namespace) |
| `voltstack_cluster_ar.k8s_cluster.tenant` | [voltstack_cluster_ar.k8s_cluster.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--k8s_cluster.md#schema-voltstack_cluster_ar--k8s_cluster--tenant) |
| `voltstack_cluster_ar.no_dc_cluster_group` | [voltstack_cluster_ar.no_dc_cluster_group](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--no_dc_cluster_group.md#section) |
| `voltstack_cluster_ar.no_forward_proxy` | [voltstack_cluster_ar.no_forward_proxy](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--no_forward_proxy.md#section) |
| `voltstack_cluster_ar.no_global_network` | [voltstack_cluster_ar.no_global_network](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--no_global_network.md#section) |
| `voltstack_cluster_ar.no_k8s_cluster` | [voltstack_cluster_ar.no_k8s_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--no_k8s_cluster.md#section) |
| `voltstack_cluster_ar.no_network_policy` | [voltstack_cluster_ar.no_network_policy](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--no_network_policy.md#section) |
| `voltstack_cluster_ar.no_outside_static_routes` | [voltstack_cluster_ar.no_outside_static_routes](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--no_outside_static_routes.md#section) |
| `voltstack_cluster_ar.node` | [voltstack_cluster_ar.node](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node.md#section) |
| `voltstack_cluster_ar.node.fault_domain` | [voltstack_cluster_ar.node.fault_domain](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node.md#schema-voltstack_cluster_ar--node--fault_domain) |
| `voltstack_cluster_ar.node.local_subnet` | [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet.md#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet` | [voltstack_cluster_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet.md#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet.subnet_name` | [voltstack_cluster_ar.node.local_subnet.subnet.subnet_name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet.md#schema-voltstack_cluster_ar--node--local_subnet--subnet--subnet_name) |
| `voltstack_cluster_ar.node.local_subnet.subnet.subnet_resource_grp` | [voltstack_cluster_ar.node.local_subnet.subnet.subnet_resource_grp](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet.md#schema-voltstack_cluster_ar--node--local_subnet--subnet--subnet_resource_grp) |
| `voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group` | [voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet--vnet_resource_group.md#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet_param` | [voltstack_cluster_ar.node.local_subnet.subnet_param](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet_param.md#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet_param.ipv4` | [voltstack_cluster_ar.node.local_subnet.subnet_param.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node--local_subnet--subnet_param.md#schema-voltstack_cluster_ar--node--local_subnet--subnet_param--ipv4) |
| `voltstack_cluster_ar.node.node_number` | [voltstack_cluster_ar.node.node_number](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node.md#schema-voltstack_cluster_ar--node--node_number) |
| `voltstack_cluster_ar.node.update_domain` | [voltstack_cluster_ar.node.update_domain](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--node.md#schema-voltstack_cluster_ar--node--update_domain) |
| `voltstack_cluster_ar.outside_static_routes` | [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list` | [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--labels.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster_ar.outside_static_routes.static_route_list.simple_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list.md#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster_ar.sm_connection_public_ip` | [voltstack_cluster_ar.sm_connection_public_ip](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--sm_connection_public_ip.md#section) |
| `voltstack_cluster_ar.sm_connection_pvt_ip` | [voltstack_cluster_ar.sm_connection_pvt_ip](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--sm_connection_pvt_ip.md#section) |
| `voltstack_cluster_ar.storage_class_list` | [voltstack_cluster_ar.storage_class_list](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--storage_class_list.md#section) |
| `voltstack_cluster_ar.storage_class_list.storage_classes` | [voltstack_cluster_ar.storage_class_list.storage_classes](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--storage_class_list--storage_classes.md#section) |
| `voltstack_cluster_ar.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster_ar.storage_class_list.storage_classes.default_storage_class](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--storage_class_list--storage_classes.md#schema-voltstack_cluster_ar--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster_ar.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster_ar.storage_class_list.storage_classes.storage_class_name](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--storage_class_list--storage_classes.md#schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name) |
| `waf_signatures` | [waf_signatures](data-sources--azure_vnet_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--azure_vnet_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--azure_vnet_site--properties--waf_signatures--manual.md#section) |

## Next pages

- [admin_password](data-sources--azure_vnet_site--properties--admin_password.md)
- [azure_cred](data-sources--azure_vnet_site--properties--azure_cred.md)
- [block_all_services](data-sources--azure_vnet_site--properties--block_all_services.md)
- [blocked_services](data-sources--azure_vnet_site--properties--blocked_services.md)
- [coordinates](data-sources--azure_vnet_site--properties--coordinates.md)
- [custom_dns](data-sources--azure_vnet_site--properties--custom_dns.md)
- [default_blocked_services](data-sources--azure_vnet_site--properties--default_blocked_services.md)
- [disable_encryption](data-sources--azure_vnet_site--properties--disable_encryption.md)
- [enable_encryption](data-sources--azure_vnet_site--properties--enable_encryption.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_gw](data-sources--azure_vnet_site--properties--ingress_gw.md)
- [ingress_gw_ar](data-sources--azure_vnet_site--properties--ingress_gw_ar.md)
- [kubernetes_upgrade_drain](data-sources--azure_vnet_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](data-sources--azure_vnet_site--properties--log_receiver.md)
- [logs_streaming_disabled](data-sources--azure_vnet_site--properties--logs_streaming_disabled.md)
- [no_worker_nodes](data-sources--azure_vnet_site--properties--no_worker_nodes.md)
- [offline_survivability_mode](data-sources--azure_vnet_site--properties--offline_survivability_mode.md)
- [os](data-sources--azure_vnet_site--properties--os.md)
- [sw](data-sources--azure_vnet_site--properties--sw.md)
- [vnet](data-sources--azure_vnet_site--properties--vnet.md)
- [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [waf_signatures](data-sources--azure_vnet_site--properties--waf_signatures.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
