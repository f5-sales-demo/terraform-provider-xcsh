---
page_title: "f5_big_ip_aws_service.market_place_image"
subcategory: ""
description: "f5_big_ip_aws_service.market_place_image for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1911, "body_sha256": "sha256:2632f1b72f2d286ac1171d2ad77d874cd6fd47968e6b3385d54e5c6810bc1031", "canonical_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "docs/guides/resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "market_place_image"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.market_place_image for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.market_place_image

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
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

- [awafpay_g200_mbps](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g200_mbps.md): complete subsection reference.

- [awafpay_g3_gbps](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g3_gbps.md): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g200_mbps.md)
- [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g3_gbps.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
