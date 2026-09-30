---
page_title: "f5_big_ip_aws_service.market_place_image"
subcategory: ""
description: "f5_big_ip_aws_service.market_place_image for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1530, "body_sha256": "sha256:60af539bb731b5d6fc826470f98867feecdd4a4fd233dd151998c7dd3037e592", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "docs/guides/data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "market_place_image"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.market_place_image for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# f5_big_ip_aws_service.market_place_image

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md)
- f5_big_ip_aws_service.market_place_image

<a id="section"></a>

Type: `"single"`. Computed.

BIG-IP AWS Pay as You Go Image Selection.

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

## Direct properties

- [awafpay_g200_mbps](data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g200_mbps.md): complete subsection reference.

- [awafpay_g3_gbps](data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g3_gbps.md): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g200_mbps.md)
- [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g3_gbps.md)
- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
