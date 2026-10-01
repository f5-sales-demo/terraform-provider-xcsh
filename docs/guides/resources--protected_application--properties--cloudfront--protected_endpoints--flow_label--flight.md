---
page_title: "cloudfront.protected_endpoints.flow_label.flight"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.flight for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1805, "body_sha256": "sha256:485cf8bce6e35d889cba2b868bce4900c93a8ac67c5940718edb8521bdad76c7", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:flight", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:flight:checkin"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:flight", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "flight"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/flight/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.flight for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.flight

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- cloudfront.protected_endpoints.flow_label.flight

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

## Direct properties

- [checkin](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight--checkin.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.flight.checkin](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight--checkin.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [xcsh_protected_application](../resources/protected_application.md)
