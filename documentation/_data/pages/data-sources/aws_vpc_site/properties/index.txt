---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 139806, "body_sha256": "sha256:fb2c89244676fdfe9f65779530c24b350c1bbb751d22863fc2b95a9517f2b3bf", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:admin_password", "xcsh-docs:data-sources:aws_vpc_site:properties:aws_cred", "xcsh-docs:data-sources:aws_vpc_site:properties:block_all_services", "xcsh-docs:data-sources:aws_vpc_site:properties:blocked_services", "xcsh-docs:data-sources:aws_vpc_site:properties:coordinates", "xcsh-docs:data-sources:aws_vpc_site:properties:custom_dns", "xcsh-docs:data-sources:aws_vpc_site:properties:custom_security_group", "xcsh-docs:data-sources:aws_vpc_site:properties:default_blocked_services", "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_disabled", "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled", "xcsh-docs:data-sources:aws_vpc_site:properties:disable_encryption", "xcsh-docs:data-sources:aws_vpc_site:properties:disable_internet_vip", "xcsh-docs:data-sources:aws_vpc_site:properties:egress_gateway_default", "xcsh-docs:data-sources:aws_vpc_site:properties:egress_nat_gw", "xcsh-docs:data-sources:aws_vpc_site:properties:egress_virtual_private_gateway", "xcsh-docs:data-sources:aws_vpc_site:properties:enable_encryption", "xcsh-docs:data-sources:aws_vpc_site:properties:enable_internet_vip", "xcsh-docs:data-sources:aws_vpc_site:properties:f5_orchestrated_routing", "xcsh-docs:data-sources:aws_vpc_site:properties:f5xc_security_group", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_gw", "xcsh-docs:data-sources:aws_vpc_site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:aws_vpc_site:properties:log_receiver", "xcsh-docs:data-sources:aws_vpc_site:properties:logs_streaming_disabled", "xcsh-docs:data-sources:aws_vpc_site:properties:manual_routing", "xcsh-docs:data-sources:aws_vpc_site:properties:no_worker_nodes", "xcsh-docs:data-sources:aws_vpc_site:properties:offline_survivability_mode", "xcsh-docs:data-sources:aws_vpc_site:properties:os", "xcsh-docs:data-sources:aws_vpc_site:properties:private_connectivity", "xcsh-docs:data-sources:aws_vpc_site:properties:sw", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "xcsh-docs:data-sources:aws_vpc_site:properties:vpc", "xcsh-docs:data-sources:aws_vpc_site:properties:waf_signatures"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:reference", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:fundamentals", "path": "documentation/data-sources/aws_vpc_site/properties/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
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

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/): complete subsection reference.

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

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/aws_cred/): complete subsection reference.

<a id="schema-aws_region"></a>

### aws_region property

Type: `"string"`. Computed.

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

- [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/block_all_services/): complete subsection reference.

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/): complete subsection reference.

- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/coordinates/): complete subsection reference.

- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_dns/): complete subsection reference.

- [custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_security_group/): complete subsection reference.

- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/default_blocked_services/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AWSVPCSite.

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

- [direct_connect_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_disabled/): complete subsection reference.

- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/): complete subsection reference.

- [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/disable_encryption/): complete subsection reference.

- [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/disable_internet_vip/): complete subsection reference.

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

- [egress_gateway_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_gateway_default/): complete subsection reference.

- [egress_nat_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_nat_gw/): complete subsection reference.

- [egress_virtual_private_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_virtual_private_gateway/): complete subsection reference.

- [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_encryption/): complete subsection reference.

- [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_internet_vip/): complete subsection reference.

- [f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/f5_orchestrated_routing/): complete subsection reference.

- [f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/f5xc_security_group/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/): complete subsection reference.

- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/): complete subsection reference.

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

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/): complete subsection reference.

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

- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/log_receiver/): complete subsection reference.

- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/logs_streaming_disabled/): complete subsection reference.

- [manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/manual_routing/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AWSVPCSite.

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

Namespace where the AWSVPCSite exists.

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

- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/no_worker_nodes/): complete subsection reference.

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

- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/offline_survivability_mode/): complete subsection reference.

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/os/): complete subsection reference.

- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/): complete subsection reference.

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

- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/sw/): complete subsection reference.

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

- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/): complete subsection reference.

- [vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/): complete subsection reference.

- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/waf_signatures/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-address) |
| `admin_password` | [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/#section) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/blindfold_secret_info/#section) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/blindfold_secret_info/#schema-admin_password--blindfold_secret_info--decryption_provider) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/blindfold_secret_info/#schema-admin_password--blindfold_secret_info--location) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/blindfold_secret_info/#schema-admin_password--blindfold_secret_info--store_provider) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/clear_secret_info/#section) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/clear_secret_info/#schema-admin_password--clear_secret_info--provider_ref) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/clear_secret_info/#schema-admin_password--clear_secret_info--url) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-annotations) |
| `aws_cred` | [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/aws_cred/#section) |
| `aws_cred.name` | [aws_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/aws_cred/#schema-aws_cred--name) |
| `aws_cred.namespace` | [aws_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/aws_cred/#schema-aws_cred--namespace) |
| `aws_cred.tenant` | [aws_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/aws_cred/#schema-aws_cred--tenant) |
| `aws_region` | [aws_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-aws_region) |
| `block_all_services` | [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/block_all_services/#section) |
| `blocked_services` | [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/blocked_service/#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/blocked_service/dns/#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/blocked_service/#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/blocked_service/ssh/#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/blocked_service/web_user_interface/#section) |
| `coordinates` | [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/coordinates/#section) |
| `coordinates.latitude` | [coordinates.latitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/coordinates/#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/coordinates/#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_dns/#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_dns/#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_dns/#schema-custom_dns--outside_nameserver) |
| `custom_security_group` | [custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_security_group/#section) |
| `custom_security_group.inside_security_group_id` | [custom_security_group.inside_security_group_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_security_group/#schema-custom_security_group--inside_security_group_id) |
| `custom_security_group.outside_security_group_id` | [custom_security_group.outside_security_group_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_security_group/#schema-custom_security_group--outside_security_group_id) |
| `default_blocked_services` | [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/default_blocked_services/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-description) |
| `direct_connect_disabled` | [direct_connect_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_disabled/#section) |
| `direct_connect_enabled` | [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/#section) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/auto_asn/#section) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/#schema-direct_connect_enabled--custom_asn) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/#schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/#section) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/#schema-direct_connect_enabled--hosted_vifs--vif_list--other_region) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/same_as_site_region/#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/#schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/standard_vifs/#section) |
| `disable_encryption` | [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/disable_encryption/#section) |
| `disable_internet_vip` | [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/disable_internet_vip/#section) |
| `disk_size` | [disk_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-disk_size) |
| `egress_gateway_default` | [egress_gateway_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_gateway_default/#section) |
| `egress_nat_gw` | [egress_nat_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_nat_gw/#section) |
| `egress_nat_gw.nat_gw_id` | [egress_nat_gw.nat_gw_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_nat_gw/#schema-egress_nat_gw--nat_gw_id) |
| `egress_virtual_private_gateway` | [egress_virtual_private_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_virtual_private_gateway/#section) |
| `egress_virtual_private_gateway.vgw_id` | [egress_virtual_private_gateway.vgw_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_virtual_private_gateway/#schema-egress_virtual_private_gateway--vgw_id) |
| `enable_encryption` | [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_encryption/#section) |
| `enable_encryption.kms_key_id` | [enable_encryption.kms_key_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_encryption/#schema-enable_encryption--kms_key_id) |
| `enable_internet_vip` | [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_internet_vip/#section) |
| `f5_orchestrated_routing` | [f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/f5_orchestrated_routing/#section) |
| `f5xc_security_group` | [f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/f5xc_security_group/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-id) |
| `ingress_egress_gw` | [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/#section) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/network_policies/#section) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/network_policies/#schema-ingress_egress_gw--active_network_policies--network_policies--name) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/network_policies/#schema-ingress_egress_gw--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/network_policies/#schema-ingress_egress_gw--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw.allowed_vip_port` | [ingress_egress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/#section) |
| `ingress_egress_gw.allowed_vip_port.custom_ports` | [ingress_egress_gw.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/custom_ports/#section) |
| `ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/custom_ports/#schema-ingress_egress_gw--allowed_vip_port--custom_ports--port_ranges) |
| `ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/disable_allowed_vip_port/#section) |
| `ingress_egress_gw.allowed_vip_port.use_http_https_port` | [ingress_egress_gw.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_http_https_port/#section) |
| `ingress_egress_gw.allowed_vip_port.use_http_port` | [ingress_egress_gw.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_http_port/#section) |
| `ingress_egress_gw.allowed_vip_port.use_https_port` | [ingress_egress_gw.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/use_https_port/#section) |
| `ingress_egress_gw.allowed_vip_port_sli` | [ingress_egress_gw.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/#section) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/#section) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/custom_ports/#schema-ingress_egress_gw--allowed_vip_port_sli--custom_ports--port_ranges) |
| `ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/disable_allowed_vip_port/#section) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_http_https_port/#section) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_http_port/#section) |
| `ingress_egress_gw.allowed_vip_port_sli.use_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/use_https_port/#section) |
| `ingress_egress_gw.aws_certified_hw` | [ingress_egress_gw.aws_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/#schema-ingress_egress_gw--aws_certified_hw) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/#section) |
| `ingress_egress_gw.az_nodes.aws_az_name` | [ingress_egress_gw.az_nodes.aws_az_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/#schema-ingress_egress_gw--az_nodes--aws_az_name) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/#schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/#schema-ingress_egress_gw--az_nodes--outside_subnet--existing_subnet_id) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.az_nodes.reserved_inside_subnet` | [ingress_egress_gw.az_nodes.reserved_inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/reserved_inside_subnet/#section) |
| `ingress_egress_gw.az_nodes.workload_subnet` | [ingress_egress_gw.az_nodes.workload_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/workload_subnet/#section) |
| `ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/workload_subnet/#schema-ingress_egress_gw--az_nodes--workload_subnet--existing_subnet_id) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/workload_subnet/subnet_param/#section) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/workload_subnet/subnet_param/#schema-ingress_egress_gw--az_nodes--workload_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#section) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#schema-ingress_egress_gw--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#schema-ingress_egress_gw--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#schema-ingress_egress_gw--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#section) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#schema-ingress_egress_gw--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#schema-ingress_egress_gw--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#schema-ingress_egress_gw--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/forward_proxy_allow_all/#section) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/#schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/no_dc_cluster_group/#section) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/no_forward_proxy/#section) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/no_global_network/#section) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/no_inside_static_routes/#section) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/no_network_policy/#section) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/no_outside_static_routes/#section) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/#schema-ingress_egress_gw--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/sm_connection_public_ip/#section) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/sm_connection_pvt_ip/#section) |
| `ingress_gw` | [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/#section) |
| `ingress_gw.allowed_vip_port` | [ingress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/#section) |
| `ingress_gw.allowed_vip_port.custom_ports` | [ingress_gw.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/#section) |
| `ingress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_gw.allowed_vip_port.custom_ports.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/#schema-ingress_gw--allowed_vip_port--custom_ports--port_ranges) |
| `ingress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_gw.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/disable_allowed_vip_port/#section) |
| `ingress_gw.allowed_vip_port.use_http_https_port` | [ingress_gw.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_https_port/#section) |
| `ingress_gw.allowed_vip_port.use_http_port` | [ingress_gw.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_port/#section) |
| `ingress_gw.allowed_vip_port.use_https_port` | [ingress_gw.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_https_port/#section) |
| `ingress_gw.aws_certified_hw` | [ingress_gw.aws_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/#schema-ingress_gw--aws_certified_hw) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/#section) |
| `ingress_gw.az_nodes.aws_az_name` | [ingress_gw.az_nodes.aws_az_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/#schema-ingress_gw--az_nodes--aws_az_name) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/#section) |
| `ingress_gw.az_nodes.local_subnet.existing_subnet_id` | [ingress_gw.az_nodes.local_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/#schema-ingress_gw--az_nodes--local_subnet--existing_subnet_id) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/subnet_param/#section) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/subnet_param/#schema-ingress_gw--az_nodes--local_subnet--subnet_param--ipv4) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `instance_type` | [instance_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-instance_type) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/#section) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-labels) |
| `log_receiver` | [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/log_receiver/#section) |
| `log_receiver.name` | [log_receiver.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/log_receiver/#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/log_receiver/#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/log_receiver/#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/logs_streaming_disabled/#section) |
| `manual_routing` | [manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/manual_routing/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-namespace) |
| `no_worker_nodes` | [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/no_worker_nodes/#section) |
| `nodes_per_az` | [nodes_per_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-nodes_per_az) |
| `offline_survivability_mode` | [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/offline_survivability_mode/#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/offline_survivability_mode/enable_offline_survivability_mode/#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/offline_survivability_mode/no_offline_survivability_mode/#section) |
| `os` | [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/os/#section) |
| `os.default_os_version` | [os.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/os/default_os_version/#section) |
| `os.operating_system_version` | [os.operating_system_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/os/#schema-os--operating_system_version) |
| `private_connectivity` | [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/cloud_link/#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/cloud_link/#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/cloud_link/#schema-private_connectivity--cloud_link--namespace) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/cloud_link/#schema-private_connectivity--cloud_link--tenant) |
| `private_connectivity.inside` | [private_connectivity.inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/inside/#section) |
| `private_connectivity.outside` | [private_connectivity.outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/outside/#section) |
| `ssh_key` | [ssh_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-ssh_key) |
| `sw` | [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/sw/#section) |
| `sw.default_sw_version` | [sw.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/sw/default_sw_version/#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/sw/#schema-sw--volterra_software_version) |
| `tags` | [tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-tags) |
| `total_nodes` | [total_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/#schema-total_nodes) |
| `voltstack_cluster` | [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/#section) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/#section) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/network_policies/#section) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/network_policies/#schema-voltstack_cluster--active_network_policies--network_policies--name) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/network_policies/#schema-voltstack_cluster--active_network_policies--network_policies--namespace) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/network_policies/#schema-voltstack_cluster--active_network_policies--network_policies--tenant) |
| `voltstack_cluster.allowed_vip_port` | [voltstack_cluster.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/#section) |
| `voltstack_cluster.allowed_vip_port.custom_ports` | [voltstack_cluster.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/custom_ports/#section) |
| `voltstack_cluster.allowed_vip_port.custom_ports.port_ranges` | [voltstack_cluster.allowed_vip_port.custom_ports.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/custom_ports/#schema-voltstack_cluster--allowed_vip_port--custom_ports--port_ranges) |
| `voltstack_cluster.allowed_vip_port.disable_allowed_vip_port` | [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/disable_allowed_vip_port/#section) |
| `voltstack_cluster.allowed_vip_port.use_http_https_port` | [voltstack_cluster.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/use_http_https_port/#section) |
| `voltstack_cluster.allowed_vip_port.use_http_port` | [voltstack_cluster.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/use_http_port/#section) |
| `voltstack_cluster.allowed_vip_port.use_https_port` | [voltstack_cluster.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/use_https_port/#section) |
| `voltstack_cluster.aws_certified_hw` | [voltstack_cluster.aws_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/#schema-voltstack_cluster--aws_certified_hw) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/#section) |
| `voltstack_cluster.az_nodes.aws_az_name` | [voltstack_cluster.az_nodes.aws_az_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/#schema-voltstack_cluster--az_nodes--aws_az_name) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/#section) |
| `voltstack_cluster.az_nodes.local_subnet.existing_subnet_id` | [voltstack_cluster.az_nodes.local_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/#schema-voltstack_cluster--az_nodes--local_subnet--existing_subnet_id) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/#schema-voltstack_cluster--az_nodes--local_subnet--subnet_param--ipv4) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/dc_cluster_group/#section) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/dc_cluster_group/#schema-voltstack_cluster--dc_cluster_group--name) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/dc_cluster_group/#schema-voltstack_cluster--dc_cluster_group--namespace) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/dc_cluster_group/#schema-voltstack_cluster--dc_cluster_group--tenant) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/default_storage/#section) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/forward_proxy_allow_all/#section) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/#section) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/k8s_cluster/#section) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/k8s_cluster/#schema-voltstack_cluster--k8s_cluster--name) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/k8s_cluster/#schema-voltstack_cluster--k8s_cluster--namespace) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/k8s_cluster/#schema-voltstack_cluster--k8s_cluster--tenant) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_dc_cluster_group/#section) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_forward_proxy/#section) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_global_network/#section) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_k8s_cluster/#section) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_network_policy/#section) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_outside_static_routes/#section) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/#schema-voltstack_cluster--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/sm_connection_public_ip/#section) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/sm_connection_pvt_ip/#section) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/storage_class_list/#section) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/storage_class_list/storage_classes/#section) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/storage_class_list/storage_classes/#schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/storage_class_list/storage_classes/#schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name) |
| `vpc` | [vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/#section) |
| `vpc.new_vpc` | [vpc.new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/new_vpc/#section) |
| `vpc.new_vpc.autogenerate` | [vpc.new_vpc.autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/new_vpc/autogenerate/#section) |
| `vpc.new_vpc.name_tag` | [vpc.new_vpc.name_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/new_vpc/#schema-vpc--new_vpc--name_tag) |
| `vpc.new_vpc.primary_ipv4` | [vpc.new_vpc.primary_ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/new_vpc/#schema-vpc--new_vpc--primary_ipv4) |
| `vpc.vpc_id` | [vpc.vpc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/#schema-vpc--vpc_id) |
| `waf_signatures` | [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/waf_signatures/#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/waf_signatures/automatic/#section) |
| `waf_signatures.manual` | [waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/waf_signatures/manual/#section) |

## Next pages

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/admin_password/)
- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/aws_cred/)
- [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/block_all_services/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/blocked_services/)
- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/coordinates/)
- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_dns/)
- [custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/custom_security_group/)
- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/default_blocked_services/)
- [direct_connect_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_disabled/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/)
- [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/disable_encryption/)
- [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/disable_internet_vip/)
- [egress_gateway_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_gateway_default/)
- [egress_nat_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_nat_gw/)
- [egress_virtual_private_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/egress_virtual_private_gateway/)
- [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_encryption/)
- [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/enable_internet_vip/)
- [f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/f5_orchestrated_routing/)
- [f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/f5xc_security_group/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_gw/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/kubernetes_upgrade_drain/)
- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/log_receiver/)
- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/logs_streaming_disabled/)
- [manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/manual_routing/)
- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/no_worker_nodes/)
- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/offline_survivability_mode/)
- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/os/)
- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/private_connectivity/)
- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/sw/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/)
- [vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/vpc/)
- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/waf_signatures/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
