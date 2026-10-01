---
page_title: "direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect"
subcategory: "Infrastructure"
description: "direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2636, "body_sha256": "sha256:81588266e243fbe1e5abfaf7c6308c6bfce48df0744d235b3d9d1cacf6e3e750", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "path": "docs/guides/resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_direct_connect"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [direct_connect_enabled](resources--aws_vpc_site--properties--direct_connect_enabled.md)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs.md)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
```

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

Terraform syntax:

```terraform
site_registration_over_direct_connect {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name"></a>

### cloudlink_network_name property

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
