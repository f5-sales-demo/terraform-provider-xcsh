---
page_title: "f5_big_ip_aws_service.aws_tgw_site_params"
subcategory: ""
description: "BIG-IP AWS TGW site specification."
xcsh_docs: {"aliases": ["f5 big ip aws service aws tgw site params"], "body_bytes": 1069, "body_sha256": "sha256:8848883f8bc2a32bc380b5ed1dec6f690851285889a1992b90ea2ceb49122674", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/data-sources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1130322021203231-2300212333020310-3210110221221221-1003120300332113-1323310310320132-3100020313313322-3031020210022022-2110121011300330", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service aws tgw site params aws tgw site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params", "aws_tgw_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "BIG-IP AWS TGW site specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.aws_tgw_site_params

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
- f5_big_ip_aws_service.aws_tgw_site_params

<a id="section"></a>

Type: `"single"`. Computed.

BIG-IP AWS TGW Site. BIG-IP AWS TGW site specification.

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

- [aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/aws_tgw_site/): complete subsection reference.
