---
page_title: "aws_parameters.new_tgw.user_assigned"
subcategory: ""
description: "Information needed when ASNs are assigned by the user."
xcsh_docs: {"aliases": ["aws parameters new tgw user assigned"], "body_bytes": 3231, "body_sha256": "sha256:892a0b3bf8859c312c20b2a9b70f0f9cd37107d1d0218dc6ed8271d7a0ab1451", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1012321201321212-3333210011321220-1300013121113023-2010222022103120-1331133131000231-2010332022011311-0221120300323130-1130322133020332", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "new_tgw", "user_assigned"], "schema_version": 1, "sections": [{"aliases": ["aws parameters new tgw user assigned tgw asn"], "anchor": "schema-aws_parameters--new_tgw--user_assigned--tgw_asn", "description": "TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned", "tgw_asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws parameters new tgw user assigned volterra site asn"], "anchor": "schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn", "description": "F5XC Site ASN.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned", "volterra_site_asn"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Information needed when ASNs are assigned by the user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
