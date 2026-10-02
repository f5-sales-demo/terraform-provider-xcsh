---
page_title: "f5_big_ip_aws_service.market_place_image"
subcategory: ""
description: "BIG-IP AWS Pay as You Go Image Selection."
xcsh_docs: {"aliases": ["f5 big ip aws service market place image"], "body_bytes": 2364, "body_sha256": "sha256:d7d8026c136215df560203f8d4bfe1ab8c65d7d63393039260d8dde320b5c7e5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3121221100302221-2312112212213100-2300030203321033-0023133210013231-0322231123223011-1213333100111231-1001202332212022-3001113302213202", "registry_path": "docs/guides/resources--nfv_service--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.market_place_image:ConflictingObjectAttributes:awafpay_g200_mbps,awafpay_g3_gbps", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.market_place_image:ConflictingObjectAttributes:awafpay_g200_mbps,awafpay_g3_gbps", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "market_place_image"], "schema_version": 1, "sections": [{"aliases": ["awafpay g200 mbps"], "anchor": "section", "description": "Configuration parameter for AWAFPayG200Mbps.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "market_place_image", "awafpay_g200_mbps"], "syntax": "attribute", "type": "object"}, {"aliases": ["awafpay g3 gbps"], "anchor": "section", "description": "Configuration parameter for AWAFPayG3Gbps.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "market_place_image", "awafpay_g3_gbps"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BIG-IP AWS Pay as You Go Image Selection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.market_place_image

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- f5_big_ip_aws_service.market_place_image

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BIG-IP AWS Pay as You Go Image Selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("awafpay_g200_mbps",
    "awafpay_g3_gbps")}
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
  "x-ves-oneof-field-ami_choice": "[\"AWAFPayG200Mbps\",\"AWAFPayG3Gbps\",\"BestPlusPayG200Mbps\",\"best_plus_payg_1gbps\"]"
}
```

Terraform syntax:

```terraform
market_place_image {
  # Configure direct properties listed below.
}
```

## Direct properties

- [awafpay_g200_mbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g200_mbps/): complete subsection reference.

- [awafpay_g3_gbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g3_gbps/): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g200_mbps/)
- [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g3_gbps/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
