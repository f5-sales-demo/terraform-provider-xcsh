---
page_title: "aws_parameters.new_tgw.user_assigned"
subcategory: ""
description: "Information needed when ASNs are assigned by the user."
xcsh_docs: {"aliases": ["aws parameters new tgw user assigned"], "body_bytes": 2785, "body_sha256": "sha256:4bde50ae5fa2045fe3888d8415e6ea76836320c8b246102e0d881b9abf97a4b9", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw", "path": "documentation/data-sources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3033000002110200-3033022233303022-3300212013011111-2102201030002331-0311023013100201-0221032010201101-1211022300233331-3331121021211111", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "new_tgw", "user_assigned"], "schema_version": 1, "sections": [{"aliases": ["aws parameters new tgw user assigned tgw asn"], "anchor": "schema-aws_parameters--new_tgw--user_assigned--tgw_asn", "description": "TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned", "tgw_asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws parameters new tgw user assigned volterra site asn"], "anchor": "schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn", "description": "F5XC Site ASN.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_tgw", "user_assigned", "volterra_site_asn"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Information needed when ASNs are assigned by the user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_tgw.user_assigned

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- [aws_parameters.new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/new_tgw/)
- aws_parameters.new_tgw.user_assigned

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-aws_parameters--new_tgw--user_assigned--tgw_asn"></a>

### tgw_asn property

Type: `"number"`. Computed.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

Enter F5XC Site ASN. F5XC Site ASN.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
