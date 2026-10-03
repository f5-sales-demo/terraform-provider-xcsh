---
page_title: "cookie_matchers.check_present"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cookie matchers check present"], "body_bytes": 1328, "body_sha256": "sha256:e43ff1e80142a3dd065adc7bec103d634256662d9afb7c09eccf0bf2f0774e7e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:check_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers", "path": "documentation/resources/service_policy_rule/properties/cookie_matchers/check_present/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0301100030213101-1111323301332122-1110231222230111-0301002120302310-2122012122030133-1220031231310322-3001102022032332-3122110120122303", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_matchers", "check_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/cookie_matchers/check_present/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_matchers.check_present

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/)
- cookie_matchers.check_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
