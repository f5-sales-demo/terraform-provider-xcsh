---
page_title: "ingress_egress_gw.outside_subnet.existing_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw.outside_subnet.existing_subnet for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2782, "body_sha256": "sha256:3dc5130ef1c58012c4e338bfa9558a1139fd36df03031ab47fdb925252ee02f7", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:existing_subnet", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/existing_subnet/index.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw", "outside_subnet", "existing_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/existing_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.outside_subnet.existing_subnet for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.outside_subnet.existing_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/)
- ingress_egress_gw.outside_subnet.existing_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--outside_subnet--existing_subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

## Next pages

- [ingress_egress_gw.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
