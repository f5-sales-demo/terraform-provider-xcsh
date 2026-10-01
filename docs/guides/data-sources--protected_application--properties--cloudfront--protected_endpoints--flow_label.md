---
page_title: "cloudfront.protected_endpoints.flow_label"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3647, "body_sha256": "sha256:b798df2e72053ac0eada044d10b7084361e2da24cbb0c3beb079fb10ee79d214", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:account_management", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:financial_services", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:flight", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- cloudfront.protected_endpoints.flow_label

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

## Direct properties

- [account_management](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management.md): complete subsection reference.

- [authentication](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md): complete subsection reference.

- [financial_services](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services.md): complete subsection reference.

- [flight](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight.md): complete subsection reference.

- [profile_management](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management.md): complete subsection reference.

- [search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md): complete subsection reference.

- [shopping_gift_cards](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management.md)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services.md)
- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight.md)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management.md)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
