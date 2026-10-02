---
page_title: "f5_big_ip_aws_service.market_place_image"
subcategory: ""
description: "BIG-IP AWS Pay as You Go Image Selection."
xcsh_docs: {"aliases": ["f5 big ip aws service market place image"], "body_bytes": 2082, "body_sha256": "sha256:053284a309434cd2a4c356bb5d10ba5fd9f7e9cdf0a85d7b3767c2275449e6f8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2210031222201112-3331110222012210-2132033313310012-3320121221230020-0131021130301020-0223033002000032-3021303122131320-2001020121212112", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "market_place_image"], "schema_version": 1, "sections": [{"aliases": ["awafpay g200 mbps"], "anchor": "section", "description": "Configuration parameter for AWAFPayG200Mbps.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "market_place_image", "awafpay_g200_mbps"], "syntax": "attribute", "type": "object"}, {"aliases": ["awafpay g3 gbps"], "anchor": "section", "description": "Configuration parameter for AWAFPayG3Gbps.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "market_place_image", "awafpay_g3_gbps"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BIG-IP AWS Pay as You Go Image Selection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.market_place_image

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
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

- [awafpay_g200_mbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g200_mbps/): complete subsection reference.

- [awafpay_g3_gbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g3_gbps/): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g200_mbps/)
- [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/awafpay_g3_gbps/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
