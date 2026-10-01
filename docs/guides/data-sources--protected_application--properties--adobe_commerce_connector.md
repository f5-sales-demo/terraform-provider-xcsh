---
page_title: "adobe_commerce_connector"
subcategory: ""
description: "adobe_commerce_connector for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1876, "body_sha256": "sha256:412250e6a6e928608e33a95fea85945b88fabaa4004a36dd1c383198da9ff04d", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:adobe_commerce_connector", "child_ids": [], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:adobe_commerce_connector", "parent_id": "xcsh-docs:data-sources:protected_application:reference", "path": "docs/guides/data-sources--protected_application--properties--adobe_commerce_connector.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["adobe_commerce_connector"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/adobe_commerce_connector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "adobe_commerce_connector for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# adobe_commerce_connector

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- adobe_commerce_connector

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

- [adobe_commerce_connector](data-sources--protected_application--properties--adobe_commerce_connector.md#section)
- [big_ip_iapp](data-sources--protected_application--properties--big_ip_iapp.md#section)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md#section)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md#section)
- [custom_connector](data-sources--protected_application--properties--custom_connector.md#section)
- [f5_big_ip](data-sources--protected_application--properties--f5_big_ip.md#section)
- [salesforce_commerce_connector](data-sources--protected_application--properties--salesforce_commerce_connector.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--protected_application--reference.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
