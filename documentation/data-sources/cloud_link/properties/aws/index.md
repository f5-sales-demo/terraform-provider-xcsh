---
page_title: "aws"
subcategory: ""
description: "CloudLink for AWS Cloud Provider."
xcsh_docs: {"aliases": ["aws"], "body_bytes": 2930, "body_sha256": "sha256:9ae3d5b09dcd5540d67bd6666228e1cc1d825f96a29dcdfd594f11084212621e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:aws:aws_cred", "xcsh-docs:data-sources:cloud_link:properties:aws:byoc"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:aws", "parent_id": "xcsh-docs:data-sources:cloud_link:reference", "path": "documentation/data-sources/cloud_link/properties/aws/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws"], "schema_version": 1, "sections": [{"aliases": ["aws cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:aws_cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "aws_cred"], "syntax": "attribute", "type": "object"}, {"aliases": ["byoc"], "anchor": "section", "description": "List of Bring You Own Connection.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "byoc"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom asn"], "anchor": "schema-aws--custom_asn", "description": "Exclusive with F5XC will use custom ASN to create a Direct Connect Gateway.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "custom_asn"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/aws/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CloudLink for AWS Cloud Provider.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
