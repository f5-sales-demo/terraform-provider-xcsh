---
page_title: "aws"
subcategory: ""
description: "CloudLink for AWS Cloud Provider."
xcsh_docs: {"aliases": ["aws"], "body_bytes": 2930, "body_sha256": "sha256:7c6569b3aeba8c3b09e953b6e482e1c71ca9eb2550641e2721cb5177269b3aab", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:aws:aws_cred", "xcsh-docs:data-sources:cloud_link:properties:aws:byoc"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:aws", "parent_id": "xcsh-docs:data-sources:cloud_link:reference", "path": "documentation/data-sources/cloud_link/properties/aws/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws"], "schema_version": 1, "sections": [{"aliases": ["aws aws cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:aws_cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "aws_cred"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws byoc"], "anchor": "section", "description": "List of Bring You Own Connection.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "byoc"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws custom asn"], "anchor": "schema-aws--custom_asn", "description": "Exclusive with F5XC will use custom ASN to create a Direct Connect Gateway.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "custom_asn"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/aws/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "CloudLink for AWS Cloud Provider.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- aws

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: aws, gcp\] Amazon Web Services(AWS) CloudLink Provider. CloudLink for AWS Cloud Provider.

Upstream description:

CloudLink for AWS Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]",
  "x-ves-oneof-field-direct_connect_gateway_asn_choice": "[\"custom_asn\"]"
}
```

OneOf alternatives in this subsection:

- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/#section)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/aws_cred/): complete subsection reference.

- [byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/): complete subsection reference.

<a id="schema-aws--custom_asn"></a>

### custom_asn property

Type: `"number"`. Computed.

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Upstream description:

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4294967294,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  }
}
```

## Next pages

- [aws.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/aws_cred/)
- [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
