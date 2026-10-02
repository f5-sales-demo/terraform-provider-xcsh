---
page_title: "cloudfront.disable_aws_configuration"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cloudfront disable aws configuration"], "body_bytes": 1360, "body_sha256": "sha256:4418b4d3fef4bf02b44226e5d21a1bcc284651cb36e1c32159e39f8c23efade6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:disable_aws_configuration", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/disable_aws_configuration/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2010231000113321-3132021233310322-1232323033122132-2011123130301301-1003110013021021-2110131123331323-0202112031020112-0133123112200100", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "disable_aws_configuration"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/disable_aws_configuration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.disable_aws_configuration

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.disable_aws_configuration

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable aws configuration.

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
disable_aws_configuration = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
