---
page_title: "f5_big_ip_aws_service.aws_tgw_site_params"
subcategory: ""
description: "BIG-IP AWS TGW site specification."
xcsh_docs: {"aliases": ["f5 big ip aws service aws tgw site params"], "body_bytes": 1705, "body_sha256": "sha256:a27aba5ca0b71b38b467e3a38789a3dc196a46cec548d98a5e1b6cfb3ff0aae5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2233311222002000-2100133010320202-2122321003331003-0221303030313102-0120212003100213-2113310222000310-3130123003223212-1021001222122210", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params"], "schema_version": 1, "sections": [{"aliases": ["aws tgw site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--name", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "type": "requires"}], "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params", "aws_tgw_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BIG-IP AWS TGW site specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.aws_tgw_site_params

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- f5_big_ip_aws_service.aws_tgw_site_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BIG-IP AWS TGW Site. BIG-IP AWS TGW site specification.

Upstream description:

BIG-IP AWS TGW site specification.

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
aws_tgw_site_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/aws_tgw_site/): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/aws_tgw_site/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
