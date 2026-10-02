---
page_title: "waf_action.waf_skip_processing"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["waf action waf skip processing"], "body_bytes": 1292, "body_sha256": "sha256:ba5a29a4a3b70f22541977e545c412b31f9dee4334a12f369b1daa05eab3b46e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "path": "documentation/resources/service_policy_rule/properties/waf_action/waf_skip_processing/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2300310201333330-3133013011222002-2102212230202020-3113323220303320-0031002131200313-1213322102302103-0221203123001033-2303330302222312", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_action", "waf_skip_processing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/waf_skip_processing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action.waf_skip_processing

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/)
- waf_action.waf_skip_processing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
waf_skip_processing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
