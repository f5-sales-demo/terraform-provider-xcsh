---
page_title: "vn_config.allowed_vip_port.use_http_port"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["vn config allowed vip port use http port"], "body_bytes": 1153, "body_sha256": "sha256:ff1bd563ecf1aac8b649941c1908206a8f7fa7ab0634bae82aa7fd7c16d93b71", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port", "path": "documentation/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_port/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3133311202220330-0233223323100310-2111323013002320-2102213311200200-3002231323100230-3011023133013131-3212302031330030-1330033223123200", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "allowed_vip_port", "use_http_port"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_port/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port.use_http_port

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/)
- vn_config.allowed_vip_port.use_http_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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
use_http_port = {}
```

This is an empty object or choice marker. It has no direct properties.
