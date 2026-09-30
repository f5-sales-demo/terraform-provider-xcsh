---
page_title: "passport"
subcategory: ""
description: "passport for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 8799, "body_sha256": "sha256:a16d896d5c06d1362ef89d56879b81e49d49618a6cf908b6ab3876c1d557da3e", "canonical_id": "xcsh-docs:data-sources:registration:properties:passport", "child_ids": ["xcsh-docs:data-sources:registration:properties:passport:default_os_version", "xcsh-docs:data-sources:registration:properties:passport:default_sw_version"], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:passport", "parent_id": "xcsh-docs:data-sources:registration:reference", "path": "docs/guides/data-sources--registration--properties--passport.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["passport"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/passport/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "passport for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# passport

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- passport

<a id="section"></a>

Type: `"single"`. Computed.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]",
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

## Direct properties

<a id="schema-passport--cluster_name"></a>

### cluster_name property

Type: `"string"`. Computed.

Cluster Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="schema-passport--cluster_size"></a>

### cluster_size property

Type: `"number"`. Computed.

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after..

Upstream description:

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after
installation. It does not interact with auto-scaling as only pool nodes are scaled.

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
    "ves.io.schema.rules.int32.in": "[0,1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.in": "[0,1,3]"
  }
}
```

<a id="schema-passport--cluster_type"></a>

### cluster_type property

Type: `"string"`. Computed.

Cluster Type. Cluster or grouping configuration

Upstream description:

Cluster or grouping configuration

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

- [default_os_version](data-sources--registration--properties--passport--default_os_version.md): complete subsection reference.

- [default_sw_version](data-sources--registration--properties--passport--default_sw_version.md): complete subsection reference.

<a id="schema-passport--latitude"></a>

### latitude property

Type: `"number"`. Computed.

Latitude. Geographic location of this site.

Upstream description:

Geographic location of this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-passport--longitude"></a>

### longitude property

Type: `"number"`. Computed.

Longitude. Geographic location of this site.

Upstream description:

Geographic location of this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-passport--operating_system_version"></a>

### operating_system_version property

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Upstream description:

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="schema-passport--private_network_name"></a>

### private_network_name property

Type: `"string"`. Computed.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Upstream description:

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="schema-passport--volterra_software_version"></a>

### volterra_software_version property

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

## Next pages

- [passport.default_os_version](data-sources--registration--properties--passport--default_os_version.md)
- [passport.default_sw_version](data-sources--registration--properties--passport--default_sw_version.md)
- [Property reference](data-sources--registration--reference.md)
- [xcsh_registration](../data-sources/registration.md)
