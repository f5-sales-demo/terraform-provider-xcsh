---
page_title: "aws_parameters.new_tgw.user_assigned"
subcategory: ""
description: "Information needed when ASNs are assigned by the user."
xcsh_docs: {"aliases": ["aws parameters new tgw user assigned"], "body_bytes": 3471, "body_sha256": "sha256:ea7cfd09461d30e2a1c040795bcb5f518a4de248dd6a328b1255363cec0c8a7d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1012321201321212-3333210011321220-1300013121113023-2010222022103120-1331133131000231-2010332022011311-0221120300323130-1130322133020332", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "new_tgw", "user_assigned"], "schema_version": 1, "sections": [{"aliases": ["aws parameters new tgw user assigned tgw asn"], "anchor": "schema-aws_parameters--new_tgw--user_assigned--tgw_asn", "description": "TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned", "tgw_asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws parameters new tgw user assigned volterra site asn"], "anchor": "schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn", "description": "F5XC Site ASN.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned", "volterra_site_asn"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Information needed when ASNs are assigned by the user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_tgw.user_assigned

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [aws_parameters.new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/)
- aws_parameters.new_tgw.user_assigned

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Information needed when ASNs are assigned by the user.

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
user_assigned {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_parameters--new_tgw--user_assigned--tgw_asn"></a>

### tgw_asn property

Type: `"number"`. Optional.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(64513, 65534),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65534,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 64513
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  }
}
```

<a id="schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn"></a>

### volterra_site_asn property

Type: `"number"`. Optional.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [aws_parameters.new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
