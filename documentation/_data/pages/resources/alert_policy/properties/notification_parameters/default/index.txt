---
page_title: "notification_parameters.default"
subcategory: "Monitoring"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["notification parameters default"], "body_bytes": 1285, "body_sha256": "sha256:09aab7d36fd5c5741df648a0d84738ad0cefd05d4d6411f0eba19342c2c4cad9", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "parent_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "path": "documentation/resources/alert_policy/properties/notification_parameters/default/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2123120301301120-3113010213203313-2220210202320102-0302201203112102-1333203130012323-3112010320233112-0102122320020203-3003231132010032", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["notification_parameters", "default"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/notification_parameters/default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters.default

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/)
- notification_parameters.default

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
default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
