---
page_title: "passport"
subcategory: ""
description: "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval."
xcsh_docs: {"aliases": ["passport"], "body_bytes": 10207, "body_sha256": "sha256:758bb198db44316280ff437a35a18f543d7ed06e642edd0a9038dc7735814367", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:properties:passport:default_os_version", "xcsh-docs:resources:registration:properties:passport:default_sw_version"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:passport", "parent_id": "xcsh-docs:resources:registration:reference", "path": "documentation/resources/registration/properties/passport/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1120101321201223-0133323222022210-1001000011103111-2322312323001321-3201033220310002-2220113311003213-3101110103313231-1211213320021213", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [{"anchor": "schema-passport--operating_system_version", "enforcement": "provider-schema", "group": "passport:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport", "type": "conflicts"}, {"anchor": "schema-passport--volterra_software_version", "enforcement": "provider-schema", "group": "passport:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "passport:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport:default_os_version", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "passport:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport:default_sw_version", "type": "conflicts"}, {"anchor": "schema-passport--cluster_name", "enforcement": "provider-schema", "group": "passport:RequiredObjectAttributes:cluster_name,cluster_type,latitude,longitude", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport", "type": "requires"}, {"anchor": "schema-passport--cluster_type", "enforcement": "provider-schema", "group": "passport:RequiredObjectAttributes:cluster_name,cluster_type,latitude,longitude", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport", "type": "requires"}, {"anchor": "schema-passport--latitude", "enforcement": "provider-schema", "group": "passport:RequiredObjectAttributes:cluster_name,cluster_type,latitude,longitude", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport", "type": "requires"}, {"anchor": "schema-passport--longitude", "enforcement": "provider-schema", "group": "passport:RequiredObjectAttributes:cluster_name,cluster_type,latitude,longitude", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:passport", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["passport"], "schema_version": 1, "sections": [{"aliases": ["cluster name"], "anchor": "schema-passport--cluster_name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "cluster_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cluster size"], "anchor": "schema-passport--cluster_size", "description": "Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time, cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after installation. It does not", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "cluster_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["cluster type"], "anchor": "schema-passport--cluster_type", "description": "Cluster or grouping configuration", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "cluster_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["default os version"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:registration:properties:passport:default_os_version", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "default_os_version"], "syntax": "attribute", "type": "object"}, {"aliases": ["default sw version"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:registration:properties:passport:default_sw_version", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "default_sw_version"], "syntax": "attribute", "type": "object"}, {"aliases": ["latitude"], "anchor": "schema-passport--latitude", "description": "Geographic location of this site.", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "latitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["longitude"], "anchor": "schema-passport--longitude", "description": "Geographic location of this site.", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "longitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["operating system version"], "anchor": "schema-passport--operating_system_version", "description": "Exclusive with Operating System Version is optional parameter, which allows to specify target SW version for particular site e.g. 7.2009.10.", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "operating_system_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["private network name"], "anchor": "schema-passport--private_network_name", "description": "Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink, CloudLink and L3VPN.", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "private_network_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["volterra software version"], "anchor": "schema-passport--volterra_software_version", "description": "Exclusive with F5XC Software Version is optional parameter, which allows to specify target SW version for particular site e.g. Crt-20210329-1002.", "document_id": "xcsh-docs:resources:registration:properties:passport", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["passport", "volterra_software_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/passport/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# passport

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- passport

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_name",
    "cluster_type",
    "latitude",
    "longitude"),
  validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version"),
  validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
```

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

Terraform syntax:

```terraform
passport {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-passport--cluster_name"></a>

### cluster_name property

Type: `"string"`. Optional.

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

Type: `"number"`. Optional.

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

Type: `"string"`. Optional.

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

- [default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/passport/default_os_version/): complete subsection reference.

- [default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/passport/default_sw_version/): complete subsection reference.

<a id="schema-passport--latitude"></a>

### latitude property

Type: `"number"`. Optional.

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

Type: `"number"`. Optional.

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

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Upstream description:

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

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

Type: `"string"`. Optional.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Upstream description:

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

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

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

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

- [passport.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/passport/default_os_version/)
- [passport.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/passport/default_sw_version/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
