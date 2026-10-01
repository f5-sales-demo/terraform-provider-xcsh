---
page_title: "adobe_commerce_connector"
subcategory: ""
description: "adobe_commerce_connector for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1910, "body_sha256": "sha256:84c500a0cf462ea94191820c6d9269156ecaf922b28284cb12cc3663dd7b5bb2", "canonical_id": "xcsh-docs:resources:protected_application:properties:adobe_commerce_connector", "child_ids": [], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:adobe_commerce_connector", "parent_id": "xcsh-docs:resources:protected_application:reference", "path": "docs/guides/resources--protected_application--properties--adobe_commerce_connector.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["adobe_commerce_connector"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/adobe_commerce_connector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "adobe_commerce_connector for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# adobe_commerce_connector

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- adobe_commerce_connector

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: adobe\_commerce\_connector, big\_ip\_iapp, cloudflare, cloudfront, custom\_connector,
f5\_big\_ip, salesforce\_commerce\_connector\] Configuration parameter for adobe commerce connector.

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

OneOf alternatives in this subsection:

- [adobe_commerce_connector](resources--protected_application--properties--adobe_commerce_connector.md#section)
- [big_ip_iapp](resources--protected_application--properties--big_ip_iapp.md#section)
- [cloudflare](resources--protected_application--properties--cloudflare.md#section)
- [cloudfront](resources--protected_application--properties--cloudfront.md#section)
- [custom_connector](resources--protected_application--properties--custom_connector.md#section)
- [f5_big_ip](resources--protected_application--properties--f5_big_ip.md#section)
- [salesforce_commerce_connector](resources--protected_application--properties--salesforce_commerce_connector.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
adobe_commerce_connector = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--protected_application--reference.md)
- [xcsh_protected_application](../resources/protected_application.md)
