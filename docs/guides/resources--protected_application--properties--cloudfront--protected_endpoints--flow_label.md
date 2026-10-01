---
page_title: "cloudfront.protected_endpoints.flow_label"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 5530, "body_sha256": "sha256:238a323575a89023b65c1a9aefe92545d6094e31af2056d5e34c2562f8f1b32d", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:account_management", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:financial_services", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:flight", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- cloudfront.protected_endpoints.flow_label

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
```

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

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

## Direct properties

- [account_management](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management.md): complete subsection reference.

- [authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md): complete subsection reference.

- [financial_services](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services.md): complete subsection reference.

- [flight](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight.md): complete subsection reference.

- [profile_management](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management.md): complete subsection reference.

- [search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md): complete subsection reference.

- [shopping_gift_cards](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management.md)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services.md)
- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight.md)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management.md)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [xcsh_protected_application](../resources/protected_application.md)
