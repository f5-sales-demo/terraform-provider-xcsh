---
page_title: "cloudfront.protected_endpoints.flow_label.financial_services"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.financial_services for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1993, "body_sha256": "sha256:8535d3b7b04f36e2ed36e078fa26ac36644cf196b17d11516ec2ca90f9bd0e15", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:financial_services", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:financial_services:apply", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:financial_services:money_transfer"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:financial_services", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "financial_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/financial_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.financial_services for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudfront.protected_endpoints.flow_label.financial_services

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- cloudfront.protected_endpoints.flow_label.financial_services

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

## Direct properties

- [apply](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services--apply.md): complete subsection reference.

- [money_transfer](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services--money_transfer.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.financial_services.apply](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services--apply.md)
- [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services--money_transfer.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
