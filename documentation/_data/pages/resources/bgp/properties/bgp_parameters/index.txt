---
page_title: "bgp_parameters"
subcategory: ""
description: "bgp_parameters for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 4343, "body_sha256": "sha256:6c81263f990e146aab8bd4f6725e6be72726db28aaf72608ea1b3f2e0e9d2abf", "child_ids": ["xcsh-docs:resources:bgp:properties:bgp_parameters:from_site", "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "documentation/resources/bgp/properties/bgp_parameters/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["bgp_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/bgp_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bgp_parameters for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bgp_parameters

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- bgp_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bgp parameters.

Upstream description:

BGP parameters for the local site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn"),
  validators.ConflictingObjectAttributes("from_site",
    "ip_address"),
  validators.ConflictingObjectAttributes("from_site",
    "local_address"),
  validators.ConflictingObjectAttributes("ip_address",
    "local_address")}
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
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

Terraform syntax:

```terraform
bgp_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bgp_parameters--asn"></a>

### asn property

Type: `"number"`. Optional.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/from_site/): complete subsection reference.

<a id="schema-bgp_parameters--ip_address"></a>

### ip_address property

Type: `"string"`. Optional.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Upstream description:

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [local_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/local_address/): complete subsection reference.

## Next pages

- [bgp_parameters.from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/from_site/)
- [bgp_parameters.local_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/local_address/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
